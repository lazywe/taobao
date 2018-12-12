package tbk

import (
	"errors"
	"fmt"
)

type tpwdCreate struct {
	base
}

type TpwdCreateInterface interface {
	Execute() (TpwdCreateResponse, error)
	SetUserId(userid string)
	SetText(text string)
	SetUrl(url string)
	SetLog(logo string)
	SetExt(ext string)
}

type TpwdCreateResponse struct {
	Data struct {
		Model string `json:"model"`
	} `json:"data"`
}

// 淘宝客淘口令
func NewTpwdCreate() TpwdCreateInterface {
	return &tpwdCreate{
		base: base{
			params:  map[string]string{},
			service: "taobao.tbk.tpwd.create",
		},
	}
}

// 执行 Execute
func (t *tpwdCreate) Execute() (TpwdCreateResponse, error) {
	var result TpwdCreateResponse
	if t.params["text"] == "" {
		return result, errors.New("text参数不可为空")
	}
	if t.params["url"] == "" {
		return result, errors.New("url参数不可为空")
	}
	if err := t.base.Execute(&result); err != nil {
		return result, err
	}
	return result, nil
}

// 生成口令的淘宝用户ID SetUserId
func (t *tpwdCreate) SetUserId(userid string) {
	t.params["user_id"] = fmt.Sprintf("%s", userid)
}

// 口令弹框内容 SetText
func (t *tpwdCreate) SetText(text string) {
	t.params["text"] = fmt.Sprintf("%s", text)
}

// 口令跳转目标页 SetUrl
func (t *tpwdCreate) SetUrl(url string) {
	t.params["url"] = fmt.Sprintf("%s", url)
}

// 口令弹框logoURL SetLog
func (t *tpwdCreate) SetLog(logo string) {
	t.params["logo"] = fmt.Sprintf("%s", logo)
}

// 扩展字段JSON格式 SetExt
func (t *tpwdCreate) SetExt(ext string) {
	t.params["ext"] = fmt.Sprintf("%s", ext)
}
