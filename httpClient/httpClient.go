package httpClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpClient struct {
	URL     string
	Method  string
	Param   interface{}
	Headers map[string]string
	request *http.Request
}

//
func New(url string) *HttpClient {
	httpClient := &HttpClient{
		URL:    url,
		Method: "GET",
		Param:  map[string]interface{}{},
		Headers: map[string]string{
			"Content-Type": "application-json",
		},
	}
	return httpClient
}

func (this *HttpClient) SetHeaders(headers map[string]string) *HttpClient {
	for key, value := range headers {
		this.Headers[key] = value
	}
	return this
}

func (this *HttpClient) Get(result interface{}) error {
	this.SetMethod("GET")
	request, err := http.NewRequest(this.Method, this.URL, nil)
	if err != nil {
		return err
	}
	this.request = request
	if len(this.Headers) > 0 {
		for key, value := range this.Headers {
			this.request.Header.Set(key, value)
		}
	}
	client := http.Client{}
	resp, err := client.Do(this.request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(body)
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

func (this *HttpClient) Post(data interface{}, result interface{}) error {
	this.SetMethod("POST")
	this.Param = data
	var reqBody *strings.Reader
	if this.Param != nil {
		bParam, err := json.Marshal(this.Param)
		if err != nil {
			return err
		}
		reqBody = strings.NewReader(string(bParam))
	}

	request, err := http.NewRequest(this.Method, this.URL, reqBody)
	if err != nil {
		return err
	}
	this.request = request
	if len(this.Headers) > 0 {
		for key, value := range this.Headers {
			this.request.Header.Set(key, value)
		}
	}
	client := http.Client{}
	resp, err := client.Do(this.request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}

func (this *HttpClient) SetMethod(method string) *HttpClient {
	this.Method = method
	return this
}

func (this *HttpClient) newRequest() {
}
