package couples

import (
	"fmt"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/initialize"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAddPhoto(t *testing.T) {
	initialize.InitServer("../../../config/config-dev.yaml")
	photo := service.Photo{}
	open, err := os.Open("D:\\img\\22pn2w2xb04fmsyb884sjrlcp_0")
	if err != nil {
		fmt.Println("文件打开失败")
	}
	addPhoto, err := photo.UploadTencent("测试", 17, open)
	assert.NoError(t, err)
	fmt.Println(addPhoto)
}
