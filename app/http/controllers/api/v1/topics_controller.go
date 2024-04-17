package v1

import (
	"fmt"
	"goWeb/app/models/topic"
	"goWeb/app/policies"
	"goWeb/app/requests"
	"goWeb/pkg/auth"
	"goWeb/pkg/logger"
	"goWeb/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.AddTopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok { //只能修改自己的文章，判断id
		response.Abort403(c)
		return
	}

	request := requests.AddTopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	request := requests.TopicListRequest{}
	c.ShouldBindJSON(&request) //为什么要多绑定依次才可以？
	fmt.Println(request)
	if ok := requests.Validate(c, &request, requests.TopicList); !ok {
		return
	}

	perPage, err1 := strconv.Atoi(request.PerPage)
	currentPage, err2 := strconv.Atoi(request.CurrentPage)
	if err1 != nil || err2 != nil {
		logger.Error("类型转换失败")
		return
	}

	topics := topic.GetList(perPage, currentPage)
	if len(topics) != 0 { //GetList里面没做错误，通过
		response.Data(c, topics)
	} else {
		response.Abort500(c, "未获取到话题列表")
	}
}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, topicModel)
}
func (ctrl *TopicsController) Search(c *gin.Context) {
	request := requests.TopicSearchRequest{}
	c.ShouldBindJSON(&request)

	topics := topic.Search(request.Keyword)
	if len(topics) == 0 {
		response.Abort404(c, "抱歉，找不到您想搜索的话题")
	} else {
		response.Data(c, topics)
	}

}
