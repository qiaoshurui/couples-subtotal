// description:
// @author renshiwei
// Date: 2022/10/6 14:36

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/cmd/api"
	"github.com/qiaoshurui/couples-subtotal/cmd/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "couples-cli",
	Short:        "couples-cli",
	SilenceUsage: true,
	Long:         `couples-cli:https://github.com/qiaoshurui/couples-subtotal`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + `couples-subtotal:` + ` 可以使用 ` + `-h` + ` 查看命令`
	usageStr1 := `也可以参考 https://github.com/qiaoshurui/couples-subtotal 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(api.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
