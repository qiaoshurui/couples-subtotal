// description:
// @author renshiwei
// Date: 2022/10/5 17:57

package initialize

import (
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"sync"
)

var once sync.Once

func InitServer(configDir string) {
	once.Do(func() {
		InitConfig(configDir)
		InitGorm()
		InitTencentCloud()
		logger.InitLog()
		InitGin()
	})
}
