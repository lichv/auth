package services

import (
	"auth/app/models"
	"fmt"
)

type WechatUser struct {
	Code       string `json:"code" form:"code"`
	ConfigCode string `json:"config_code" form:"config_code"`
	Openid     string `json:"openid" form:"openid"`
	Unionid    string `json:"unionid" form:"unionid"`
	Nickname   string `json:"nickname" form:"nickname"`
	Sex        string `json:"sex" form:"sex"`
	Headimage  string `json:"headimage" form:"headimage"`
	Country    string `json:"country" form:"country"`
	Province   string `json:"province" form:"province"`
	City       string `json:"city" form:"city"`
	Phone      string `json:"phone" form:"phone"`
	Privilege  string `json:"privilege" form:"privilege"`
	FLag       string `json:"flag" form:"flag"`
	State      string `json:"state" form:"state"`
}

func ExistWechatUserByCode(code string) (b bool, err error) {
	b, err = models.ExistWechatUserByCode(code)
	return b, err
}

func GetWechatUserTotal(maps interface{}) (count int, err error) {
	count, err = models.GetWechatUserTotal(map[string]interface{}{})
	return count, err
}
func GetWechatUserOne(query map[string]interface{}, orderBy interface{}) (WechatUser *WechatUser, err error) {
	var nu *models.WechatUser
	nu, err = models.GetWechatUserOne(query, orderBy)
	return TransferWechatUserModel(nu), nil
}

func GetWechatUserPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) (WechatUsers []*WechatUser, total int, errs []error) {
	count, err := models.GetWechatUserTotal(query)
	fmt.Println(count)
	if err != nil {
		return nil, 0, errs
	}
	us, errs := models.GetWechatUserPages(query, orderBy, pageNum, pageSize)
	WechatUsers = TransferWechatUsers(us)
	return WechatUsers, total, nil
}

func AddWechatUser(data map[string]interface{}) (err error) {
	err = models.AddWechatUser(data)
	return err
}

func EditWechatUser(code string, data map[string]interface{}) (err error) {
	err = models.EditWechatUser(code, data)
	return err
}

func DeleteWechatUser(maps map[string]interface{}) (err error) {
	err = models.DeleteWechatUsers(maps)
	return nil
}

func ClearAllWechatUser() (err error) {
	err = models.ClearAllWechatUser()
	return err
}

func TransferWechatUserModel(wu *models.WechatUser) (wechatUser *WechatUser) {
	wechatUser = &WechatUser{
		Code:       wu.Code,
		ConfigCode: wu.ConfigCode,
		Openid:     wu.Openid,
		Unionid:    wu.Unionid,
		Nickname:   wu.Nickname,
		Sex:        wu.Sex,
		Headimage:  wu.Headimage,
		Country:    wu.Country,
		Province:   wu.Province,
		City:       wu.City,
		Phone:      wu.Phone,
		Privilege:  wu.Privilege,
		FLag:       wu.FLag,
		State:      wu.State,
	}
	return
}
func TransferWechatUsers(us []*models.WechatUser) (WechatUsers []*WechatUser) {
	for _, value := range us {
		WechatUser := TransferWechatUserModel(value)
		WechatUsers = append(WechatUsers, WechatUser)
	}
	return WechatUsers
}
