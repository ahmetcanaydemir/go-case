package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var PersonsDir = "persons"

type Person struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	BirthDay    time.Time `json:"birthDay"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
}

type PersonController struct {
}

func NewPersonController(app *fiber.App) {
	ct := &PersonController{}
	app.Get("/persons", ct.GetPersons)
}

func (ct PersonController) GetPersons(c *fiber.Ctx) error {
	var persons []Person

	files, err := ioutil.ReadDir(PersonsDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".json") {
			continue
		}

		jsonFile, _ := ioutil.ReadFile(filepath.Join(PersonsDir, fileName))
		person := Person{}

		err = json.Unmarshal(jsonFile, &person)
		if err != nil {
			log.Fatal(err)
		}

		persons = append(persons, person)
	}

	if persons == nil {
		return c.Next()
	}

	return c.JSON(persons)
}
