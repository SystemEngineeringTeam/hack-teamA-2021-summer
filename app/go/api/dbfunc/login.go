package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
)

func PostLoginInfo(email string, password string) (user models.User,err error){
	db := sqlConnect()

	// データベースからユーザー情報を取得
	var u models.User
	if err = db.Where("email = ?", email).First(&u).Error; err != nil {
		return u,err
	}
	
	
	defer db.Close()
	return u,err
}
