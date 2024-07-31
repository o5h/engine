package glm

import (
	"fmt"
	"testing"
)

func TestLine2I_NaiveLineDrawing(t *testing.T) {

	line := Line2I{Start: Vec2I{0, 10}, End: Vec2I{0, 0}}
	line.NaiveLineDrawing(func(x, y int32) { fmt.Println(x, y) })

	line2 := Line2I{Start: Vec2I{10, 10}, End: Vec2I{10, 10}}
	line2.NaiveLineDrawing(func(x, y int32) { fmt.Println(x, y) })

	line3 := Line2I{Start: Vec2I{10, 2}, End: Vec2I{1, 5}}
	line3.NaiveLineDrawing(func(x, y int32) { fmt.Println(x, y) })

}
