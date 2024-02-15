package main

// go run main.go pet.go petController.go

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/mod/sumdb/storage"
	"gorm.io/gorm"
)

type Pet struct {
	name    string
	age  	int
	owner   string
	size    string
	paidThisMonth bool
}

type Repository struct {
	DB *gorm.db
}

func (r *Repository) CreatePet(context *fiber.Ctx) error {
	pet := Pet{}
	err := context.BodyParser(&pet)
	if err != nil {
        context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
        return err
	}

	err = r.DB.Create(&pet).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"couldn't create pet"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message":"pet has been added"})
	return nil
}

func (r *Repository) GetPet(context *fiber.Ctx) error{
	petModels :=&models.Pet{}

	err := r.DB.Find(petModels).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"could not get pets"})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{	"message": "pets fetched successfully",
													"data": petModels})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_pets", r.CreatePet)
	api.Delete("delete_pet/:id", r.DeletePet)
	api.Get("/get_pets/:id", r.GetPetByID)
	api.Get("/pet", r.GetPet)
}

//type pet struct {
//	title    string
//	pubYear  int
//	author   string
//	genre    [2]string
//	borrowed bool
//}

func getStringInput(texto string) string {
	fmt.Print(texto)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ReplaceAll(input, "'", "''")
	return input
}

func getIntInput(message string) int {
	var userInput int
	for {
		fmt.Print(message)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Only integer numbers are allowed, please, try again: ")
			continue
		}
		userInput = value
		break
	}
	return userInput
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
