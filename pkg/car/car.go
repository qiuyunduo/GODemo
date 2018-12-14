package car

import "fmt"

type Engine interface {
	Start()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func (c *Car) numberOfWheels() {
	fmt.Println("This car has ", c.wheelCount, " wheels")
}

func (c *Car) Start() {
	fmt.Println("the engine start working")
}

func (c *Car) goWorking() {
	c.Start()
}

func (m *Mercedes) SayHiToMerced1() {
	fmt.Println("Hi Merced1!!!")
}

func CarTest() {
	car := Car{nil, 4}
	car.goWorking()
	car.numberOfWheels()

	mercedes := Mercedes{car}
	mercedes.numberOfWheels()
	mercedes.SayHiToMerced1()
}
