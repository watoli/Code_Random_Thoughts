package AbstractFactory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	factory := NewSimpleLunchFactory()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
