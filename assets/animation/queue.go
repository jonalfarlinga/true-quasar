package animation

import "github.com/hajimehoshi/ebiten/v2"

type AnimationQueue struct {
	queue []string
}

func NewAnimationQueue() *AnimationQueue {
	return &AnimationQueue{
		queue: make([]string, 0),
	}
}

func (a *AnimationQueue) Length() int {
	return len(a.queue)
}

func (a *AnimationQueue) AddAnimation(animation string) {
	a.queue = append(a.queue, animation)
}

func (a *AnimationQueue) Animate(screen *ebiten.Image) {
	if len(a.queue) == 0 {
		return
	}
	
}
