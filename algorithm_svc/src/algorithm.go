package src

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type FindPositionRequest struct {
	Array  []int `json:"array"`
	Search int   `json:"search"`
}

type AlgorithmController struct {
}

func NewAlgorithmController(app *fiber.App) {
	ct := &AlgorithmController{}
	app.Post("/find-position", ct.PostFindPosition)
}

func (ct AlgorithmController) PostFindPosition(c *fiber.Ctx) error {
	var body FindPositionRequest
	if err := c.BodyParser(&body); err != nil {
		log.Printf("error while decoding body: %v", err)
	}

	result := binarySearch(body.Search, body.Array)

	if result == -1 {
		c.Status(fiber.StatusNotFound)
	}

	return c.JSON(result)
}

func binarySearch(search int, array []int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] < search {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(array) || array[low] != search {
		return -1
	}

	return low + 1
}
