package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间降序排列
	}

	query := DB.Where(model)

	// 统计总数
	countQuery := DB.Model(&model).Where(model)
	countQuery.Count(&count)

	// 如果未指定分页，则返回全部数据
	if option.Limit == 0 {
		err = query.Order(option.Sort).Find(&list).Error
		return list, count, err
	}

	// 分页查询
	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit
	if offset < 0 {
		offset = 0
	}

	err = query.Limit(option.PageInfo.Limit).
		Offset(offset).
		Order(option.Sort).
		Find(&list).Error

	return list, count, err
}
