package lib

/// 定义http请求类型
type Method int

// iota 初始化后会自动递增
const (
	Get  Method = iota // value --> 0
	Post               // value --> 1
)

func (this Method) String() string {
	switch this {
	case Get:
		return "GET"
	case Post:
		return "POST"
	default:
		return ""
	}
}
