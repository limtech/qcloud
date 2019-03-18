package qcloud

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/limtech/utils"
)

const (
	QCLOUD_SMS_SEND_URL string = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=%s&random=%s"
)

type Sms struct {
	AppID  string
	AppKey string
	Random string

	Config SmsConfig
	Result SmsResult
}

type SmsConfig struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    struct {
		Mobile     string `json:"mobile"`
		Nationcode string `json:"nationcode"`
	} `json:"tel"`
	Time  int64 `json:"time"`
	TplId int64 `json:"tpl_id"`
}

type SmsResult struct {
	Result int64  `json:"result"`
	ErrMsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int64  `json:"fee"`
	Sid    string `json:"sid"`
}

func NewSms(appid string, appkey string) *Sms {
	sms := &Sms{}
	sms.AppID = appid
	sms.AppKey = appkey
	sms.Random = utils.RandomString(8, 3)
	sms.Config.Time = time.Now().Unix()
	sms.Config.Tel.Nationcode = "86"
	return sms
}

func (self *Sms) Send(mobile string, params []string, tplId int64, sign string, nationCode string) (SmsResult, error) {
	url := fmt.Sprintf(QCLOUD_SMS_SEND_URL, self.AppID, self.Random)
	// make signature
	signature := fmt.Sprintf("%x", sha256.Sum256([]byte(
		fmt.Sprintf(
			"appkey=%s&random=%s&time=%d&mobile=%s",
			self.AppKey,
			self.Random,
			self.Config.Time,
			mobile,
		),
	)))
	self.Config.Sig = signature
	self.Config.Params = params
	self.Config.Tel.Mobile = mobile
	if nationCode != "" {
		self.Config.Tel.Nationcode = nationCode
	}
	self.Config.TplId = tplId

	// http post
	content, err := utils.HttpPostJson(url, self.Config, nil)
	if err != nil {
		return self.Result, err
	}

	err = json.Unmarshal(content, &self.Result)
	return self.Result, err
}
