package couples

import (
	"fmt"
	"github.com/qiaoshurui/couples-subtotal/common/utils"
	"testing"
)

func TestCode(t *testing.T) {
	code := "123456"
	encryptionCode := utils.PasswordEncryption(code)
	decryptionCode := utils.PasswordDecryption(encryptionCode)
	fmt.Println(decryptionCode)
}
