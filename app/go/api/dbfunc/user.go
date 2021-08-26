package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
	"errors"

	"golang.org/x/crypto/bcrypt" //パスワードをハッシュ化するために使用
)

func GetUserInfo(email string) (user models.User, err error) {
	db := sqlConnect()

	// データベースからユーザー情報を取得
	var u models.User
	err = db.Where("email = ?", email).First(&u).Error

	defer db.Close()
	return u, err
}

func PostUser(email string, password string, name string) (err error) {
	db := sqlConnect()

	// おなじemailが既に登録されていないか確認
	var tmp models.User
	if err = db.Where("email = ?", email).First(&tmp).Error; err == nil {
		return errors.New("it already exists")
	}

	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashStr := string(hash)

	// データベースに登録
	var u models.User = models.User{Email: email, Password: hashStr, Name: name, Coin: 100}
	err = db.Create(&u).Error

	defer db.Close()
	return err
}

func PutUser(email string, password string, name string) (err error) {
	db := sqlConnect()

	// データベースに存在しているか確認
	var u models.User
	if err = db.Where("email = ?", email).First(&u).Error; err != nil {
		return err
	}
	// パスワードを確かめる
	if !ComparePassword(u.Password, password) {
		return errors.New("wrong password")
	}

	// データベースを更新
	err = db.Model(models.User{}).Where("email = ?", email).Select("email", "name").Updates(models.User{Email: email, Name: name}).Error

	defer db.Close()
	return err
}

// ComparePassword パスワードを確認する
// args:
//   hashStr ハッシュ化されたデータベースに登録されているパスワード
//   inputPass 送られてきたパスワード
func ComparePassword(hashStr string, inputPass string) (Ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(inputPass))
	if err == nil {
		Ok = true
	} else {
		Ok = false
	}
	return Ok
}
