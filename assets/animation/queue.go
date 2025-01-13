package animation

type AnimationQueue struct {
	queue []string
}

func (a *AnimationQueue) Length() int {
	return len(a.queue)
}
