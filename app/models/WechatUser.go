package models

import (
	"auth/utils"
	"github.com/jinzhu/gorm"
)

type WechatUser struct {
	Code string `json:"code" form:"code"`
	ConfigCode string `json:"config_code" form:"config_code"`
	Openid string `json:"openid" form:"openid"`
	Unionid string `json:"unionid" form:"unionid"`
	Nickname string `json:"nickname" form:"nickname"`
	Sex string `json:"sex" form:"sex"`
	Headimage string `json:"headimage" form:"headimage"`
	Country string `json:"country" form:"country"`
	Province string `json:"province" form:"province"`
	City string `json:"city" form:"city"`
	Phone string `json:"phone" form:"phone"`
	Privilege string `json:"privilege" form:"privilege"`
	FLag string `json:"flag" form:"flag"`
	State string `json:"state" form:"state"`
}

func ExistWechatUserByCode(code string) (b bool,err error) {
	var wechatUser WechatUser
	err = db.Model(&WechatUser{}).Select("code").Where("code = ? ",code).First(&wechatUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false,err
	}
	return true, err
}

func GetWechatUserTotal(maps interface{}) (count int,err error) {
	err = db.Model(&WechatUser{}).Where("state = ?",true).Count(&count).Error
	if err != nil {
		return 0,err
	}
	return count, err
}

func FindWechatUserByCode( code string) ( wechatUser *WechatUser, err error) {
	err = db.Model(&WechatUser{}).Where("code = ? ",code).First(&wechatUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &WechatUser{},err
	}
	return
}

func GetWechatUserOne( query map[string]interface{},orderBy interface{}) ( *WechatUser,error) {
	var wechatUser WechatUser
	model := db.Model(&WechatUser{})
	for key, value := range query {
		b,err := utils.In ([]string{"code", "WechatUsername", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatUser", "flag", "state"},key)
		if  err != nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	err := model.First(&wechatUser).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return &WechatUser{},nil
	}
	return &wechatUser, nil
}

func GetWechatUserPages( query map[string]interface{},orderBy interface{},pageNum int,pageSize int) ( []*WechatUser, []error) {
	var WechatUsers []*WechatUser
	var errs []error
	model := db.Where("state=?",true)
	for key, value := range query {
		b,err := utils.In ([]string{"code", "WechatUsername", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatUser", "flag", "state"},key)
		if  err != nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	model =model.Offset(pageNum).Limit(pageSize).Order(orderBy)
	model = model.Find(&WechatUsers)
	errs = model.GetErrors()
	//err = model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&WechatUsers).Error

	return WechatUsers, errs
}

func AddWechatUser( data map[string]interface{}) error {
	WechatUser := WechatUser{
		Code:data["Code"].(string),
		ConfigCode:data["ConfigCode"].(string),
		Openid:data["Openid"].(string),
		Unionid:data["Unionid"].(string),
		Nickname:data["Nickname"].(string),
		Sex:data["Sex"].(string),
		Headimage:data["Headimage"].(string),
		Country:data["Country"].(string),
		Province:data["Province"].(string),
		City:data["City"].(string),
		Phone:data["Phone"].(string),
		Privilege:data["Privilege"].(string),
		FLag:data["FLag"].(string),
		State:data["State"].(string),
	}
	if err:= db.Create(&WechatUser).Error;err != nil{
		return err
	}
	return nil
}

func EditWechatUser( code string,data map[string]interface{}) error {
	if err:= db.Model(&WechatUser{}).Where("code=?",code).Updates(data).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWechatUser(code string) error {
	if err := db.Where("code=?",code).Delete(WechatUser{}).Error;err != nil{
		return err
	}
	return nil
}

func DeleteWechatUsers(maps map[string]interface{}) error{
	model := db
	for key, value := range maps {
		b,err := utils.In ([]string{"code", "WechatUsername", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatUser", "flag", "state"},key)
		if  err != nil && b{
			model = model.Where(key + "= ?", value)
		}
	}
	if err :=model.Delete(&WechatUser{}).Error;err != nil{
		return err
	}
	return nil
}

func ClearAllWechatUser() error {
	if err := db.Unscoped().Delete(&WechatUser{}).Error; err != nil {
		return err
	}
	return nil
}
