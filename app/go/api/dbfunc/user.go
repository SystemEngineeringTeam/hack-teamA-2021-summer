package dbfunc

import (
	"SystemEngineeringTeam/hack-teamA-2021-summer/models"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt" //パスワードをハッシュ化するために使用
)

func GetUserInfo(c echo.Context) (user models.User, err error) {
	db := sqlConnect()
	defer db.Close()

	user, err = GetUserFromToken(c)
	if err != nil {
		return user, err
	}

	return user, err //接続できなかったり、データがないときはエラーを出す
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
	var u models.User = models.User{Email: email, Password: hashStr, Name: name, Coin: 100, SpinCount: 0, UUID: uid.String()}
	err = db.Create(&u).Error
	var c models.ChartData = models.ChartData{UserUID: uid.String(), Spin: 0, Coin: 100}
	err = db.Create(&c).Error
	var imgPath []string
	files, err := ioutil.ReadDir("./static/img")
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			images, err := ioutil.ReadDir("./static/img/" + file.Name())
			if err != nil {
				return err
			}
			imgPath = append(imgPath, images[0].Name())
		}
	}
	Path := strings.Join(imgPath, ":")
	var s models.Setting = models.Setting{UserUID: uid.String(), Path: Path}
	err = db.Create(&s).Error

	return err
}

func PutUser(c echo.Context, name string) (err error) {
	db := sqlConnect()
	defer db.Close()

	// データベースに存在しているか確認
	user, err := GetUserFromToken(c)
	if err != nil {
		return err
	}

	// データベースを更新
	err = db.Model(models.User{}).Where("uuid = ?", user.UUID).Select("name").Updates(models.User{Name: name}).Error

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

func GetUserFromToken(c echo.Context) (u models.User, err error) {
	db := sqlConnect()
	defer db.Close()

	// トークンを取得
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	uid := claims["uid"].(string)

	// データベースからユーザー情報を取得
	err = db.Where("uuid = ?", uid).First(&u).Error

	return u, err
}
