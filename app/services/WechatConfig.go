package services

import (
	"auth/app/models"
	"fmt"
)

type WechatConfig struct {
	Code            string      `json:"code" form:"code"`
	Type            string      `json:"type" form:"type"`
	Appid           string      `json:"appid" form:"appid"`
	Appsecret       string      `json:"appsecret" form:"appsecret"`
	Scope           string      `json:"scope" form:"scope"`
	AuthRedirectUrl string      `json:"auth_redirect_url" form:"auth_redirect_url"`
	NoticeUrl       string      `json:"notice_url" form:"notice_url"`
	Group           string      `json:"group" form:"group"`
	Company         string      `json:"company" form:"company"`
	FLag            string      `json:"flag" form:"flag"`
	State           string      `json:"state" form:"state"`
}

func ExistWechatConfigByCode(code string) (b bool, err error) {
	b, err = models.ExistWechatConfigByCode(code)
	return b, err
}

func FindWechatConfigOne(code string) (*WechatConfig,error) {
	wc,_ := models.FindWechatConfigByCode(code)
	return TransferWechatConfigModel(wc),nil
}
func GetWechatConfigTotal(maps interface{}) (count int, err error) {
	count, err = models.GetWechatConfigTotal(map[string]interface{}{})
	return count, err
}
func GetWechatConfigOne(query map[string]interface{}, orderBy interface{}) (WechatConfig *WechatConfig, err error) {
	var nu *models.WechatConfig
	nu, err = models.GetWechatConfigOne(query, orderBy)
	return TransferWechatConfigModel(nu), nil
}

func GetWechatConfigPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) (WechatConfigs []*WechatConfig, total int, errs []error) {
	count, err := models.GetWechatConfigTotal(query)
	fmt.Println(count)
	if err != nil {
		return nil, 0, errs
	}
	us, errs := models.GetWechatConfigPages(query, orderBy, pageNum, pageSize)
	WechatConfigs = TransferWechatConfigs(us)
	return WechatConfigs, total, nil
}

func AddWechatConfig(data map[string]interface{}) (err error) {
	err = models.AddWechatConfig(data)
	return err
}

func EditWechatConfig(code string, data map[string]interface{}) (err error) {
	err = models.EditWechatConfig(code, data)
	return err
}

func DeleteWechatConfig(maps map[string]interface{}) (err error) {
	err = models.DeleteWechatConfigs(maps)
	return nil
}

func ClearAllWechatConfig() (err error) {
	err = models.ClearAllWechatConfig()
	return err
}

func TransferWechatConfigModel(wc *models.WechatConfig) (wechatConfig *WechatConfig) {
	wechatConfig = &WechatConfig{
		Code:            wc.Code,
		Type:            wc.Type,
		Appid:           wc.Appid,
		Appsecret:       wc.Appsecret,
		Scope:           wc.Scope,
		AuthRedirectUrl: wc.AuthRedirectUrl,
		NoticeUrl:       wc.NoticeUrl,
		Group:           wc.Group,
		Company:         wc.Company,
		FLag:            wc.FLag,
		State:           wc.State,
	}
	return
}
func TransferWechatConfigs(us []*models.WechatConfig) (WechatConfigs []*WechatConfig) {
	for _, value := range us {
		WechatConfig := TransferWechatConfigModel(value)
		WechatConfigs = append(WechatConfigs, WechatConfig)
	}
	return WechatConfigs
}
