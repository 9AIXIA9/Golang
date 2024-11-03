package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/myerrors"
	"bluebell/settings"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// PostCreateHandler 用于处理创建帖子的路由
func PostCreateHandler(c *gin.Context) {
	//1.获取及处理参数
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error(myerrors.InvalidParam.Error(), zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c取到当前发请求的用户ID

	userID, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error(err.Error())
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error(myerrors.PostCreate.Error(), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)
}

// PostGetHandler 通过id查询帖子的路由
func PostGetHandler(c *gin.Context) {
	//获取数据(帖子的id)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
		ResponseError(c, CodeInvalidParam)
		return
	}
	//查询(根据id取数据)
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error(err.Error())
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
	return
}

// GetPostListHandler 用于处理获取帖子列表的路由
func GetPostListHandler(c *gin.Context) {
	pageNum, pageSize, err := GetPageInfo(c)
	if err != nil {
		zap.L().Error(myerrors.PostPageGet.Error())
		ResponseError(c, CodeServerBusy)
		return
	}
	//获取数据
	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error(myerrors.PostListGet.Error())
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	//返回响应
}

// GetPostListHandler2 升级版帖子接口实现
// 按创建时间排序或者按照分数排序
// 获取参数去redis查找id
// 根据id查找帖子详情
// 返回响应
func GetPostListHandler2(c *gin.Context) {
	//GET请求参数：/api/v1/posts2?page=1&size=10&order=time
	//获取分页信息
	//指定默认值
	p := &models.ParamPostList{
		Page:  settings.Config.PostDefaultPage,
		Size:  settings.Config.PostDefaultSize,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("PostGetListHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	} //请求用query传参
	//c.ShouldBind()      //动态选择请求数据类型获取传参
	//c.ShouldBindJSON() // 请求用json传参

	//获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error(myerrors.PostListGet.Error())
		ResponseError(c, CodeServerBusy)

		return
	}
	ResponseSuccess(c, data)
	//返回响应
}

// GetCommunityPostListHandler 根据社区去查询帖子列表
func GetCommunityPostListHandler(c *gin.Context) {
	//GET请求参数：/api/v1/posts2?page=1&size=10&order=time
	//获取分页信息
	//指定默认值
	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  settings.Config.PostDefaultPage,
			Size:  settings.Config.PostDefaultSize,
			Order: models.OrderTime,
		},
		CommunityID: settings.Config.CommunityDefault,
	}
	//请求用query传参
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	zap.L().Debug("Controller GetCommunityPostListHandler:", zap.Any("p:", p))

	//c.ShouldBind()      //动态选择请求数据类型获取传参
	//c.ShouldBindJSON() // 请求用json传参

	//获取数据
	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error(myerrors.PostListGet.Error())
		ResponseError(c, CodeServerBusy)

		return
	}
	ResponseSuccess(c, data)
	//返回响应
}
