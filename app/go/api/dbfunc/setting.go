package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
	"log"

	"github.com/labstack/echo/v4"
)

func SettingPost(path string, c echo.Context) (err error) {
	db := sqlConnect()
	defer db.Close()

	var u models.User
	u, err = GetUserFromToken(c)
	if err != nil {
		return err
	}

	var tmp models.Setting
	if err = db.Where("user_uid = ?", u.UUID).First(&tmp).Error; err == nil {
		log.Println("SettingPost: Setting already exists")
		err = db.Model(models.Setting{}).Where("user_uid = ?", u.UUID).Select("path").Updates(models.Setting{Path: path}).Error
	} else {
		log.Println("SettingPost: Setting does not exist")
		var s models.Setting = models.Setting{Path: path, UserUID: u.UUID}
		err = db.Create(&s).Error
	}

	return err
}

func SettingGet(c echo.Context) (s models.Setting, err error) {
	db := sqlConnect()
	defer db.Close()

	var u models.User
	u, err = GetUserFromToken(c)
	if err != nil {
		return s, err
	}
	err = db.Where("user_uid = ?", u.UUID).First(&s).Error

	return s, err
}
