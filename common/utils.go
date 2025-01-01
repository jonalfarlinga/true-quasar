package common

import "math/rand"

var r *rand.Rand = rand.New(rand.NewSource(99))

type Interactable interface {
	GetBounds() (int, int, int, int)
}

func Collide(x, y int, button Interactable) bool {
	left, top, right, bottom := button.GetBounds()
	return (x > left &&
		x < right &&
		y > top &&
		y < bottom)
}

func RandInt(min, max int) int {
	return r.Intn(max-min) + min
}
