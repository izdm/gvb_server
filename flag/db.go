package flag

import (
	"gvb_server/global"
	"gvb_server/models"
)

func MakeMigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.User2Collects{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.BannerModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error(" [ error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info(" [ success ] 生成数据库表结构成功!")
}
