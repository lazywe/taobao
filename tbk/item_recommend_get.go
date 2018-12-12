package tbk

import (
	"errors"
	"fmt"
)

type itemRecommendGet struct {
	base
}

type ItemRecommendGetInterface interface {
	Execute() (ItemRecommendGetResponse, error)
	SetFields(fields []string)
	SetNumIid(id int)
	SetCount(count int)
	SetPlatform(platform int)
}

type RecommendItem struct {
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
}

type ItemRecommendGetResponse struct {
	Results struct {
		RecommendItems []RecommendItem `json:"n_tbk_item"`
	} `json:"results"`
	RequestId string `json:"request_id"`
}

// 淘宝客商品关联推荐查询
func NewItemRecommendGet() ItemRecommendGetInterface {
	return &itemRecommendGet{
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
			},
			service: "taobao.tbk.item.recommend.get",
		},
	}
}

// 执行 Execute
func (t *itemRecommendGet) Execute() (ItemRecommendGetResponse, error) {
	var result struct {
		Tbk ItemRecommendGetResponse `json:"tbk_item_recommend_get_response"`
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
func (t *itemRecommendGet) SetFields(fields []string) {
	t.base.SetFields(fields)
}

// 	商品Id
func (t *itemRecommendGet) SetNumIid(id int) {
	t.params["num_iid"] = fmt.Sprintf("%d", id)
}

// 返回数量，默认20，最大值40
func (t *itemRecommendGet) SetCount(count int) {
	if count >= 40 {
		count = 40
	}
	t.params["count"] = fmt.Sprintf("%d", count)
}

// 链接形式：1：PC，2：无线，默认：１
func (t *itemRecommendGet) SetPlatform(platform int) {
	t.base.SetPlatform(platform)
}
