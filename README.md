
# QCloud SDK for golang

### send sms
```go
import "github.com/limtech/qcloud"

func main {
    QCloudSDKAppID := ""
    QCloudAppKey := ""
    mobile := "13912345678"

	sms := qcloud.NewSms(QCloudSDKAppID, QCloudAppKey)
	action, err := sms.Send(
		mobile,
		[]string{"注册", "1234", "10"}, // 您的{1}验证码是{2}，请于{3}分钟内填写。如非本人操作，请忽略本短信。
		101010,
		"LIM.TECH",
		"86",
	)
}

type Sms struct{ ... }
    func NewSms(appid string, appkey string) *Sms
type SmsConfig struct{ ... }
type SmsResult struct{ ... }
```

### captcha

```go
import "github.com/limtech/qcloud"

const (
	QcloudCaptchaAid string = "1234567890" // change to your own
	QcloudCaptchaKey string = "xxxxxxxxxx" // change to your own
)

func main {
	Randstr := "from front end"
	Ticket := "from front end"
	ClientIP := "client IP"

	// do qcloud captcha verify
	captcha := qcloud.NewCaptcha(QcloudCaptchaAid, QcloudCaptchaKey)
	if ok, err := captcha.Verify(Randstr, Ticket, ClientIP); err != nil || !ok {
		log.Println(err)
		return
	}
}
```