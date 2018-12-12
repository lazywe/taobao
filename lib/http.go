package lib

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type httpRequest struct {
	AppKey    string
	AppSecret string
}

type HttpRequestInterface interface {
	Request(method Method, service string, params map[string]string, result interface{}) error
}

var (
	httpUrl = "http://gw.api.taobao.com/router/rest"
)

// Http请求 NewHttp
func NewHttp() HttpRequestInterface {
	c := GetConfig()
	return &httpRequest{
		AppKey:    c.AppKey,
		AppSecret: c.SecretKey,
	}
}

// 网络请求 Request
func (t *httpRequest) Request(method Method, service string, params map[string]string, result interface{}) error {
	if err := t.checkConfig(); err != nil {
		return err
	}
	if service == "'" {
		return errors.New("service不能为空")
	}
	if c.Debug == true {
		fmt.Println("------------------")
		fmt.Println("请求参数:")
		fmt.Println(params)
	}
	params["method"] = service
	requestBody := strings.NewReader(t.makeRequestParams(params))
	req, err := http.NewRequest(method.String(), httpUrl, requestBody)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpclient := &http.Client{}
	res, err := httpclient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("请求错误码:[%d]，请求body:[%x]", res.StatusCode, res.Body)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if c.Debug == true {
		fmt.Println("返回数据:")
		fmt.Println(string(b))
		fmt.Println("------------------")
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return err
	}
	return nil
}

func (t *httpRequest) checkConfig() error {
	if t.AppKey == "" {
		return errors.New("AppKey未配置")
	}
	if t.AppSecret == "" {
		return errors.New("AppSecret未配置")
	}
	return nil
}

// 默认参数 defaultParams
func (t *httpRequest) defaultParams() url.Values {
	args := url.Values{}
	//默认参数
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	args.Add("timestamp", timestamp)
	args.Add("format", "json")
	args.Add("app_key", t.AppKey)
	args.Add("v", "2.0")
	args.Add("sign_method", "md5")

	return args
}

// 生成请求参数 makeRequestParams
func (t *httpRequest) makeRequestParams(params map[string]string) string {
	args := t.defaultParams()
	for key, val := range params {
		args.Set(key, val)
	}
	args.Add("sign", t.sign(args))
	return args.Encode()
}

// 签名 sign
func (t *httpRequest) sign(args url.Values) string {
	keys := []string{}
	for k := range args {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pstr := ""
	for _, k := range keys {
		pstr += k + args.Get(k)
	}
	sign := md5.Sum([]byte(t.AppSecret + pstr + t.AppSecret))
	return strings.ToUpper(hex.EncodeToString(sign[:]))
}
