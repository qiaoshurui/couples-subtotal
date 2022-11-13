// description:
// @author renshiwei
// Date: 2022/10/5 22:42

package hello

import (
	"fmt"
	"github.com/qiaoshurui/couples-subtotal/app/examples/service"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 目前有问题
func TestPingRoute(t *testing.T) {
	//initialize.InitServer()
	//api.Run()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	global.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
}

func TestSayHello(t *testing.T) {
	helloService := new(service.Hello)
	helloStr, err := helloService.SayHello("xiaoqiao")
	assert.Nil(t, err)
	fmt.Println(helloStr)
}
