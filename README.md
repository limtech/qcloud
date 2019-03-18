
### qcloud sms sdk for golang

```go
import "github.com/limtech/qcloud"

func main {
    qCloudSDKAppID := ""
    qCloudAppKey := ""
    mobile := "13912345678"

	sms := qcloud.NewSms(qCloudSDKAppID, qCloudAppKey)
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