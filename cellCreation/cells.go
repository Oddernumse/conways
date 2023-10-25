package cellCreation

import (
	"math/rand"
	gamebehaviour "render/gameBehaviour"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	rows    = 100
	columns = 100

	threshold = 0.15
)

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

func chooseColor() []float32 {
	color := []float32{rand.Float32(), rand.Float32(), rand.Float32()}

	var (
		colors = []float32{
			color[0], color[1], color[2], 1.0,
			color[0], color[1], color[2], 1.0,
			color[0], color[1], color[2], 1.0,

			color[0], color[1], color[2], 1.0,
			color[0], color[1], color[2], 1.0,
			color[0], color[1], color[2], 1.0,
		}
	)

	return colors
}

// This is our drawable object
func makeVao(points []float32, colors []float32) uint32 {
	// Creates buffer object
	var colorBuffer uint32

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	gl.GenBuffers(1, &colorBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(colors), gl.Ptr(colors), gl.STATIC_DRAW)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(1)

	// creates actual object to draw
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

// Fills the board with our cell struct
func MakeCells() [][]*gamebehaviour.Cell {
	rand.Seed(time.Now().UnixNano())

	cells := make([][]*gamebehaviour.Cell, rows, rows)
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			c := newCell(x, y)

			c.Alive = rand.Float64() < threshold
			c.AliveNext = c.Alive

			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

func newCell(x, y int) *gamebehaviour.Cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)

	// This decides whether we're at an x or y of the shape
	// aka i % 3 = 0 = x and = 1 = y ty
	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(rows)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	colors := chooseColor()

	return &gamebehaviour.Cell{
		Drawable: makeVao(points, colors),

		X: x,
		Y: y,
	}
}
