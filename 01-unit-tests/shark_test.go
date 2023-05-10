package hunt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test the Hunt method of the Shark struct. The shark successfully hunts the prey.
func TestSharkHuntsSuccessfully(t *testing.T) {
	// Given
	shark := &Shark{
		hungry: true,
		tired:  false,
		speed:  10,
	}
	prey := &Prey{
		name:  "fish",
		speed: 5,
	}

	// When
	err := shark.Hunt(prey)

	// Then
	// Check that there is no error from the Hunt method
	if assert.NoError(t, err, "the Hunt method should not return an error") {
		// Check that the shark is not hungry anymore
		assert.False(t, shark.hungry, "shark should not be hungry")
		// Check that the shark is tired
		assert.True(t, shark.tired, "shark should be tired")
	}
}

// Test the Hunt method of the Shark struct. The shark cannot hunt because it is tired.
func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Given
	shark := &Shark{
		hungry: true,
		tired:  true,
		speed:  10,
	}
	prey := &Prey{
		name:  "fish",
		speed: 5,
	}

	// When
	err := shark.Hunt(prey)

	// Then
	// Check that there is an error from the Hunt method
	if assert.Error(t, err, "the Hunt method should return an error") {
		// Check that the error message is correct
		assert.ErrorIs(t, err, ErrTired, "the error message is not correct")
	}

}

// Test the Hunt method of the Shark struct. The shark cannot hunt because it is not hungry.
func TestSharkCannotHuntBecauseIsNotHungry(t *testing.T) {
	// Given
	shark := &Shark{
		hungry: false,
		tired:  false,
		speed:  10,
	}
	prey := &Prey{
		name:  "fish",
		speed: 5,
	}

	// When
	err := shark.Hunt(prey)

	// Then
	// Check that there is an error from the Hunt method
	if assert.Error(t, err, "the Hunt method should return an error") {
		// Check that the error message is correct
		assert.ErrorIs(t, err, ErrNotHungry, "the error message is not correct")
	}
}

// Test the Hunt method of the Shark struct. The shark cannot hunt because it cannot reach the prey.
func TestSharkCannotReachThePrey(t *testing.T) {
	// Given
	shark := &Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}
	prey := &Prey{
		name:  "fish",
		speed: 10,
	}

	// When
	err := shark.Hunt(prey)

	// Then
	// Check that there is an error from the Hunt method
	if assert.Error(t, err, "the Hunt method should return an error") {
		// Check that the error message is correct
		assert.ErrorIs(t, err, ErrPreyEscape, "the error message is not correct")
	}
}

// Test the Hunt method of the Shark struct. The Hunt method is called with a nil prey.
func TestSharkHuntNilPrey(t *testing.T) {
	// Given
	shark := &Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}
	var prey *Prey

	// When
	err := shark.Hunt(prey)

	// Then
	// Check that there is an error from the Hunt method and that the error is an ErrNoPrey error
	if assert.Error(t, err, "the Hunt method should return an error") {
		assert.ErrorIs(t, err, ErrNoPrey, "the error message is not correct")
	}
}

// Test all possible errors from the Hunt method of the Shark struct, using a test table approach.
func TestShark_Hunt_Errors(t *testing.T) {
	// Given

	// Define a test table struct to test all possible errors from the Hunt method
	type errorTests struct {
		name        string
		shark       *Shark
		prey        *Prey
		expectedErr error
	}

	// Define the test table
	tests := []errorTests{
		{
			name: "shark is tired",
			shark: &Shark{
				hungry: true,
				tired:  true,
				speed:  10,
			},
			prey: &Prey{
				name:  "fish",
				speed: 5,
			},
			expectedErr: ErrTired,
		},
		{
			name: "shark is not hungry",
			shark: &Shark{
				hungry: false,
				tired:  false,
				speed:  10,
			},
			prey: &Prey{
				name:  "fish",
				speed: 5,
			},
			expectedErr: ErrNotHungry,
		},
		{
			name: "shark cannot reach the prey",
			shark: &Shark{
				hungry: true,
				tired:  false,
				speed:  5,
			},
			prey: &Prey{
				name:  "fish",
				speed: 10,
			},
			expectedErr: ErrPreyEscape,
		},
		{
			name: "prey is nil",
			shark: &Shark{
				hungry: true,
				tired:  false,
				speed:  5,
			},
			prey:        nil,
			expectedErr: ErrNoPrey,
		},
	}

	// Loop over the test table
	for _, test := range tests {
		// When
		err := test.shark.Hunt(test.prey)

		// Then
		// Check that there is an error from the Hunt method
		if assert.Error(t, err, "the Hunt method should return an error") {
			// Check that the error message is correct
			assert.ErrorIs(t, err, test.expectedErr, "the error message is not correct")
		}
	}
}
