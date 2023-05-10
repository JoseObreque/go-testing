package prey

// PreyStub is a stub for Prey interface
type PreyStub struct {
	GetSpeedFunc func() float64
}

// GetSpeed return the result of GetSpeedFunc of the PreyStub struct
func (s *PreyStub) GetSpeed() float64 {
	return s.GetSpeedFunc()
}

// NewPreyStub return a new PreyStub struct
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}
