package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
	"log"
)

func SettingPost(path string, email string) (err error) {
	db := sqlConnect()
	defer db.Close()

	var tmp models.Setting
	if err = db.Where("email = ?", email).First(&tmp).Error; err == nil {
		log.Println("SettingPost: Setting already exists")
		err = db.Model(models.Setting{}).Where("email = ?", email).Select("path").Updates(models.Setting{Path: path}).Error
	} else {
		log.Println("SettingPost: Setting does not exist")
		var s models.Setting = models.Setting{Path: path, Email: email}
		err = db.Create(&s).Error
	}

	return err
}

func SettingGet(email string) (s models.Setting, err error) {
	db := sqlConnect()
	defer db.Close()

	err = db.Where("email = ?", email).First(&s).Error

	return s, err
}
