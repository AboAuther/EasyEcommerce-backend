package api

type (
	PayType      int
	ResponseType int
	SellStatus   int
	PayStatus    int
)

const (
	OperateOk   ResponseType = 200
	OperateFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case OperateOk:
		return "Ok"
	case OperateFail:
		return "Fail"
	default:
		return "UNKNOWN"
	}
}

const (
	Selling  SellStatus = 0
	StopSell SellStatus = 1
)

func (p SellStatus) String() string {
	switch p {
	case Selling:
		return "销售中"
	case StopSell:
		return "停止销售"
	default:
		return "UNKNOWN"
	}
}

const (
	Bank   PayType = 0
	WeChat PayType = 1
	AliPay PayType = 2
)

func (p PayType) String() string {
	switch p {
	case Bank:
		return "银行卡"
	case WeChat:
		return "微信"
	case AliPay:
		return "支付宝"
	default:
		return "UNKNOWN"
	}
}

const (
	UnPay PayStatus = 0
	Payed PayStatus = 1
)

func (p PayStatus) String() string {
	switch p {
	case UnPay:
		return "未支付"
	case Payed:
		return "已支付"
	default:
		return "UNKNOWN"
	}
}
