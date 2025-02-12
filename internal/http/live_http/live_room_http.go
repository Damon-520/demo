package live_http

import (
	"demoapi/internal/conf"
	apies "demoapi/internal/elasticsearch"
	"demoapi/internal/pkg/errorx"
	"demoapi/internal/pkg/response"
	liveSc "demoapi/internal/service/live_service"
	"demoapi/libs/timex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"net/http"
)

type LiveRoomHttp struct {
	liveRoomSc *liveSc.LiveRoomService
	log        *log.Helper
	config     *conf.Config
}

func NewLiveRoomHttp(
	liveRoomSc *liveSc.LiveRoomService,
	logger log.Logger,
	config *conf.Config,
) *LiveRoomHttp {
	return &LiveRoomHttp{
		liveRoomSc: liveRoomSc,
		config:     config,
		log:        log.NewHelper(log.With(logger, "x_module", "http/NewLiveRoomHttp")),
	}
}

// Create 直播间详情
func (l *LiveRoomHttp) Create(c *gin.Context) {

	var err error
	resp := response.NewResponse()

	// 参数验证
	var req LiveRoomCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Create:接口参数解析与验证错误:Err:%v", err)
		resp.Error(c, errorx.Cause(err))
		return
	}

	// fmt.Printf("req: %+v\n", req)

	// 权限
	// uInfo, ok := auth.GetAdminInfo(c)
	// if !ok {
	// 	l.log.WithContext(c).Info("LiveVideoPointHttp:Create:认证错误")
	// 	resp.Error(c, errorx.ErrAuthFail)
	// 	return
	// }

	params := liveSc.LiveRoomAddParams_{
		LiveType: 1,
	}

	_ = copier.Copy(&params, &req)

	fmt.Printf("params: %+v\n", params)

	lastId, err := l.liveRoomSc.Add(c, params)
	if err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Create:添加错误:Err:%v", err)
		resp.Error(c, errorx.Cause(err))
		return
	}

	resp.JsonRaw(c, LiveRoomCreateResult{LastId: lastId})
}

// Info 直播间详情
func (l *LiveRoomHttp) Info(c *gin.Context) {

	resp := response.NewResponse()

	// 参数验证
	var req LiveRoomInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Info:接口参数解析与验证错误:Err:%v", err)
		resp.Error(c, errorx.ErrRequest)
		return
	}

	params := liveSc.LiveRoomInfoParams_{LiveRoomId: req.Id}

	liveRoom, err := l.liveRoomSc.Info(c, params)
	if err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Info:获取结果错误:Err:%v", err)
		resp.Error(c, errorx.ErrRequest)
		return
	}

	var liveRoomVo LiveRoomVo
	_ = copier.Copy(&liveRoomVo, &liveRoom)

	liveRoomVo.UpdatedAt = timex.TimeFormat(liveRoom.UpdatedAt, timex.DefaultLayout)
	liveRoomVo.CreatedAt = timex.TimeFormat(liveRoom.CreatedAt, timex.DefaultLayout)

	resp.JsonRaw(c, liveRoomVo)
}

// List 直播间列表
func (l *LiveRoomHttp) List(c *gin.Context) {

	resp := response.NewResponse()

	// 参数验证
	var req LiveRoomListRequest
	if err := c.ShouldBind(&req); err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:List:接口参数解析与验证错误:Err:%v", err)
		resp.Error(c, errorx.ErrRequest)
		return
	}

	// 查询条件
	params := liveSc.LiveRoomListParams_{
		LikeName:   req.Name,
		IsDisabled: req.IsDisabled,
		StartDate:  req.DateRange.StartDate,
		EndDate:    req.DateRange.EndDate,
		Page:       req.Page,
		Limit:      req.Limit,
	}

	// 查询数据
	list, total, err := l.liveRoomSc.List(c, params)
	if err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:List:查询结果失败:Err:%v", err)
		resp.Error(c, errorx.ErrRequest)
		return
	}

	// 返回结果
	result := LiveRoomListResult{}
	result.List = liveRoom2LiveRoomVo(list)
	result.PageInfo.Limit = req.Limit
	result.PageInfo.Total = total

	resp.JsonRaw(c, result)
}

// Edit 直播间详情
func (l *LiveRoomHttp) Edit(c *gin.Context) {

	resp := response.NewResponse()

	// 参数验证
	var req LiveRoomEditRequest
	if err := c.ShouldBind(&req); err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Edit:接口参数解析与验证错误:Err:%v", err)
		resp.Error(c, errorx.ErrRequest)
		return
	}

	// 权限
	// uInfo, ok := auth.GetAdminInfo(c)
	// if !ok {
	// 	l.log.WithContext(c).Info("LiveVideoPointHttp:Edit:认证错误")
	// 	resp.Error(c, errorx.ErrAuthFail)
	// 	return
	// }

	params := liveSc.LiveRoomEditParams_{}

	_ = copier.Copy(&params, &req)

	// fmt.Printf("params: %+v\n", params)

	rows, err := l.liveRoomSc.Edit(c, params)
	if err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Edit:编辑错误:Err:%v", err)
		resp.Error(c, errorx.Cause(err))
		return
	}

	resp.JsonRaw(c, LiveRoomEditResult{Rows: int(rows)})
}

// Update 直播间详情
func (l *LiveRoomHttp) Update(c *gin.Context) {

	resp := response.NewResponse()

	// 参数验证
	var req LiveRoomUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Update:接口参数解析与验证错误:Err:%v", err)
		resp.Error(c, errorx.ErrRequest)
		return
	}

	// 权限
	// uInfo, ok := auth.GetAdminInfo(c)
	// if !ok {
	// 	l.log.WithContext(c).Info("LiveVideoPointHttp:Update:认证错误")
	// 	resp.Error(c, errorx.ErrAuthFail)
	// 	return
	// }

	params := liveSc.LiveRoomUpdateParams_{}

	_ = copier.Copy(&params, &req)

	// fmt.Printf("params: %+v\n", params)

	rows, err := l.liveRoomSc.Update(c, params)
	if err != nil {
		l.log.WithContext(c).Infof("LiveRoomHttp:Update:编辑错误:Err:%v", err)
		resp.Error(c, errorx.Cause(err))
		return
	}

	resp.JsonRaw(c, LiveRoomEditResult{Rows: int(rows)})
}

func (h *LiveRoomHttp) Es(c *gin.Context) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "Test Document",
			},
		},
	}

	res, err := apies.SearchDocument("test_index", query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
