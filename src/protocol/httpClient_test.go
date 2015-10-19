package protocol
import (
	"testing"
	"strings"
	"fmt"
	"encoding/hex"
	"encoding/base64"
)


func Test_Login(t *testing.T){
//	PwdEncode2("","\\x00\\x00\\x00\\x00\\x2e\\x42\\x14\\x7d","")

	XLogin("776082557","a123789")

//	PwdEncode("776082557","a123789","!FOX")

//	Hexchar2bin("c4ca4238a0b923820dcc509a6f75849b")
}



func PwdEncode2(vcode string, uin string, password string) string {
	salt := strings.Replace(uin,"\\x","",-1)
	data,err := hex.DecodeString(salt)
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Println(string(data))
	}
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

func Test_ha(t *testing.T)  {
	str := "8gzgC65TYfj6OunO+klTYv99obpij2SjR/CowBK/CyVKMM2Sq//npu4NxCTLYWb4gZ76W8yyDt+0rQLkEsz1ebHKcR1VuLCzrrYBU9XgaToqhvMWfXhHoMuLAABHFqkJXZutyXfLuATb3LpgKalxCGmkU/J9/d+DwBbZKLPL9Mc="

	data,_ := base64.StdEncoding.DecodeString(str)
	fmt.Println(len(data))

	publicKeyStr := "F20CE00BAE5361F8FA3AE9CEFA495362FF7DA1BA628F64A347F0A8C012BF0B254A30CD92ABFFE7A6EE0DC424CB6166F8819EFA5BCCB20EDFB4AD02E412CCF579B1CA711D55B8B0B3AEB60153D5E0693A2A86F3167D7847A0CB8B00004716A9095D9BADC977CBB804DBDCBA6029A9710869A453F27DFDDF83C016D928B3CBF4C7"

	data,_ = hex.DecodeString(publicKeyStr)
	fmt.Println(len(data),len(publicKeyStr))

	base64Str := base64.StdEncoding.EncodeToString(data);
	fmt.Println(base64Str)
}