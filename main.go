package main

import (
	"render/cellCreation"
	gamebehaviour "render/gameBehaviour"
	"render/opengl"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	fps = 100
)

func main() {
	// All our shit with this window gotta happen on the same thread
	runtime.LockOSThread()

	window := opengl.InitGlfw()
	defer glfw.Terminate()

	program := opengl.InitOpenGL()

	// vao := makeVao(square)
	cells := cellCreation.MakeCells()

	for !window.ShouldClose() {
		t := time.Now()

		for x := range cells {
			for _, c := range cells[x] {
				c.CheckState(cells)
			}
		}

		draw(cells, window, program)

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

// Clears whats currently displayed and displays the next thing/switches the buffer
func draw(cells [][]*gamebehaviour.Cell, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	// Draws the cells duh
	for x := range cells {
		for _, c := range cells[x] {
			c.Draw()
		}
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
