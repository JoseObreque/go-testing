package hunt

import (
	"errors"
)

var (
	ErrTired      = errors.New("cannot hunt, i am really tired")
	ErrNotHungry  = errors.New("cannot hunt, i am not hungry")
	ErrPreyEscape = errors.New("could not catch it")
	ErrNoPrey     = errors.New("there is no prey to hunt")
)

// Shark is a struct that represents a shark.
type Shark struct {
	hungry bool
	tired  bool
	speed  int
}

// Prey is a struct that represents a prey that the shark can hunt.
type Prey struct {
	name  string
	speed int
}

// Hunt is a method of the Shark struct. It takes a pointer to a Prey struct as an argument and returns an error if
// the shark cannot hunt the prey. If the shark can hunt the prey, then the shark is no longer hungry and is now tired.
func (s *Shark) Hunt(p *Prey) error {
	// If there is no prey (nil pointer), return an error
	if p == nil {
		return ErrNoPrey
	}

	// If the shark is tired, return an error
	if s.tired {
		return ErrTired
	}

	// If the shark is not hungry, return an error
	if !s.hungry {
		return ErrNotHungry
	}

	// If the prey is faster than the shark, return an error
	if p.speed >= s.speed {
		s.tired = true
		return ErrPreyEscape
	}

	// The shark catch the prey and is now tired and not hungry
	s.hungry = false
	s.tired = true

	// Return nil to indicate no error
	return nil
}
