package shark

import (
	"github.com/stretchr/testify/assert"
	"integrationtests/pkg/storage"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"
)

func TestWhiteShark_Hunt(t *testing.T) {
	t.Run("shark successfully catch the prey", func(t *testing.T) {
		// Arrange
		dataStorage := storage.NewStorageMock()
		dataStorage.Mock.On("GetValue", "white_shark_speed").Return(144.0)
		dataStorage.Mock.On("GetValue", "white_shark_x").Return(5.0)
		dataStorage.Mock.On("GetValue", "white_shark_y").Return(5.0)
		dataStorage.Mock.On("GetValue", "tuna_speed").Return(0.0)

		testSimulator := simulator.NewCatchSimulator(200)

		shark := CreateWhiteShark(testSimulator, dataStorage)
		tuna := prey.CreateTuna(dataStorage)

		// Act
		err := shark.Hunt(tuna)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("shark could not catch the prey (shark too slow)", func(t *testing.T) {
		// Assert
		dataStorage := storage.NewStorageMock()
		dataStorage.Mock.On("GetValue", "white_shark_speed").Return(10.0)
		dataStorage.Mock.On("GetValue", "white_shark_x").Return(5.0)
		dataStorage.Mock.On("GetValue", "white_shark_y").Return(5.0)
		dataStorage.Mock.On("GetValue", "tuna_speed").Return(50.0)

		testSimulator := simulator.NewCatchSimulator(200)

		shark := CreateWhiteShark(testSimulator, dataStorage)
		tuna := prey.CreateTuna(dataStorage)

		// Act
		err := shark.Hunt(tuna)

		// Assert
		assert.Error(t, err)
	})

	t.Run("shark could not catch the prey (shark too far)", func(t *testing.T) {
		// Assert
		dataStorage := storage.NewStorageMock()
		dataStorage.Mock.On("GetValue", "white_shark_speed").Return(10.0)
		dataStorage.Mock.On("GetValue", "white_shark_x").Return(0.0)
		dataStorage.Mock.On("GetValue", "white_shark_y").Return(1500.0)
		dataStorage.Mock.On("GetValue", "tuna_speed").Return(5.0)

		testSimulator := simulator.NewCatchSimulator(200)

		shark := CreateWhiteShark(testSimulator, dataStorage)
		tuna := prey.CreateTuna(dataStorage)

		// Act
		err := shark.Hunt(tuna)

		// Assert
		assert.Error(t, err)
	})
}
