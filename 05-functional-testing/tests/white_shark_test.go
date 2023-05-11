package tests

import (
	"functional/prey"
	"functional/shark"
	"functional/simulator"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestWhiteShark_Hunt_SharkTooSlow sets a prey faster than the shark and checks that the shark cannot catch it.
func TestWhiteShark_Hunt_SharkTooSlow(t *testing.T) {
	// Arrange
	testSimulator := simulator.NewCatchSimulator(10)
	testPrey := prey.CreateTuna()
	testPrey.SetSpeed(10)

	testShark := shark.CreateWhiteShark(testSimulator)
	testShark.Configure([2]float64{10, 20}, 5)

	// Act
	err, timeToCatch := testShark.Hunt(testPrey)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0.0, timeToCatch)
}

// TestWhiteShark_Hunt_SharkTooFarAway sets a prey slower than the shark and checks that the shark cannot catch it
// because it is too far away.
func TestWhiteShark_Hunt_SharkTooFarAway(t *testing.T) {
	// Arrange
	testSimulator := simulator.NewCatchSimulator(10)
	testPrey := prey.CreateTuna()
	testPrey.SetSpeed(5)

	testShark := shark.CreateWhiteShark(testSimulator)
	testShark.Configure([2]float64{0, 500}, 10)

	// Act
	err, timeToCatch := testShark.Hunt(testPrey)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0.0, timeToCatch)
}

// TestWhiteShark_Hunt_Success sets a prey slower than the shark and checks that the shark can catch it
// in exactly 24 seconds.
func TestWhiteShark_Hunt_Success(t *testing.T) {
	// Arrange
	testSimulator := simulator.NewCatchSimulator(30)
	testPrey := prey.CreateTuna()
	testPrey.SetSpeed(5)

	testShark := shark.CreateWhiteShark(testSimulator)
	testShark.Configure([2]float64{0, 240}, 15)

	// Act
	err, timeToCatch := testShark.Hunt(testPrey)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 24.0, timeToCatch)
}
