package qcloud

// for more info go to https://007.qq.com/product.html

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/limtech/utils"
)

const (
	QCLOUD_CAPTCHA_VERIFY_URL string = "https://ssl.captcha.qq.com/ticket/verify?aid=%s&AppSecretKey=%s&Ticket=%s&Randstr=%s&UserIP=%s"
)

type (
	Captcha struct {
		Aid     string
		Key     string
		Ticket  string
		Randstr string
		UserIP  string
	}
	CaptchaVerifyResult struct {
		Response  string `json:"response"`   // 1:验证成功，0:验证失败，100:AppSecretKey参数校验错误[required]
		EvilLevel string `json:"evil_level"` // [0,100]，恶意等级[optional]
		ErrMsg    string `json:"err_msg"`    // err msg 见下表
	}

	// 错误信息              | 详细说明                           | 错误信息                    | 详细说明
	// OK                   | 验证通过                           | cmd no match              | 验证码系统命令号不匹配
	// user code len error  | 验证码长度不匹配                    | uin no match              | 号码不匹配
	// captcha no match     | 验证码答案不匹配/Randstr参数不匹配    | seq redirect              | 重定向验证
	// verify timeout       | 验证码签名超时                      | opt no vcode              | 操作使用pt免验证码校验错误
	// Sequnce repeat       | 验证码签名重放	                  | diff                      | 差别，验证错误
	// Sequnce invalid      | 验证码签名序列	                  | captcha type not match    | 验证码类型与拉取时不一致
	// Cookie invalid       | 验证码cookie信息不合法               | verify type error        | 验证类型错误
	// verify ip no match   | ip不匹配                           | invalid pkg               | 非法请求包
	// decrypt fail	        | 验证码签名解密失败                   | bad visitor               | 策略拦截
	// appid no match	    | 验证码强校验appid错误                | system busy               | 系统内部错误
	// param err	        | AppSecretKey参数校验错误            |

)

func NewCaptcha(aid, key string) *Captcha {
	return &Captcha{
		Aid: aid,
		Key: key,
	}
}

func (self *Captcha) Verify(randstr, ticket, ip string) (bool, error) {
	self.Randstr = randstr
	self.Ticket = ticket
	self.UserIP = ip
	url := fmt.Sprintf(QCLOUD_CAPTCHA_VERIFY_URL, self.Aid, self.Key, self.Ticket, self.Randstr, self.UserIP)

	rs := CaptchaVerifyResult{}
	// http request
	content, err := utils.HttpGet(url)
	if err != nil {
		return false, err
	}

	// json decode
	if err := json.Unmarshal(content, &rs); err != nil {
		return false, err
	}

	// response
	if rs.Response != "1" {
		return false, errors.New(rs.ErrMsg)
	}
	return true, nil
}
