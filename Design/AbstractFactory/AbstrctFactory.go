package AbstractFactory

import "fmt"

type Lunch interface {
	Cook()
}
type Rice struct {
}

func (r *Rice) Cook() {
	fmt.Println("正在烹饪米饭")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("正在烹饪土豆")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

func (f *SimpleLunchFactory) CreateFood() Lunch {
	return &Rice{}
}

func (f *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
