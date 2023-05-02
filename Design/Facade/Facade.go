package Facade

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}
func (m *CarModel) SetModel() {
	fmt.Println("设置汽车中")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}
func (e CarEngine) SetEngine() {
	fmt.Println("设置引擎中")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}
func (e CarBody) SetBody() {
	fmt.Println("设置外观中")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (f *CarFacade) CreateCompleteCar() {
	f.model.SetModel()
	f.engine.SetEngine()
	f.body.SetBody()
}
