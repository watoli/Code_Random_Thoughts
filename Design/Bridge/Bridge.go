package Bridge

import "fmt"

type Draw interface {
	DrawCircle(radius, x, y int)
}
type RedCircle struct {
}

func (c *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("画红圆半径%d 横坐标%d 纵坐标%d\n", radius, x, y)
}

type YellowCircle struct {
}

func (c *YellowCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("画黄圆半径%d 横坐标%d 纵坐标%d\n", radius, x, y)
}

type Shape struct {
	draw Draw
}

func (s *Shape) Shape(d Draw) {
	s.draw = d
}

type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func (c *Circle) Constructor(x, y, radius int, draw Draw) {
	c.x = x
	c.y = y
	c.radius = radius
	c.shape.Shape(draw)
}
func (c *Circle) Draw() {
	c.shape.draw.DrawCircle(c.radius, c.x, c.y)
}
