package life

import (
	"errors"
	"math/rand"
	"time"
)

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if width <= 0 || height <= 0 {
		return &World{}, errors.New("Ширина и высота должны быть положительными")
	}

	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)
	alive := w.Cells[y][x]
	if n < 4 && n > 1 && alive {
		return true
	}
	if n == 3 && !alive {
		return true
	}

	return false
}

func (w *World) neighbors(x, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			x_pos, y_pos := x+j, y+i
			if x_pos < 0 {
				x_pos += w.Width
			}
			if x_pos >= w.Width {
				x_pos -= w.Width
			}
			if y_pos < 0 {
				y_pos += w.Height
			}
			if y_pos >= w.Height {
				y_pos -= w.Height
			}
			if !(i == 0 && j == 0) && w.Cells[y_pos][x_pos] {
				count++
			}
		}
	}

	return count
}

func NextState(oldWorld, newWorld World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			// для каждой клетки получим новое состояние
			newWorld.Cells[i][j] = oldWorld.next(j, i)
		}
	}
}

func (w *World) RandInit(percentage int) {
	numAlive := percentage * w.Height * w.Width / 100
	w.fillAlive(numAlive)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
