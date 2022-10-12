// description:
// @author renshiwei
// Date: 2022/10/13 00:47

package couples

import (
	"fmt"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCouplesInfo(t *testing.T) {
	initialize.InitServer("../../../config/config-dev.yaml")

	relationship := service.Relationship{}
	couplesInfo, err := relationship.GetRelationship(1)
	assert.NoError(t, err)
	fmt.Println(couplesInfo)
}
