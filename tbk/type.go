package tbk

/// 定义排序
type Sort int

// iota 初始化后会自动递增
const (
	Desc Sort = iota // value --> 0
	Asc              // value --> 1
	TotalSales
	TkRate
	TkTotalSales
	TkTotalCommi
	Price
)

func (this Sort) String() string {
	switch this {
	case Desc:
		return "_des"
	case Asc:
		return "_asc"
	case TotalSales:
		return "total_sales"
	case TkRate:
		return "tk_rate"
	case TkTotalSales:
		return "tk_total_sales"
	case TkTotalCommi:
		return "tk_total_commi"
	case Price:
		return "price"
	default:
		return ""
	}
}
