package VoltFuzer

import (

)

type VoltError string
func (v VoltError) Error() string {
	return string(v)
}

type Conduit interface {
	IsValid() bool
	Set(name string, val float64) 
	Get(name string) float64
	Parent() ParentConduit
}

type ParentConduit interface {
	IsValid() bool
	Set(name string, val float64) 
	Get(name string) float64
	Parent() ParentConduit
	ChildChanged(child Conduit, name string)
}

type Wire struct {
	U, I, R float64
	parent ParentConduit
}

type Parallel struct {
	U, I, R float64
	childA, childB Conduit
	parent ParentConduit
}

type Series struct {
	U, I, R float64
	childA, childB Conduit
	parent ParentConduit
}

func (w *Wire) Parent() ParentConduit {
	return w.parent
}
func (p *Parallel) Parent() ParentConduit {
	return p.parent
}
func (s *Series) Parent() ParentConduit {
	return s.parent
}

func (w *Wire) IsValid() bool {
	return true
}

func (w *Wire) Set(name string, val float64) {
	if val == 0.0 {
		panic("Set() : " + name + " must not be zero")
	}
	var ptr *float64
	switch name {
	case "U":
		ptr = &w.U
	case "I":
		ptr = &w.I
	case "R":
		ptr = &w.R
	default:
		panic("Set() : invalid name " + name)
	}
	if *ptr != 0.0 {
		panic("Set() : " + name + " is already set")
	} else {
		*ptr = val
	}
}

func (w *Wire) Get(name string) float64 {
	switch name {
	case "U":
		return w.U
	case "I":
		return w.I
	case "R":
		return w.R
	default:
		panic("Get() : invalid name " + name)
	}
}

func (p *Parallel) IsValid() bool {
	if p.childA != nil && 
	p.childA.Parent() == p && 
	p.childB != nil && 
	p.childB.Parent() == p && 
	p.childA.IsValid() && 
	p.childB.IsValid() {
		return true
	} else {
		return false
	}
}

func (s *Series) IsValid() bool {
	if s.childA != nil && 
	s.childA.Parent() == s && 
	s.childB != nil && 
	s.childB.Parent() == s && 
	s.childA.IsValid() && 
	s.childB.IsValid() {
		return true
	} else {
		return false
	}
}

func (p *Parallel) Set(name string, val float64) {
	// TODO: calculate new possible values
	if val == 0.0 {
		panic("Set() : " + name + " must not be zero")
	}
	var ptr *float64
	switch name {
	case "U":
		ptr = &p.U
	case "I":
		ptr = &p.I
	case "R":
		ptr = &p.R
	default:
		panic("Set() : invalid name " + name)
	}
	if *ptr != 0.0 {
		panic("Set() : " + name + " is already set")
	} else {
		*ptr = val
	}
}

func (p *Parallel) Get(name string) float64 {
	switch name {
	case "U":
		return p.U
	case "I":
		return p.I
	case "R":
		return p.R
	default:
		panic("Get() : invalid name " + name)
	}
}

func (p *Parallel) ChildChanged(child Conduit, name string) {
	var otherChild Conduit
	switch child {
	case p.childA:
		otherChild = p.childB
	case p.childB:
		otherChild = p.childA
	default:
		panic("ChildChanged() : this is not my child")
	}
	switch name {
	case "U":
		if p.U == 0.0 {
			p.Set("U", child.Get("U"))
		}
		if otherChild.Get("U") == 0.0 {
			otherChild.Set("U", child.Get("U"))
		}
	case "I":
		if p.I == 0.0 && otherChild.Get("I") != 0.0 {
			p.Set("I", child.Get("I") + otherChild.Get("I"))
		}
		if otherChild.Get("I") == 0.0 && p.I != 0.0 {
			otherChild.Set("I", p.I - child.Get("I"))
		}
	case "R":
		if p.R == 0.0 && otherChild.Get("R") != 0.0 {
			p.Set("R", 1 / ( 1 / child.Get("R") + 1 / otherChild.Get("R")))
		}
		if otherChild.Get("R") == 0.0 && p.R != 0.0 {
			otherChild.Set("R", 1 / ( 1 / p.R - 1 / child.Get("R")))
		}
	default:
		panic("ChildChanged() : invalid name " + name)
	}
}

func (s *Series) Set(name string, val float64) {
	// TODO: calculate new possible values
	if val == 0.0 {
		panic("Set() : " + name + " must not be zero")
	}
	var ptr *float64
	switch name {
	case "U":
		ptr = &s.U
	case "I":
		ptr = &s.I
	case "R":
		ptr = &s.R
	default:
		panic("Set() : invalid name " + name)
	}
	if *ptr != 0.0 {
		panic("Set() : " + name + " is already set")
	} else {
		*ptr = val
	}
}

func (s *Series) Get(name string) float64 {
	switch name {
	case "U":
		return s.U
	case "I":
		return s.I
	case "R":
		return s.R
	default:
		panic("Get() : invalid name " + name)
	}
}

func (s *Series) ChildChanged(child Conduit, name string) {
	var otherChild Conduit
	switch child {
	case s.childA:
		otherChild = s.childB
	case s.childB:
		otherChild = s.childA
	default:
		panic("ChildChanged() : this is not my child")
	}
	switch name {
	case "I":
		if s.I == 0.0 {
			s.Set("I", child.Get("I"))
		}
		if otherChild.Get("I") == 0.0 {
			otherChild.Set("I", child.Get("I"))
		}
	case "U":
		if s.U == 0.0 && otherChild.Get("U") != 0.0 {
			s.Set("U", child.Get("U") + otherChild.Get("U"))
		}
		if otherChild.Get("U") == 0.0 && s.U != 0.0 {
			otherChild.Set("U", s.U - child.Get("U"))
		}
	case "R":
		if s.R == 0.0 && otherChild.Get("R") != 0.0 {
			s.Set("R", child.Get("U") + otherChild.Get("U"))
		}
		if otherChild.Get("R") == 0.0 && s.R != 0.0 {
			otherChild.Set("R", s.R - child.Get("R"))
		}
	default:
		panic("ChildChanged() : invalid name " + name)
	}
}

