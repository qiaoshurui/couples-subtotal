// description:
// @author renshiwei
// Date: 2022/10/5 22:00

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	examples "github.com/qiaoshurui/couples-subtotal/app/examples/router"

	couples "github.com/qiaoshurui/couples-subtotal/app/couples/router"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/qiaoshurui/couples-subtotal/common/initialize"
	_ "github.com/qiaoshurui/couples-subtotal/docs"
	"github.com/spf13/cobra"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "couples-cli server -c config/config.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yml", "Start server with provided configuration file")
}

func setup() {
	initialize.InitServer(configYml)

	usageStr := `starting api server...`
	log.Println(usageStr)
}

func run() error {
	return initRouter()
}

func initRouter() error {
	global.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	examples.InitHelloRouter(global.Router)
	couples.InitCoupleRouter(global.Router)
	global.Router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := global.Router.Run(fmt.Sprintf("%s:%d", global.Config.Server.Http.Host, global.Config.Server.Http.Port))
	if err != nil {
		return errors.Wrap(err, "Failed to start http server")
	}

	return nil
}
