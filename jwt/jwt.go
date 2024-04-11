// Package jwt 处理 JWT 认证
package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
)

var ( //定义错误
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// JWT 定义一个jwt对象
type JWT struct {
	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte
	// 刷新 Token 的最大过期时间
	MaxRefresh time.Duration
}

// JWTCustomClaims 自定义载荷:userID，用户名、过期时间
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`
	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	//如下↓
	jwtpkg.StandardClaims //
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte("33446a9dcf9ea060a0a6532b166da32f304af0de"), //配置方案中的app.key:这是密钥
		MaxRefresh: time.Duration(10000000) * time.Minute,
	}
}

// ParserToken 解析 Token，中间件中调用
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {

	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	// 1. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)
	// 2. 解析出错

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 3. 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 更新 Token，用以提供 refresh token 接口
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {

	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*JWTCustomClaims)

	// 5. 检查是否过了『最大允许刷新的时间』
	x := time.Now().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		// 修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh

	/*这个函数是用于更新 Token 的，通常用于提供刷新 Token 的接口。让我们逐步解析这个函数的逻辑：

	首先，从请求的 Header 中获取 Token 字符串。
	调用 parseTokenString 函数解析传入的 Token 字符串，并获得解析后的 Token 对象。
	如果解析 Token 时发生错误，进一步检查错误类型。如果错误不是过期错误 (ValidationErrorExpired)，则返回该错误，表示 Token 不合法。
	如果解析 Token 成功，获取 Token 中的声明信息 (JWTCustomClaims)。
	检查 Token 是否已经过了最大允许刷新的时间。这里通过与当前时间比较 Token 的签发时间 (IssuedAt) 来判断是否超过了最大允许刷新的时间。
	如果 Token 还在允许刷新的时间范围内，则修改 Token 的过期时间 (ExpiresAt) 为新的过期时间，并调用 createToken 函数生成新的 Token 字符串并返回。
	如果 Token 已经超过最大允许刷新的时间，则返回错误 ErrTokenExpiredMaxRefresh，表示无法刷新 Token。*/
}

// IssueToken 生成  Token，在登录成功时调用
func (jwt *JWT) CreatToken(userID string, userName string) string {

	// 1. 构造用户 claims 信息(负荷)
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		userID,
		userName,
		expireAtTime,
		jwtpkg.StandardClaims{
			NotBefore: time.Now().Unix(), // 签名生效时间
			IssuedAt:  time.Now().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireAtTime,      // 签名过期时间
			Issuer:    "goWeb",           // 签名颁发者
		},
	}

	// 2. 根据 claims 生成token对象
	token, err := jwt.createToken(claims)
	if err != nil {
		//logger.LogIf(err)
		fmt.Println(err)
		return ""
	}

	return token
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	// 使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey) //生成jwt
}

// expireAtTime 过期时间
func (jwt *JWT) expireAtTime() int64 {

	//timenow := app.TimenowInTimezone()
	timenow := time.Now()

	var expireTime int64 = 60

	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}

// parseTokenString 使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	}) //ParseWithClaims函数传入jwt、载荷对象、密钥解密方法
}

/*并解析出 Token 中的声明信息，将其填充到 JWTCustomClaims 结构体中。如果 Token 有效且解析成功，
函数会返回解析后的 Token 对象。
解析后的 Token 对象，其中包含了 Token 的声明信息。
这个 Token 对象可以用于后续的授权和身份验证操作。


jwt.Parse函数只负责解析令牌的结构，需要手动验证签名的有效性，
而 jwt.ParseWithClaims 在解析令牌的同时也会验证签名的有效性，
并将声明信息存储在指定的声明结构体中。*/

// getTokenFromHeader 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization") //从authorization头中获得jwt
	if authHeader == "" {                               //为空
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2) // 获得的是"Bearer 'jwt'",所以我们现在只要得到jwt部分
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil //返回jwt
}
