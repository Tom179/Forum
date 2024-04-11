// Package models 模型通用属性和方法
package models

import (
	"github.com/spf13/cast"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"` //绑定数据库字段、json格式
} //omitempty标签，该标签表示在将BaseModel结构体转换成JSON格式时，如果ID属性值为空，则不会将该属性输出

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

func (a BaseModel) GetStringID() string {
	//fmt.Println(a.ID)
	return cast.ToString(a.ID)
}
