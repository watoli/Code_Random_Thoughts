package Simple

import "fmt"

type Restaurant interface {
	GetFood()
}

type LanzhouNoodles struct {
}

func (n LanzhouNoodles) GetFood() {
	fmt.Println("煮一碗兰州拉面")
}

type KFC struct {
}

func (k KFC) GetFood() {
	fmt.Println("制作一份疯狂星期四套餐")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "L":
		return &LanzhouNoodles{}
	case "K":
		return &KFC{}
	}
	return nil
}
