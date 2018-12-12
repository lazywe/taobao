package tbk

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type itemGet struct {
	base
}

type ItemGetInterface interface {
	Execute() (ItemGetResponse, error)
	SetFields(fields []string)
	SetQ(str string)
	SetCat(cats []string)
	SetItemLoc(str string)
	SetSort(sort Sort)
	SetIsTmall(b bool)
	SetIsOverseas(b bool)
	SetStartPrice(num int)
	SetEndPrice(num int)
	SetStartTkRate(num int)
	SetEndTkRate(num int)
	SetPlatform(platform int)
	SetPageNo(num int)
	SetPageSize(num int)
}

type Item struct {
	NumIid      int    `json:"num_iid"`
	Title       string `json:"title"`
	PictURL     string `json:"pict_url"`
	SmallImages struct {
		String []string `json:"string"`
	} `json:"small_images"`
	ReservePrice string `json:"reserve_price"`
	ZkFinalPrice string `json:"zk_final_price"`
	UserType     int    `json:"user_type"`
	Provcity     string `json:"provcity"`
	ItemURL      string `json:"item_url"`
	Nick         string `json:"nick"`
	SellerID     int    `json:"seller_id"`
	Volume       int    `json:"volume"`
}

type ItemGetResponse struct {
	Results struct {
		Items []Item `json:"n_tbk_item"`
	} `json:"results"`
	TotalResults int    `json:"total_results"`
	RequestId    string `json:"request_id"`
}

// 淘宝客商品查询
func NewItemGet() ItemGetInterface {
	return &itemGet{
		base: base{
			params: map[string]string{},
			fields: []string{
				"num_iid",
				"title",
				"pict_url",
				"small_images",
				"reserve_price",
				"zk_final_price",
				"user_type",
				"provcity",
				"item_url",
				"seller_id",
				"volume",
				"nick",
			},
			service: "taobao.tbk.item.get",
		},
	}
}

// 执行 Execute
func (t *itemGet) Execute() (ItemGetResponse, error) {
	var result struct {
		Tbk ItemGetResponse `json:"tbk_item_get_response"`
	}
	if t.params["q"] == "" && t.params["cats"] == "" {
		return result.Tbk, errors.New("q参数与cats参数不可同时为空")
	}
	if t.params["fields"] == "" {
		return result.Tbk, errors.New("fields参数不可同时为空")
	}
	if err := t.base.Execute(&result); err != nil {
		return result.Tbk, err
	}
	return result.Tbk, nil
}

// 设置字段 SetFields
func (t *itemGet) SetFields(fields []string) {
	t.base.SetFields(fields)
}

// 设置查询词 SetQ
func (t *itemGet) SetQ(str string) {
	t.params["q"] = str
}

// 设置分类 SetCat
func (t *itemGet) SetCat(cats []string) {
	if len(cats) <= 0 {
		log.Fatalln(errors.New("fileds 不能为空"))
	}
	t.params["cats"] = strings.Join(cats, ",")
}

// 设置所在地 SetItemLoc
func (t *itemGet) SetItemLoc(str string) {
	t.params["itemloc"] = str
}

// 设置排序 SetSort
func (t *itemGet) SetSort(sort Sort) {
	t.base.SetSort(sort)
}

// 设置是否商城商品 SetIsTmall
func (t *itemGet) SetIsTmall(b bool) {
	t.params["is_tmall"] = fmt.Sprintf("%t", b)
}

// 设置是否海外商品 SetIsOverseas
func (t *itemGet) SetIsOverseas(b bool) {
	t.params["is_overseas"] = fmt.Sprintf("%t", b)

}

// 设置折扣价范围下限，单位：元 SetStartPrice
func (t *itemGet) SetStartPrice(num int) {
	t.params["start_price"] = fmt.Sprintf("%d", num)
}

// 设置折扣价范围上限，单位：元 SetEndPrice
func (t *itemGet) SetEndPrice(num int) {
	t.params["end_price"] = fmt.Sprintf("%d", num)
}

// 设置淘客佣金比率上限，如：1234表示12.34% SetStartTkRate
func (t *itemGet) SetStartTkRate(num int) {
	t.params["start_tk_rate"] = fmt.Sprintf("%d", num)
}

// 设置淘客佣金比率下限，如：1234表示12.34% SetEndTkRate
func (t *itemGet) SetEndTkRate(num int) {
	t.params["end_tk_rate"] = fmt.Sprintf("%d", num)
}

// 设置链接形式：1：PC，2：无线，默认：１ SetPlatform
func (t *itemGet) SetPlatform(platform int) {
	t.base.SetPlatform(platform)

}

// 设置	第几页，默认：１ SetPageNo
func (t *itemGet) SetPageNo(num int) {
	t.base.SetPageNo(num)
}

// 设置 页大小，默认20，1~100 SetPageSize
func (t *itemGet) SetPageSize(num int) {
	t.base.SetPageSize(num)
}
