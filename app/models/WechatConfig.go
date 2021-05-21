package models

import (
	"auth/utils"
	"github.com/jinzhu/gorm"
)

type WechatConfig struct {
	Code            string      `json:"code" form:"code" gorm:"code"`
	Type            string      `json:"type" form:"type" gorm:"type"`
	Appid           string      `json:"appid" form:"appid" gorm:"appid"`
	Appsecret       string      `json:"appsecret" form:"appsecret" gorm:"appsecret"`
	Scope           string      `json:"scope" form:"scope" gorm:"scope"`
	AuthRedirectUrl string      `json:"auth_redirect_url" form:"auth_redirect_url" gorm:"auth_redirect_url"`
	NoticeUrl       string      `json:"notice_url" form:"notice_url" gorm:"notice_url"`
	Group           string      `json:"group" form:"group" gorm:"group"`
	Company         string      `json:"company" form:"company" gorm:"company"`
	FLag            string      `json:"flag" form:"flag" gorm:"flag"`
	State           string      `json:"state" form:"state" gorm:"state"`
}

func ExistWechatConfigByCode(code string) (b bool, err error) {
	var wechatConfig WechatConfig
	err = db.Model(&WechatConfig{}).Select("code").Where("code = ? ", code).First(&wechatConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, err
}

func GetWechatConfigTotal(maps interface{}) (count int, err error) {
	err = db.Model(&WechatConfig{}).Where("state = ?", true).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func FindWechatConfigByCode(code string) ( *WechatConfig, error) {
	var wechatConfig WechatConfig
	err := db.Where("code = ? ", code).First(&wechatConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &WechatConfig{}, err
	}
	return &wechatConfig,nil
}

func GetWechatConfigOne(query map[string]interface{}, orderBy interface{}) (*WechatConfig, error) {
	var wechatConfig WechatConfig
	model := db.Model(&WechatConfig{})
	for key, value := range query {
		b, err := utils.In([]string{"code", "WechatConfigname", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatConfig", "flag", "state"}, key)
		if err != nil && b {
			model = model.Where(key+"= ?", value)
		}
	}
	err := model.First(&wechatConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &WechatConfig{}, nil
	}
	return &wechatConfig, nil
}

func GetWechatConfigPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*WechatConfig, []error) {
	var wechatConfigs []*WechatConfig
	var errs []error
	model := db.Where("state=?", true)
	for key, value := range query {
		b, err := utils.In([]string{"code", "WechatConfigname", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatConfig", "flag", "state"}, key)
		if err != nil && b {
			model = model.Where(key+"= ?", value)
		}
	}
	model = model.Offset(pageNum).Limit(pageSize).Order(orderBy)
	model = model.Find(&wechatConfigs)
	errs = model.GetErrors()
	//err = model.Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&WechatConfigs).Error

	return wechatConfigs, errs
}

func AddWechatConfig(data map[string]interface{}) error {
	WechatConfig := WechatConfig{
		Code:            data["Code"].(string),
		Type:            data["Type"].(string),
		Appid:           data["Appid"].(string),
		Appsecret:       data["Appsecret"].(string),
		Scope:           data["Scope"].(string),
		AuthRedirectUrl: data["AuthRedirectUrl"].(string),
		NoticeUrl:       data["NoticeUrl"].(string),
		Group:           data["Group"].(string),
		Company:         data["Company"].(string),
		FLag:            data["FLag"].(string),
		State:           data["State"].(string),
	}
	if err := db.Create(&WechatConfig).Error; err != nil {
		return err
	}
	return nil
}

func EditWechatConfig(code string, data map[string]interface{}) error {
	if err := db.Model(&WechatConfig{}).Where("code=?", code).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteWechatConfig(code string) error {
	if err := db.Where("code=?", code).Delete(WechatConfig{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteWechatConfigs(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		b, err := utils.In([]string{"code", "WechatConfigname", "password", "name", "sex", "birthday", "phone", "email", "province", "city", "county", "address", "reference", "regtime", "remark", "is_active", "is_superWechatConfig", "flag", "state"}, key)
		if err != nil && b {
			model = model.Where(key+"= ?", value)
		}
	}
	if err := model.Delete(&WechatConfig{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllWechatConfig() error {
	if err := db.Unscoped().Delete(&WechatConfig{}).Error; err != nil {
		return err
	}
	return nil
}
