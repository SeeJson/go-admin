package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type EmMajorInfo struct {
	api.Api
}

// GetPage 获取专业信息列表
// @Summary 获取专业信息列表
// @Description 获取专业信息列表
// @Tags 专业信息
// @Param majorName query string false "所属专业"
// @Param collegeName query string false "所属学院"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EmMajorInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-major-info [get]
// @Security Bearer
func (e EmMajorInfo) GetPage(c *gin.Context) {
	req := dto.EmMajorInfoGetPageReq{}
	s := service.EmMajorInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EmMajorInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取专业信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取专业信息
// @Summary 获取专业信息
// @Description 获取专业信息
// @Tags 专业信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EmMajorInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/em-major-info/{id} [get]
// @Security Bearer
func (e EmMajorInfo) Get(c *gin.Context) {
	req := dto.EmMajorInfoGetReq{}
	s := service.EmMajorInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.EmMajorInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取专业信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建专业信息
// @Summary 创建专业信息
// @Description 创建专业信息
// @Tags 专业信息
// @Accept application/json
// @Product application/json
// @Param data body dto.EmMajorInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/em-major-info [post]
// @Security Bearer
func (e EmMajorInfo) Insert(c *gin.Context) {
	req := dto.EmMajorInfoInsertReq{}
	s := service.EmMajorInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建专业信息失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改专业信息
// @Summary 修改专业信息
// @Description 修改专业信息
// @Tags 专业信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EmMajorInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/em-major-info/{id} [put]
// @Security Bearer
func (e EmMajorInfo) Update(c *gin.Context) {
	req := dto.EmMajorInfoUpdateReq{}
	s := service.EmMajorInfo{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改专业信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除专业信息
// @Summary 删除专业信息
// @Description 删除专业信息
// @Tags 专业信息
// @Param data body dto.EmMajorInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/em-major-info [delete]
// @Security Bearer
func (e EmMajorInfo) Delete(c *gin.Context) {
	s := service.EmMajorInfo{}
	req := dto.EmMajorInfoDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除专业信息失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
