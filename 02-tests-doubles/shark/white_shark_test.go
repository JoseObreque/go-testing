package shark

import (
	"github.com/stretchr/testify/assert"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"
)

// TestWhiteShark_Hunt_Success tests the success scenario for the Hunt method of the WhiteShark struct.
// The prey does not move, so the shark can catch it.
func TestWhiteShark_Hunt_Success(t *testing.T) {
	// Arrange
	preyStub := prey.NewPreyStub()
	preyStub.GetSpeedFunc = func() float64 {
		return 0
	}

	simulatorMock := simulator.NewCatchSimulatorMock()
	simulatorMock.CanCatchFunc = func(distance float64, speed float64, catchSpeed float64) bool {
		return true
	}
	simulatorMock.GetLinearDistanceFunc = func(position [2]float64) float64 {
		return 0
	}

	shark := CreateWhiteShark(simulatorMock)

	// Act
	err := shark.Hunt(preyStub)

	// Assert
	assert.NoError(t, err)
	assert.True(t, simulatorMock.GetLinearDistanceSpy)
}

// TestWhiteShark_Hunt_Fail_SharkTooSlow tests the fail scenario for the Hunt method of the WhiteShark struct.
// The prey moves faster than the shark, so the shark cannot catch it.
func TestWhiteShark_Hunt_Fail_SharkTooSlow(t *testing.T) {
	// Arrange
	preyStub := prey.NewPreyStub()
	preyStub.GetSpeedFunc = func() float64 {
		return 200
	}

	simulatorMock := simulator.NewCatchSimulatorMock()
	simulatorMock.CanCatchFunc = func(distance float64, speed float64, catchSpeed float64) bool {
		return false
	}
	simulatorMock.GetLinearDistanceFunc = func(position [2]float64) float64 {
		return 100
	}

	shark := CreateWhiteShark(simulatorMock)

	// Act
	err := shark.Hunt(preyStub)

	// Assert
	assert.Error(t, err)
	assert.True(t, simulatorMock.GetLinearDistanceSpy)
}

// TestWhiteShark_Hunt_Fail_PreyTooFarAway tests the fail scenario for the Hunt method of the WhiteShark struct.
// The prey is too far away from the shark, so the shark cannot catch it.
func TestWhiteShark_Hunt_Fail_PreyTooFarAway(t *testing.T) {
	// Arrange
	preyStub := prey.NewPreyStub()
	preyStub.GetSpeedFunc = func() float64 {
		return 100
	}

	simulatorMock := simulator.NewCatchSimulatorMock()
	simulatorMock.CanCatchFunc = func(distance float64, speed float64, catchSpeed float64) bool {
		return false
	}
	simulatorMock.GetLinearDistanceFunc = func(position [2]float64) float64 {
		return 1000
	}

	shark := CreateWhiteShark(simulatorMock)

	// Act
	err := shark.Hunt(preyStub)

	// Assert
	assert.Error(t, err)
	assert.True(t, simulatorMock.GetLinearDistanceSpy)
}
