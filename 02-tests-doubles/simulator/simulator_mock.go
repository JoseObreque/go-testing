package simulator

// CatchSimulatorMock is a mock implementation of CatchSimulator interface.
type CatchSimulatorMock struct {
	CanCatchFunc          func(distance float64, speed float64, catchSpeed float64) bool
	GetLinearDistanceFunc func(position [2]float64) float64
	GetLinearDistanceSpy  bool
}

// CanCatch is a mock implementation of CatchSimulator.CanCatch
func (s *CatchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	return s.CanCatchFunc(distance, speed, catchSpeed)
}

// GetLinearDistance is a mock implementation of CatchSimulator.GetLinearDistance
func (s *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	s.GetLinearDistanceSpy = true
	return s.GetLinearDistanceFunc(position)
}

// NewCatchSimulatorMock return a new CatchSimulatorMock struct
func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}
