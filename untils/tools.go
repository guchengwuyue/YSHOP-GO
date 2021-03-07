package untils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strings"
	"yixiang.co/yshop/controllers"
)

//加密
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd,bcrypt.MinCost)
	if err != nil {
		controllers.ErrMsg(err.Error())
	}
	return string(hash)
}

//密码验证
func ComparePwd(hashPwd string,plainPwd []byte) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash,plainPwd)
	if err != nil {
		return false
	}

	return true
}

//判断array contain item
func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice: {
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	}
	return
}


//[a] -> a -> a
//[a b c] -> a b c -> a,b,c
func Convert(array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func IntToBool (num int8) bool {
	if num > 0 {
		return true
	}
	return false
}

func  ReturnQ(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += ",?"
	}
	return str[1:]
}

