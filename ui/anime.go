package ui

// EasingFunc is easing function for animation
type EasingFunc func(t float64) float64

// AnimationFunc is animate function for animation
type AnimationFunc func(el Element, t float64)

// Animation is animation
type Animation struct {
	IsLoop   bool
	Duration int
	Delay    int
	Ease     EasingFunc
	OnStart  AnimationFunc
	OnAnime  AnimationFunc
	OnEnd    AnimationFunc
}

// animation is internal animation setting
type animation struct {
	ID        int
	Element   Element
	IsStarted bool
	Start     int
	End       int
	Animation
}

// animationManager is animationManager
type animationManager struct {
	Map map[int]*animation
}

// SetAnimation set animation to animationManager
func (am *animationManager) SetAnimation(el Element, a Animation) {
	id := el.ID()
	if _, ok := am.Map[id]; ok {
		return
	}
	anime := &animation{
		ID:        id,
		Element:   el,
		IsStarted: false,
		Start:     gm.Now + a.Delay,
		End:       gm.Now + a.Delay + a.Duration,
		Animation: a,
	}
	am.Map[id] = anime
}

// StopAnimation stop animation
func (am *animationManager) StopAnimation(el Element) {
	id := el.ID()
	a, ok := am.Map[id]
	if !ok {
		return
	}
	if !a.IsStarted {
		delete(am.Map, id)
		return
	}
	elapsed := gm.Now - a.Start
	if a.IsLoop {
		elapsed %= a.Duration
	}
	t := float64(elapsed) / float64(a.Duration)
	if a.OnEnd != nil {
		a.OnEnd(a.Element, t)
	}
	delete(am.Map, id)
}

// Animate animation
func (am *animationManager) Animate() {
	for id, a := range am.Map {
		switch {
		case gm.Now < a.Start:
			// nothing to do in delay time.
		case a.End < gm.Now && !a.IsLoop:
			// close
			if a.OnEnd != nil {
				t := float64(gm.Now-a.Start) / float64(a.Duration)
				a.OnEnd(a.Element, t)
			}
			delete(am.Map, id)
		default:
			elapsed := gm.Now - a.Start
			if a.IsLoop {
				elapsed %= a.Duration
			}
			t := float64(elapsed) / float64(a.Duration)
			if !a.IsStarted {
				if a.OnStart != nil {
					a.OnStart(a.Element, a.Ease(t))
				}
				a.IsStarted = true
			}
			a.OnAnime(a.Element, a.Ease(t))
		}
	}
}

// Clear all animation
func (am *animationManager) Clear() {
	am.Map = map[int]*animation{}
}

// ClearAllAnimations clear all animations
func ClearAllAnimations() {
	gm.animationManager.Clear()
}
