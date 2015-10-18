package protocol
import (
	"testing"
	"strings"
	"fmt"
)


func Test_Login(t *testing.T){
	PwdEncode2("","\\x00\\x00\\x00\\x00\\x2e\\x42\\x14\\x7d","")
}

func PwdEncode2(vcode string, uin string, password string) string {
	salt := strings.Replace(uin,"\\x","",-1)
	fmt.Println(salt)
	return ""
}

func Test_PtuiCheckVC(t *testing.T) {
//	respStr := "ptui_checkVC('0','!UIF','\\x00\\x00\\x00\\x00\\x2e\\x42\\x14\\x7d','46545c20d11742f3caaeadb7780f56f815407858305f80d7f1a035d2e855f8104f8b348195d7b812bf8d722b9ad7f6c25b2b6ee73e6f4559','0')"
//	reg := regexp.MustCompile("'(.*?)'")
//
//	v := reg.FindAllString(respStr,-1)
//
//	fmt.Println(v[2])
}