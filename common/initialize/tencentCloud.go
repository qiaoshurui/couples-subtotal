package initialize

import (
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func InitTencentCloud() {
	u, _ := url.Parse("https://couples-subtotal-1300745270.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.Config.TencentCloud.SecretId,
			SecretKey: global.Config.TencentCloud.SecretKey,
		},
	})
	global.Client = client
}
