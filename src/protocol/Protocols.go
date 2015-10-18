package protocol
import (
	"fmt"
	"net/url"
	"io/ioutil"
	"regexp"
	"bytes"
	"strings"
)

var (
	httpClient *HttpClient
)

const (
	APP_ID = "717054801"
	S_URL = "http://mp.qq.com/pre/login.php"
	UIN = "751254754"
	ACTION = "4-22-1445154711643"
	PASSWORD = ""
)

func init() {
	httpClient = NewHttpClient()
}

func XLogin(username string, password string, imageCode string) {

	url := "http://xui.ptlogin2.qq.com/cgi-bin/xlogin?appid=" + APP_ID + "&daid=296&s_url=" + S_URL + "&style=33&hide_title_bar=1&fontcolor=ffffff&enable_qlogin=0&self_regurl=http://zc.qq.com/chs/index.html"
	fmt.Println(url)
	resp, err := httpClient.client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}else {
		//		cookies := httpClient.client.Jar.Cookies(resp.Request.URL)
		cookies := resp.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "pt_login_sig" {
				Check(UIN, cookie.Value)
			}
		}
	}
}

func Check(uin string, login_sig string) {
	s_url := url.QueryEscape(S_URL)
	url := "http://check.ptlogin2.qq.com/check?regmaster=&pt_tea=1&pt_vcode=1&uin=" + uin + "&appid=" + APP_ID + "&js_ver=10136&js_type=1&login_sig=" + login_sig +"&u1=" + s_url + "&r=0.9126439963001758"
	fmt.Println(s_url)
	resp, err := httpClient.client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	}else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		respStr := string(data)
		reg := regexp.MustCompile("'(.*?)'")
		v := reg.FindAllString(respStr, -1)

		vcode := v[1]
		uin := v[2]
		if v[0] != "0" {
			vcode = getVerifyCode(vcode)
		}
		Login(uin, vcode, PASSWORD, login_sig)

	}
}

func Login(uin  string, vcode string, password string, login_sig string) {
	urlBuffer := bytes.Buffer{}
	urlBuffer.WriteString("http://ptlogin2.qq.com/login?");
	urlBuffer.WriteString("u=" + uin)
	urlBuffer.WriteString("&verifycode=" + vcode)
	urlBuffer.WriteString("&pt_vcode_v1=0")
	urlBuffer.WriteString("&pt_verifysession_v1=")
	urlBuffer.WriteString("&p=" + pwdEncode(vcode, uin, password))
	urlBuffer.WriteString("&pt_randsalt=0")
	urlBuffer.WriteString("&u1=" + url.QueryEscape(S_URL))
	urlBuffer.WriteString("&ptredirect=1")
	urlBuffer.WriteString("&h=1")
	urlBuffer.WriteString("&ptlang=2052")
	urlBuffer.WriteString("&daid=296")
	urlBuffer.WriteString("&from_ui=1")
	urlBuffer.WriteString("&action=loginerroralert")
	urlBuffer.WriteString("&from_ui=1")
	urlBuffer.WriteString("&action=" + ACTION)
	urlBuffer.WriteString("&js_vsr=10136")
	urlBuffer.WriteString("&js_type=1")
	urlBuffer.WriteString("&login_sig=" + login_sig)
	urlBuffer.WriteString("&pt_uistyle=33")
	urlBuffer.WriteString("&aid=717054801&")

	fmt.Println(urlBuffer.String())
}

func pwdEncode(vcode string, uin string, password string) string {
	salt := strings.Replace(uin, "\\x", "", -1)
	fmt.Println(salt)

	return ""
}

func getVerifyCode(vcode string) string {
	return vcode
}