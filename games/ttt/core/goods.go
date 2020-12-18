package core

//TODO apparently bad idea, rely on specific types
const (
	Passengers GoodsType = 0
	Mail       GoodsType = 1
)

const (
	PassengersDefaultCost = 1
	MailDefaultCost       = 2
)

type GoodsType int

type Goods struct {
	Name           string
	Type           GoodsType
	Count          int
	DeliveredCount int
	Cost           int
}

func NewPassenger(count int, cost int) *Goods {
	return &Goods{"Passengers", Passengers, count, 0, cost}
}

func NewMail(count int, cost int) *Goods {
	return &Goods{"Mail", Mail, count, 0, cost}
}
