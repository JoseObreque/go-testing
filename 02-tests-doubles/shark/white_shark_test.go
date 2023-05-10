package shark

import (
	"github.com/stretchr/testify/assert"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"
)

// TestWhiteShark_Hunt_Success tests the success scenario for the Hunt method of the WhiteShark struct.
func TestWhiteShark_Hunt_Success(t *testing.T) {
	//-------------------------------------------------------------------------
	// Arrange
	//-------------------------------------------------------------------------

	// Create a new prey stub
	preyStub := prey.NewPreyStub()

	// The prey stands still
	preyStub.GetSpeedFunc = func() float64 {
		return 0
	}

	// Create a new mocked simulator
	simulatorMock := simulator.NewCatchSimulatorMock()

	// The prey is 0 meters away from the shark (oh no!)
	simulatorMock.GetLinearDistanceMock.Position = [2]float64{0, 0}
	simulatorMock.GetLinearDistanceMock.Return = 0

	// The shark is 10 meters per second faster than the prey
	simulatorMock.CanCatchMock.Distance = 0
	simulatorMock.CanCatchMock.Speed = preyStub.GetSpeed() // 0
	simulatorMock.CanCatchMock.CatchSpeed = 10
	simulatorMock.CanCatchMock.Return = true

	// Create a new white shark
	whiteShark := CreateWhiteShark(simulatorMock)

	//-------------------------------------------------------------------------
	// Act
	//-------------------------------------------------------------------------

	// Hunt the prey
	err := whiteShark.Hunt(preyStub)

	//-------------------------------------------------------------------------
	// Assert
	//-------------------------------------------------------------------------

	// Check that there is no error returned from the Hunt method
	assert.NoError(t, err)

	// Check that the GetLinearDistance method was called once
	assert.True(t, simulatorMock.GetLinearDistanceSpy)
}

// TestWhiteShark_Hunt_Failure tests the failure scenario for the Hunt method of the WhiteShark struct.
// In this case, the shark cannot catch the prey because it is too slow.
func TestWhiteShark_Hunt_Failure_SharkTooSlow(t *testing.T) {
	//-------------------------------------------------------------------------
	// Arrange
	//-------------------------------------------------------------------------

	// Create a new prey stub
	preyStub := prey.NewPreyStub()

	// The prey moves very fast
	preyStub.GetSpeedFunc = func() float64 {
		return 300
	}

	// Create a new mocked simulator
	simulatorMock := simulator.NewCatchSimulatorMock()

	// The prey is 100 meters away from the shark
	simulatorMock.GetLinearDistanceMock.Position = [2]float64{0, 100}
	simulatorMock.GetLinearDistanceMock.Return = 100

	// The shark is 250 meters per second slower than the prey
	simulatorMock.CanCatchMock.Distance = 100
	simulatorMock.CanCatchMock.Speed = preyStub.GetSpeed()            // 300
	simulatorMock.CanCatchMock.CatchSpeed = preyStub.GetSpeed() - 250 // 50
	simulatorMock.CanCatchMock.Return = false

	// Create a new white shark
	whiteShark := CreateWhiteShark(simulatorMock)

	//-------------------------------------------------------------------------
	// Act
	//-------------------------------------------------------------------------

	// Hunt the prey
	err := whiteShark.Hunt(preyStub)

	//-------------------------------------------------------------------------
	// Assert
	//-------------------------------------------------------------------------

	// Check that there is an error returned from the Hunt method
	assert.Error(t, err)

	// Check that the GetLinearDistance method was called once
	assert.True(t, simulatorMock.GetLinearDistanceSpy)
}

// TestWhiteShark_Hunt_Failure tests the failure scenario for the Hunt method of the WhiteShark struct.
// In this case, the shark cannot catch the prey because it is too far away, even though it is fast enough.
func TestWhiteShark_Hunt_Failure_TooFarAway(t *testing.T) {
	//-------------------------------------------------------------------------
	// Arrange
	//-------------------------------------------------------------------------

	// Create a new prey stub
	preyStub := prey.NewPreyStub()

	// The prey moves very slow
	preyStub.GetSpeedFunc = func() float64 {
		return 1
	}

	// Create a new mocked simulator
	simulatorMock := simulator.NewCatchSimulatorMock()

	// The prey is 1000 meters away from the shark
	simulatorMock.GetLinearDistanceMock.Position = [2]float64{0, 1000}
	simulatorMock.GetLinearDistanceMock.Return = 1000

	// The shark is 10 meters per second faster than the prey
	simulatorMock.CanCatchMock.Distance = 10
	simulatorMock.CanCatchMock.Speed = preyStub.GetSpeed() + 10 // 11
	simulatorMock.CanCatchMock.CatchSpeed = 10
	simulatorMock.CanCatchMock.Return = false

	// Create a new white shark
	whiteShark := CreateWhiteShark(simulatorMock)

	//-------------------------------------------------------------------------
	// Act
	//-------------------------------------------------------------------------

	// Hunt the prey
	err := whiteShark.Hunt(preyStub)

	//-------------------------------------------------------------------------
	// Assert
	//-------------------------------------------------------------------------

	// Check that there is an error returned from the Hunt method
	assert.Error(t, err)

	// Check that the GetLinearDistance method was called once
	assert.True(t, simulatorMock.GetLinearDistanceSpy)
}
