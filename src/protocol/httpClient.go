package protocol
import (
	"net/http"
)

type HttpClient struct {
	client *http.Client
//	jar *cookiejar.Jar
}

func NewHttpClient() *HttpClient {
	httpClient := new(HttpClient)
//	httpClient.jar = new(cookiejar.Jar)
	httpClient.client = &http.Client{}
	return httpClient
}