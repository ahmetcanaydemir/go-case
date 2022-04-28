package src_test

import (
	"bytes"
	"encoding/json"
	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/src"
	"github.com/gofiber/fiber/v2"
	"io"
	"math/rand"
	"net/http/httptest"
	"sort"
	"testing"
)

func Test_PostFindPositionController(t *testing.T) {

	var bigSortedArray []int
	bigSortedArray = append(bigSortedArray, 0)

	for i := 0; i < 700000; i++ {
		bigSortedArray = append(bigSortedArray, rand.Intn(100_000))
	}
	sort.Ints(bigSortedArray)

	bigSortedArrStr, _ := json.Marshal(bigSortedArray)
	var tests = []struct {
		name           string
		bodyjson       string
		expectedResult int
	}{
		{"found request", `{"array":[1,1,2,3,3,4,5,6,6,19,20],"search":3}`, 4},
		{"not found", `{"array":[1,1,2,3,3,4,5,6,6,19,20],"search":8}`, -1},
		{"big array found", `{"array":` + string(bigSortedArrStr) + `,"search":0}`, 1},
		{"big array not found", `{"array":` + string(bigSortedArrStr) + `,"search":100001}`, -1},
		{"empty array", `{"array":[],"search":3}`, -1},
		{"negative search", `{"array":[1,1,2,3,3,4,5,6,6,19,20],"search":-14}`, -1},
	}

	app := fiber.New()
	src.NewAlgorithmController(app)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest("POST", "/find-position", bytes.NewBuffer([]byte(tt.bodyjson)))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)
			defer resp.Body.Close()

			var bodyString string
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyString = string(bodyBytes)
			var response int
			json.Unmarshal([]byte(bodyString), &response)

			if response != tt.expectedResult {
				t.Errorf("PostAlgorithmController(%s) got %v, want %v", tt.name, resp, tt.expectedResult)
			}
		})
	}
}
