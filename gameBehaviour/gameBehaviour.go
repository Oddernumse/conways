package gamebehaviour

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Cell struct {
	Drawable uint32

	Alive     bool
	AliveNext bool

	X int
	Y int
}

var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func (c *Cell) CheckState(cells [][]*Cell) {
	c.Alive = c.AliveNext
	c.AliveNext = c.Alive

	liveCount := c.liveNeighbours(cells)
	if c.Alive {
		if liveCount < 2 {
			c.AliveNext = false
		}

		if liveCount == 2 || liveCount == 3 {
			c.Alive = true
		}

		if liveCount > 3 {
			c.AliveNext = false
		}
	} else {
		if liveCount == 3 {
			c.AliveNext = true
		}
	}
}

func (c *Cell) liveNeighbours(cells [][]*Cell) int {
	var liveCount int
	add := func(x, y int) {
		if x == len(cells) {
			x = 0
		} else if x == -1 {
			x = len(cells) - 1
		}
		if y == len(cells[x]) {
			y = 0
		} else if y == -1 {
			y = len(cells[x]) - 1
		}

		if cells[x][y].Alive {
			liveCount++
		}
	}

	add(c.X-1, c.Y)
	add(c.X+1, c.Y)
	add(c.X, c.Y+1)
	add(c.X, c.Y-1)
	add(c.X-1, c.Y+1)
	add(c.X+1, c.Y+1)
	add(c.X-1, c.Y-1)
	add(c.X+1, c.Y-1)

	return liveCount
}

func (c *Cell) Draw() {
	if !c.Alive {
		return
	}

	gl.BindVertexArray(c.Drawable)
	// telling it to draw however many vertices we defined
	// Gets the specific number by removing X, Y, Z hence the / 3
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}
