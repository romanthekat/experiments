package core

import (
	"fmt"
)

const (
	cityTicksToGrowth = ticksPerSecond
)

type City struct {
	Name       string
	Population int
	Goods      map[GoodsType]*Goods
	Ticks      int
}

func (c City) String() string {
	return fmt.Sprintf("%s %d %s", c.Name, c.Population, c.Goods, c.Goods)
}

func (c *City) Tick() {
	c.Ticks += 1
	if c.Ticks < cityTicksToGrowth {
		return
	}
	c.Ticks = 0

	passengers := c.getGoods(Passengers)
	passengers.Count = Min(passengers.Count+c.Population/100, c.Population/7) + Min(passengers.DeliveredCount/10, 42)

	mail := c.getGoods(Mail)
	mail.Count = Min(mail.Count+passengers.Count/5+c.Population/300, c.Population/5) + Min(mail.DeliveredCount/10, 42)

	populationDelta := c.Population/1000 + passengers.Count/100 + mail.Count/100
	if populationDelta == 0 {
		populationDelta += 1
	}

	c.Population += populationDelta
}

func (c *City) getGoods(goodsType GoodsType) *Goods {
	goods := c.Goods[goodsType]
	if goods == nil {
		if goodsType == Passengers {
			goods = NewPassenger(0, PassengersDefaultCost)
		} else if goodsType == Mail {
			goods = NewMail(0, MailDefaultCost)
		}

		c.Goods[goodsType] = goods
	}

	return goods
}

func NewCity(name string, population int) *City {
	return &City{Name: name, Population: population, Goods: map[GoodsType]*Goods{}}
}
