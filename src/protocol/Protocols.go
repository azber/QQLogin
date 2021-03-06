package protocol

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	httpClient *HttpClient
	cookies    []*http.Cookie
)

const (
	URL_IMAGE_CODE = "http://captcha.qq.com/getimgbysig"
	APP_ID = "776082557"
	S_URL = "http://mp.qq.com/pre/login.php"
	UIN = "751254754"
	ACTION = "4-22-1445154711643"
	PASSWORD = "a123789"
	PUBLIC_KEY = "F20CE00BAE5361F8FA3AE9CEFA495362FF7DA1BA628F64A347F0A8C012BF0B254A30CD92ABFFE7A6EE0DC424CB6166F8819EFA5BCCB20EDFB4AD02E412CCF579B1CA711D55B8B0B3AEB60153D5E0693A2A86F3167D7847A0CB8B00004716A9095D9BADC977CBB804DBDCBA6029A9710869A453F27DFDDF83C016D928B3CBF4C7"
)

func init() {
	httpClient = NewHttpClient()
}

func XLogin(username string, password string) {

	url := "http://xui.ptlogin2.qq.com/cgi-bin/xlogin?appid=" + APP_ID + "&daid=296&s_url=" + S_URL + "&style=33&hide_title_bar=1&fontcolor=ffffff&enable_qlogin=0&self_regurl=http://zc.qq.com/chs/index.html"
	fmt.Println(url)
	resp, err := httpClient.client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		//		cookies := httpClient.client.Jar.Cookies(resp.Request.URL)
		cookies = resp.Cookies()

		for _, cookie := range cookies {
			if cookie.Name == "pt_login_sig" {
				Check(UIN, cookie.Value)
			}
		}
	}
}

func Check(uin string, login_sig string) {
	s_url := url.QueryEscape(S_URL)
	url := "http://check.ptlogin2.qq.com/check?regmaster=&pt_tea=1&pt_vcode=1&uin=" + uin + "&appid=" + APP_ID + "&js_ver=10136&js_type=1&login_sig=" + login_sig + "&u1=" + s_url + "&r=0.9126439963001758"
	fmt.Println(s_url)
	resp, err := httpClient.client.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
		respStr := string(data)
		reg := regexp.MustCompile("'(.*?)'")
		v := reg.FindAllString(respStr, -1)

		vcode := v[1]
		//		uin := v[2]
		if v[0] != "0" {
			sig, _ := getVerifyCode(vcode)
			vcode = inputImageCode()
			vcode, sig, _ = cap_union_verify(UIN, vcode, sig)
			Login(uin, vcode, PASSWORD, sig)
		} else {
			Login(uin, vcode, PASSWORD, login_sig)
		}
	}
}

func Login(uin string, vcode string, password string, login_sig string) {
	urlBuffer := bytes.Buffer{}
	urlBuffer.WriteString("http://ptlogin2.qq.com/login?")
	urlBuffer.WriteString("u=" + UIN)
	urlBuffer.WriteString("&verifycode=" + vcode)
	urlBuffer.WriteString("&pt_vcode_v1=1")
	urlBuffer.WriteString("&pt_verifysession_v1=" + login_sig)
	urlBuffer.WriteString("&p=" + PwdEncode(UIN, password, vcode))
	urlBuffer.WriteString("&pt_randsalt=0")
	urlBuffer.WriteString("&u1=" + url.QueryEscape(S_URL))
	urlBuffer.WriteString("&ptredirect=1")
	urlBuffer.WriteString("&h=1")
	urlBuffer.WriteString("&t=1")
	urlBuffer.WriteString("&g=1")
	urlBuffer.WriteString("&ptlang=2052")
	urlBuffer.WriteString("&daid=296")
	urlBuffer.WriteString("&from_ui=1")
	urlBuffer.WriteString("&action=" + ACTION)
	urlBuffer.WriteString("&js_ver=10136")
	urlBuffer.WriteString("&js_type=1")
	urlBuffer.WriteString("&login_sig=")
	urlBuffer.WriteString("&pt_uistyle=33")
	urlBuffer.WriteString("&aid=717054801&")

	fmt.Println(urlBuffer.String())

	resq, err := httpClient.client.Get(urlBuffer.String())
	defer resq.Body.Close()
	if err != nil {
		panic(err)
		return
	}
	data, _ := ioutil.ReadAll(resq.Body)
	respStr := string(data)
	fmt.Println(respStr)

}

func PwdEncode(uin string, password string, vcode string) string {
	value, err := JsRun("Encryption().getEncryption('" + password + "', '" + uin + "', '" + vcode + "')")
	if err != nil {
		panic(err)
		return ""
	} else {
		//		fmt.Println(value)
		return value.String()
	}
}

func Hexchar2bin(hexStr string) []byte {
	stringBuffer := bytes.Buffer{}

	for i := 0; i < len(hexStr); i = i + 2 {
		//		stringBuffer.WriteString("")
		stringBuffer.WriteString(hexStr[i : i + 2])
	}

	data, err := hex.DecodeString(stringBuffer.String())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(data)
		return data
	}
	return nil
}

func callCap_union_show(uin string, vcode string) string {
	vcode = strings.Replace(vcode, "'", "", -1)
	url := "http://captcha.qq.com/cap_union_show?clientype=2&uin=" + uin + "&aid=" + uin + "&cap_cd=" + vcode
	fmt.Println(url)
	resp, err := httpClient.client.Get(url)
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		respStr := string(data)

		respStr = respStr[strings.Index(respStr, "var g_click_cap_sig=") + len("var g_click_cap_sig=") + 1:]
		respStr = respStr[:strings.Index(respStr, ";") - 1]
		return respStr
	}
}

func inputImageCode() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println("input")
	return text
}

func getVerifyCode(vcode string) (string, error) {
	sig := callCap_union_show(UIN, vcode)
	url := "http://captcha.qq.com/getimgbysig?aid=" + UIN + "&uin=" + UIN + "&sig=" + sig

	resp, err := httpClient.client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	} else {
		userFile := "/Users/XmacZone/Documents/java/golang/gopath/src/QQLogin/file/temp.png"
		fout, err := os.Create(userFile)
		if err != nil {
			fmt.Println(userFile, err)
			return "", errors.New("存储验证码失败！")
		}

		fileData, err := ioutil.ReadAll(resp.Body)
		fout.Write(fileData)
		fout.Close()
	}
	return sig, nil
}

func cap_union_verify(uin string, vcode string, sig string) (string, string, error) {
	url := "http://captcha.qq.com/cap_union_verify?aid=" + uin + "&uin=" + uin + "&captype=2&ans=" + vcode + "&sig=" + sig + "&0.06299702124670148"
	resp, err := httpClient.client.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
		return "", "", err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	respStr := string(data)

	respStr = strings.Replace(respStr, "cap_InnerCBVerify(", "", -1)
	respStr = strings.Replace(respStr, ");", "", -1)
	fmt.Println(respStr)

	var (
		ret string
		randstr string
		sigStr string
	)

	ret = respStr[strings.Index(respStr, "{rcode:") + len("{rcode:"):]
	ret = ret[0:strings.Index(ret, ",")]

	fmt.Println(ret)
	if ret == "0" {
		randstr = respStr[strings.Index(respStr, "randstr:\"") + len("randstr:\""):]
		randstr = randstr[:strings.Index(randstr, "\",")]

		sigStr = respStr[strings.Index(respStr, "sig:\"") + len("sig:\""):]
		sigStr = sigStr[:strings.Index(sigStr, "\",")]

		fmt.Println(randstr + ", sig = " + sigStr)
		return randstr, sigStr, nil
	}
	return "", "", errors.New("验证码错误")
}
