package common

type Interactable interface {
	GetBounds() (int, int, int, int)
}

func Collide(x, y int, button Interactable) bool {
	left, top, right, bottom := button.GetBounds()
	return x > left && x < right && y > top && y < bottom
}
