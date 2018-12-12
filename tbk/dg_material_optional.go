package tbk

import (
	"errors"
	"fmt"
)

type dgMaterialOptional struct {
	base
}

type DgMaterialOptionalInterface interface {
	Execute() (DgMaterialOptionalResponse, error)
	SetQ(str string)
	SetCat(cats []string)
	SetPlatform(platform int)
	SetPageNo(num int)
	SetAdzoneId(num int)
	SetPageSize(num int)
	SetSort(sort Sort)
	SetMaterialId(id int)
	SetHasCoupon(isHave bool)
	SetIp(ip string)
}

type DgMaterialOptionalItem struct {
	CouponStartTime string `json:"coupon_start_time"`
	CouponEndTime   string `json:"coupon_end_time"`
	InfoDxjh        string `json:"info_dxjh"`
	TkTotalSales    string `json:"tk_total_sales"`
	TkTotalCommi    string `json:"tk_total_commi"`
	CouponID        string `json:"coupon_id"`
	NumIid          int    `json:"num_iid"`
	Title           string `json:"title"`
	PictURL         string `json:"pict_url"`
	SmallImages     struct {
		String []string `json:"string"`
	} `json:"small_images"`
	ReservePrice         string `json:"reserve_price"`
	ZkFinalPrice         string `json:"zk_final_price"`
	UserType             int    `json:"user_type"`
	Provcity             string `json:"provcity"`
	ItemURL              string `json:"item_url"`
	IncludeMkt           string `json:"include_mkt"`
	IncludeDxjh          string `json:"include_dxjh"`
	CommissionRate       string `json:"commission_rate"`
	Volume               int    `json:"volume"`
	SellerID             int    `json:"seller_id"`
	CouponTotalCount     int    `json:"coupon_total_count"`
	CouponRemainCount    int    `json:"coupon_remain_count"`
	CouponInfo           string `json:"coupon_info"`
	CommissionType       string `json:"commission_type"`
	ShopTitle            string `json:"shop_title"`
	ShopDsr              int    `json:"shop_dsr"`
	CouponShareURL       string `json:"coupon_share_url"`
	URL                  string `json:"url"`
	LevelOneCategoryName string `json:"level_one_category_name"`
	LevelOneCategoryID   int    `json:"level_one_category_id"`
	CategoryName         string `json:"category_name"`
	CategoryID           int    `json:"category_id"`
	ShortTitle           string `json:"short_title"`
	WhiteImage           string `json:"white_image	"`
	Oetime               string `json:"oetime"`
	Ostime               string `json:"ostime"`
	JddNum               int    `json:"jdd_num"`
	JddPrice             int    `json:"jdd_price"`
	UvSumPreSale         int    `json:"uv_sum_pre_sale"`
}

type DgMaterialOptionalResponse struct {
	Results struct {
		Items []DgMaterialOptionalItem `json:"map_data"`
	} `json:"result_list"`
	TotalResults int `json:"total_results"`
}

// 淘宝客商品查询
func NewdgMaterialOptional() DgMaterialOptionalInterface {
	return &dgMaterialOptional{
		base: base{
			params:  map[string]string{},
			service: "taobao.tbk.dg.material.optional",
		},
	}
}

// 执行 Execute
func (t *dgMaterialOptional) Execute() (DgMaterialOptionalResponse, error) {
	var result struct {
		Tbk DgMaterialOptionalResponse `json:"tbk_dg_material_optional_response"`
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
func (t *dgMaterialOptional) SetAdzoneId(num int) {
	t.params["adzone_id"] = fmt.Sprintf("%d", num)
}

// 设置链接形式：1：PC，2：无线，默认：１ SetPlatform
func (t *dgMaterialOptional) SetPlatform(platform int) {
	t.base.SetPlatform(platform)
}

// 设置分类 SetCat
func (t *dgMaterialOptional) SetCat(cats []string) {
	t.base.SetCat(cats)
}

// 设置查询词 SetQ
func (t *dgMaterialOptional) SetQ(str string) {
	t.base.SetQ(str)
}

// 设置	第几页，默认：１ SetPageNo
func (t *dgMaterialOptional) SetPageNo(num int) {
	t.base.SetPageNo(num)
}

// 设置 页大小，默认20，1~100 SetPageSize
func (t *dgMaterialOptional) SetPageSize(num int) {
	t.base.SetPageSize(num)
}

// 排序_des（降序），
// 排序_asc（升序），
// 销量（total_sales），
// 淘客佣金比率（tk_rate），
// 累计推广量（tk_total_sales），
// 总支出佣金（tk_total_commi），
// 价格（price）
func (t *dgMaterialOptional) SetSort(sort Sort) {
	t.base.SetSort(sort)
}

// 官方的物料Id
// (详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)，
// 不传时默认为2836
func (t *dgMaterialOptional) SetMaterialId(id int) {
	t.params["material_id"] = fmt.Sprintf("%d", id)
}

// 是否有优惠券，
// 设置为true表示该商品有优惠券，
// 设置为false或不设置表示不判断这个属性
func (t *dgMaterialOptional) SetHasCoupon(isHave bool) {
	t.params["has_coupon"] = fmt.Sprintf("%t", isHave)
}

// ip参数影响邮费获取，
// 如果不传或者传入不准确，
// 邮费无法精准提供
func (t *dgMaterialOptional) SetIp(ip string) {
	t.params["ip"] = fmt.Sprintf("%s", ip)
}
