package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EmClassInfo struct {
	service.Service
}

// GetPage 获取EmClassInfo列表
func (e *EmClassInfo) GetPage(c *dto.EmClassInfoGetPageReq, p *actions.DataPermission, list *[]models.EmClassInfo, count *int64) error {
	var err error
	var data models.EmClassInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmClassInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmClassInfo对象
func (e *EmClassInfo) Get(d *dto.EmClassInfoGetReq, p *actions.DataPermission, model *models.EmClassInfo) error {
	var data models.EmClassInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmClassInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmClassInfo对象
func (e *EmClassInfo) Insert(c *dto.EmClassInfoInsertReq) error {
    var err error
    var data models.EmClassInfo
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmClassInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmClassInfo对象
func (e *EmClassInfo) Update(c *dto.EmClassInfoUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.EmClassInfo{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("EmClassInfoService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除EmClassInfo
func (e *EmClassInfo) Remove(d *dto.EmClassInfoDeleteReq, p *actions.DataPermission) error {
	var data models.EmClassInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEmClassInfo error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
