package couples

import (
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddDynamic(t *testing.T) {
	initialize.InitServer("../../../config/config-dev.yaml")
	dynamicService := service.DynamicService{}
	data := &dto.AddDynamic{
		Content: "这是使用测试用例编写的动态",
		Status:  2,
	}
	err := dynamicService.AddDynamic(data)
	assert.NoError(t, err, "动态添加失败")
}
