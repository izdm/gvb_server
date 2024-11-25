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
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间往前面排序
	}

	count = DB.Debug().Select("id").Find(&list).RowsAffected

	//如果传的值
	if option.Limit == 0 {
		err = DB.Debug().Order(option.Sort).Find(&list).Error
		return list, count, err
	}

	offset := (option.PageInfo.Page - 1) * option.PageInfo.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Debug().Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
