package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt" //パスワードをハッシュ化するために使用
)

func GetUserInfo(email string) (user models.User, err error) {
	db := sqlConnect()
	defer db.Close()

	// データベースからユーザー情報を取得
	var u models.User
	err = db.Where("email = ?", email).First(&u).Error //e-mailを元にユーザー情報を取得

	return u, err //接続できなかったり、データがないときはエラーを出す
}

func PostUser(email string, password string, name string) (err error) {
	db := sqlConnect()
	defer db.Close()

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

	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	// データベースに登録
	var u models.User = models.User{Email: email, Password: hashStr, Name: name, Coin: 100, UUID: uid.String()}
	err = db.Create(&u).Error

	return err
}

func PutUser(email string, password string, name string) (err error) {
	db := sqlConnect()
	defer db.Close()

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
