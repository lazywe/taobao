package tbk

import (
	"errors"
	"fmt"
)

type dgItemCouponGet struct {
	base
}

type DgItemCouponGetInterface interface {
	Execute() (DgItemCouponGetResponse, error)
	SetQ(str string)
	SetCat(cats []string)
	SetPlatform(platform int)
	SetPageNo(num int)
	SetAdzoneId(num int)
	SetPageSize(num int)
}

type DgItemCouponItem struct {
	ShopTitle         string `json:"shop_title"`
	CouponTotalCount  int    `json:"coupon_total_count"`
	CommissionRate    string `json:"commission_rate"`
	CouponInfo        string `json:"coupon_info"`
	Category          int    `json:"category"`
	CouponRemainCount int    `json:"coupon_remain_count"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponClickURL    string `json:"coupon_click_url"`
	ItemCescription   string `json:"item_description"`
	NumIid            int    `json:"num_iid"`
	Title             string `json:"title"`
	PictURL           string `json:"pict_url"`
	SmallImages       struct {
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

type DgItemCouponGetResponse struct {
	Results struct {
		Items []DgItemCouponItem `json:"tbk_coupon"`
	} `json:"results"`
	TotalResults int    `json:"total_results"`
	RequestId    string `json:"request_id"`
}

// 淘宝客商品查询
func NewDgItemCouponGet() DgItemCouponGetInterface {
	return &dgItemCouponGet{
		base: base{
			params:  map[string]string{},
			service: "taobao.tbk.dg.item.coupon.get",
		},
	}
}

// 执行 Execute
func (t *dgItemCouponGet) Execute() (DgItemCouponGetResponse, error) {
	var result struct {
		Tbk DgItemCouponGetResponse `json:"tbk_dg_item_coupon_get_response"`
	}
	if t.params["adzone_id"] == "" {
		return result.Tbk, errors.New("adzone_id参数不可为空")
	}
	if t.params["q"] == "" && t.params["cats"] == "" {
		return result.Tbk, errors.New("q参数与cats参数不可为空")
	}
	if err := t.base.Execute(&result); err != nil {
		return result.Tbk, err
	}
	return result.Tbk, nil
}

// 设置mm_xxx_xxx_xxx的第三位 SetAdzoneId
func (t *dgItemCouponGet) SetAdzoneId(num int) {
	t.params["adzone_id"] = fmt.Sprintf("%d", num)
}

// 设置链接形式：1：PC，2：无线，默认：１ SetPlatform
func (t *dgItemCouponGet) SetPlatform(platform int) {
	t.base.SetPlatform(platform)
}

// 设置分类 SetCat
func (t *dgItemCouponGet) SetCat(cats []string) {
	t.base.SetCat(cats)
}

// 设置查询词 SetQ
func (t *dgItemCouponGet) SetQ(str string) {
	t.base.SetQ(str)
}

// 设置	第几页，默认：１ SetPageNo
func (t *dgItemCouponGet) SetPageNo(num int) {
	t.base.SetPageNo(num)
}

// 设置 页大小，默认20，1~100 SetPageSize
func (t *dgItemCouponGet) SetPageSize(num int) {
	t.base.SetPageSize(num)
}
