package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
	"yixiang.co/yshop/common"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/redis"
	"yixiang.co/yshop/vo"
)

type userStdClaims struct {
	vo.JwtUser
	//*models.User
	jwt.StandardClaims
}

var (
	verifyKey  string
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
	ErrExpired = "token expired" // 令牌过期
	ErrOther   = "other error"   // 其他错误
)

const bearerLength = len("Bearer ")

func init()  {
	verifyKey,_ = beego.AppConfig.String("jwt_token")
}

func GenerateToken(m *models.User,d time.Duration) (string,error) {
	m.Password = ""
	//m.Permissions = []string{}
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Id: strconv.FormatInt(m.Id,10),
		Issuer:    "YshopGo",
	}

	var jwtUser = vo.JwtUser{
		Id: m.Id,
		Avatar: m.Avatar,
		Email: m.Email,
		Username: m.Username,
		Phone: m.Phone,
		NickName: m.NickName,
		Sex: m.Sex,
		Dept: m.Depts.Name,
		Job: m.Jobs.Name,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		JwtUser:  jwtUser,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,uClaims)
	tokenString,err := token.SignedString([]byte(verifyKey))
	if err != nil {
		logs.Info(err)
	}
	//set redis
	var key = common.REDIS_PREFIX_AUTH + tokenString
	json, _ := json.Marshal(m)
	redis.SetEx(key,string(json),expireTime.Unix())

	return tokenString,err
}

func ValidateToken(tokenString string) (*vo.JwtUser,error)  {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(verifyKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtUser, err

}


//返回id
func GetAdminUserId(c *context.BeegoInput) (int64, error) {
	u := c.GetData(common.ContextKeyUserObj)
	user, ok := u.(*vo.JwtUser)
	if ok {
		return user.Id, nil
	}
	return 0,errors.New("can't convert to user struct")
}

//返回user
func GetAdminUser(c *context.BeegoInput) (*vo.JwtUser, error) {
	u := c.GetData(common.ContextKeyUserObj)
	user, ok := u.(*vo.JwtUser)
	if ok {
		return user, nil
	}
	return nil,errors.New("can't convert to user struct")
}

//返回 detail user
func GetAdminDetailUser(c *context.BeegoInput) *models.User {
	mytoken := c.Header("Authorization")
	token := strings.TrimSpace(mytoken[bearerLength:])
	var key = common.REDIS_PREFIX_AUTH + token
	userMap, _:= redis.Get(key)
	jsonStr := userMap[key]
	user := &models.User{}
	json.Unmarshal([]byte(jsonStr),user)
	return user
}

func RemoveUser(c *context.BeegoInput) error{
	mytoken := c.Header("Authorization")
	token := strings.TrimSpace(mytoken[bearerLength:])
	var key = common.REDIS_PREFIX_AUTH + token
	return redis.Del(key)
}
