package goods

import "fmt"

type Good int

const (
	Undefined Good = iota
	Wood
	Timber
	Clay
	Brick
	Peat
	Food
	Grain
	Hide
	Flax
	Wool
	Linen
	Woolen
	Leather
	SummerWear
	WinterWear
	LeatherWear
)

var maxGood int = 15

var goodsTrack = map[Good]struct{}{
	Grain: struct{}{},
	Hide:  struct{}{},
	Flax:  struct{}{},
	Wool:  struct{}{}}

type Goods struct {
	m map[Good]int
}

func NewGoods() *Goods {
	return &Goods{map[Good]int{
		Wood:  4,
		Clay:  4,
		Peat:  3,
		Food:  5,
		Wool:  4,
		Flax:  3,
		Hide:  2,
		Grain: 1,
	}}
}

func (g Goods) Get(good Good) int {
	return g.m[good]
}

func (g Goods) Add(good Good) {
	g.Increment(good, 1)
}

func (g Goods) Increase(good Good, i int) {
	g.Increment(good, i)
}

//TODO put cap on goods
//Deprecated. use Increase()
func (g Goods) Increment(good Good, i int) {
	newVal := g.m[good] + i
	if good == Food {
		if newVal > 2*maxGood {
			newVal = 2 * maxGood
		}
	} else if _, ok := goodsTrack[good]; ok {
		if newVal > maxGood {
			newVal = maxGood
		}
	}
	g.m[good] = newVal
}

func (g Goods) Use(good Good, i int) error {
	if good == Wood {
		return g.UseWithSubstitute(Wood, Timber, i)
	} else if good == Clay {
		return g.UseWithSubstitute(Clay, Brick, i)
	}
	if g.m[good] < i {
		return InsufficientGoodsError{good}
	}
	g.m[good] -= i
	return nil
}

func (g Goods) UseWithSubstitute(good Good, sub Good, i int) error {
	if g.m[good]+g.m[sub] < i {
		return InsufficientGoodsError{good}
	}
	g.m[good] -= i
	if g.m[good] < 0 {
		g.m[sub] -= g.m[good]
		g.m[good] = 0
	}
	return nil
}

func (g Goods) UseWithSubstitutes(good Good, sub Good, sub2 Good, i int) error {
	if g.m[good]+g.m[sub]+g.m[sub2] < i {
		return InsufficientGoodsError{good}
	}
	g.m[good] -= i
	if g.m[good] < 0 {
		g.m[sub] -= g.m[good]
		g.m[good] = 0
	}
	if g.m[sub] < 0 {
		g.m[sub2] -= g.m[sub]
		g.m[sub] = 0
	}
	return nil
}

type InsufficientGoodsError struct {
	Good
}

func (i InsufficientGoodsError) Error() string {
	return fmt.Sprintf("Did not have enough of %s good", i.Good)
}
