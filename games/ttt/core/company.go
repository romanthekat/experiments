package core

type Company struct {
	Name  string
	Money int64
}

func NewCompany(name string, money int64) *Company {
	return &Company{Name: name, Money: money}
}
