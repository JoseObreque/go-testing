package server

import (
	"functional/prey"
	"functional/shark"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position" binding:"required"`
		YPosition float64 `json:"y_position" binding:"required"`
		Speed     float64 `json:"speed" binding:"required"`
	}

	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var req request

		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(400, response{Success: false})
			return
		}

		context.JSON(200, response{Success: true})
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed" required:"true"`
	}

	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var req request

		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(400, response{Success: false})
			return
		}

		context.JSON(200, response{Success: true})
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		// Configure the shark and prey
		h.prey.SetSpeed(140)
		h.shark.Configure([2]float64{0, 300}, 240)

		// Execute the simulation (catch the prey in 3 seconds)
		err, timeToCatch := h.shark.Hunt(h.prey)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response{Success: false, Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, response{Success: true, Message: "Simulation completed", Time: timeToCatch})
	}
}
