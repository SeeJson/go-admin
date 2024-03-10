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

type EmTextbook struct {
	service.Service
}

// GetPage 获取EmTextbook列表
func (e *EmTextbook) GetPage(c *dto.EmTextbookGetPageReq, p *actions.DataPermission, list *[]models.EmTextbook, count *int64) error {
	var err error
	var data models.EmTextbook

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmTextbookService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmTextbook对象
func (e *EmTextbook) Get(d *dto.EmTextbookGetReq, p *actions.DataPermission, model *models.EmTextbook) error {
	var data models.EmTextbook

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmTextbook error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmTextbook对象
func (e *EmTextbook) Insert(c *dto.EmTextbookInsertReq) error {
    var err error
    var data models.EmTextbook
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmTextbookService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmTextbook对象
func (e *EmTextbook) Update(c *dto.EmTextbookUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.EmTextbook{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("EmTextbookService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除EmTextbook
func (e *EmTextbook) Remove(d *dto.EmTextbookDeleteReq, p *actions.DataPermission) error {
	var data models.EmTextbook

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEmTextbook error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
