package tbk

import (
	"fmt"
	"strings"
	"taobao/lib"
)

type base struct {
	params map[string]string
	// 可用字段
	fields []string
	// 服务
	service string
}

// 执行淘宝客商品查询 Execute
func (t *base) Execute(result interface{}) error {
	err := lib.NewHttp().Request(lib.Post, t.service, t.params, result)
	if err != nil {
		return err
	}
	return nil
}

// 设置字段 SetFields
func (t *base) SetFields(fields []string) {
	if len(fields) <= 0 {
		return
	}
	validFields := lib.GetValidFields(t.fields, fields)
	if len(validFields) <= 0 {
		return
	}
	t.params["fields"] = strings.Join(validFields, ",")
}

// 设置链接形式：1：PC，2：无线，默认：１ SetPlatform
func (t *base) SetPlatform(platform int) {
	if platform != 1 || platform != 2 {
		platform = 1
	}
	t.params["platform"] = fmt.Sprintf("%d", platform)
}

// 设置分类 SetCat
func (t *base) SetCat(cats []string) {
	t.params["cats"] = strings.Join(cats, ",")
}

// 设置查询词 SetQ
func (t *base) SetQ(str string) {
	t.params["q"] = str
}

// 设置 页大小，默认20，1~100 SetPageSize
func (t *base) SetPageSize(num int) {
	if num <= 0 {
		num = 1
	}
	t.params["page_size"] = fmt.Sprintf("%d", num)
}

// 设置	第几页，默认：１ SetPageNo
func (t *base) SetPageNo(num int) {
	if num <= 0 {
		num = 1
	}
	t.params["page_no"] = fmt.Sprintf("%d", num)
}

// 设置排序 SetSort
func (t *base) SetSort(sort Sort) {
	t.params["sort"] = sort.String()
}
