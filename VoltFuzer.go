package VoltFuzer

import (

)

type Conduit interface {
	isValid() bool
	setU(float64) error
	setR(float64) error
	setI(float64) error
	Parent() Conduit
}

type Wire struct {
	u, i, r float64
	parent Conduit
}

type Parallel struct {
	u, i, r float64
	parent, childA, childB Conduit
}

type Series struct {
	u, i, r float64
	parent, childA, childB Conduit
}

func (w *Wire) Parent() Conduit {
	return w.parent
}
func (p *Parallel) Parent() Conduit {
	return p.parent
}
func (s *Series) Parent() Conduit {
	return s.parent
}

func (w *Wire) isValid() bool {
	return true
}

func (p *Parallel) isValid() bool {
	if p.childA != nil && 
	p.childA.Parent() == p && 
	p.childB != nil && 
	p.childB.Parent() == p && 
	p.childA.isValid() && 
	p.childB.isValid() {
		return true
	} else {
		return false
	}
}

func (s *Series) isValid() {
	if s.childA != nil && 
	s.childA.Parent() == s && 
	s.childB != nil && 
	s.childB.Parent() == s && 
	s.childA.isValid() && 
	s.childB.isValid() {
		return true
	} else {
		return false
	}
}

func (p *Parallel) setI(val float64) error {
	if p.i != 0.0 {
		// I already set
		return error("setI : I already set")
	} else if val == 0.0 {
		return error("setI : I must not be zero")
	} else {
		p.i = val
		// calculate all possible values
		
	}
}