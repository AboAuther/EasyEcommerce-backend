package api

import "math"

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

const (
	AllCategory = iota
	VegetableSoybean
	MeatPoultryEggMilk
	SeafoodAquaculture
	Fruit
	FrozenFood
)

const (
	AllPrice = iota
	ZeroToFifty
	FiftyToHundred
	HundredToDouble
	BiggerThanTwoHundred
)

var (
	PriceMap = map[int][2]int{
		AllPrice:             {0, math.MaxInt},
		ZeroToFifty:          {0, 50},
		FiftyToHundred:       {51, 100},
		HundredToDouble:      {101, 200},
		BiggerThanTwoHundred: {200, math.MaxInt},
	}
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
