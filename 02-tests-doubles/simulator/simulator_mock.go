package simulator

// CatchSimulatorMock is a mock implementation of CatchSimulator interface.
type CatchSimulatorMock struct {
	// CanCatchMock is a struct representing the CatchSimulator.CanCatch inputs and outputs
	CanCatchMock struct {
		Distance   float64
		Speed      float64
		CatchSpeed float64
		Return     bool
	}
	// GetLinearDistanceMock is a struct representing the CatchSimulator.GetLinearDistance inputs and outputs
	GetLinearDistanceMock struct {
		Position [2]float64
		Return   float64
	}
	// CanCatchSpy is a spy (boolean value) on CatchSimulator.CanCatch, checking if the method was called
	CanCatchSpy bool
	// GetLinearDistanceSpy is a spy (boolean value) on CatchSimulator.GetLinearDistance, checking if the
	// method was called
	GetLinearDistanceSpy bool
}

// CanCatch is a mock implementation of CatchSimulator.CanCatch
func (s *CatchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	// Set the spy to true, indicating that the method was called
	s.CanCatchSpy = true

	// Set the mock CanCatchMock inputs
	s.CanCatchMock.Distance = distance
	s.CanCatchMock.Speed = speed
	s.CanCatchMock.CatchSpeed = catchSpeed

	// Return the mock CanCatchMock return value
	return s.CanCatchMock.Return
}

// GetLinearDistance is a mock implementation of CatchSimulator.GetLinearDistance
func (s *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	// Set the spy to true, indicating that the method was called
	s.GetLinearDistanceSpy = true

	// Set the mock GetLinearDistanceMock inputs
	s.GetLinearDistanceMock.Position = position

	// Return the mock GetLinearDistanceMock return value
	return s.GetLinearDistanceMock.Return
}

// NewCatchSimulatorMock return a new CatchSimulatorMock struct
func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}
