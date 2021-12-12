package enum

type SellStatus int

const (
	Selling SellStatus = 0
	StopSell SellStatus = 1
)

func (s SellStatus) String() string {

	switch s {
	case Selling:
		return "销售中"
	case StopSell:
		return "停止销售"
	default:
		return "未知销售状态"

	}

}