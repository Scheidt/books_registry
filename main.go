package main

// go run main.go

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pet_shop_registry/models"
	"pet_shop_registry/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Pet struct {
	ID            int		`json: "id"`
	Name          string	`json: "name"`
	Age           int		`json: "age"`
	Owner         string	`json: "owner"`
	Size          string	`json: "size"`
	Weight        float32	`json: "weight"`
	PaidThisMonth bool		`json: "paid"`
}

type Repository struct {
	DB *gorm.DB
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
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "couldn't create pet"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "pet has been added successfully"})
	return nil
}

func (r *Repository) DeletePet(context *fiber.Ctx) error {
	petModel := models.Pet{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "must provide an ID"})
		return nil
	}

	err := r.DB.Delete(petModel, id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "couldn't delete pet"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "pet has been deleted successfully"})
	return nil
}

func (r *Repository) GetPetByID(context *fiber.Ctx) error { //WIP
	id := context.Params("id")
	petModel := &models.Pet{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "must provide and ID"})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(petModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "no pet with ID " + id})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success in fetching pet",
		"data":    petModel,
	})
	return nil
}

func (r *Repository) GetPet(context *fiber.Ctx) error {
	petModels := &models.Pet{}

	err := r.DB.Find(petModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get pets"})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "pets fetched successfully",
		"data": petModels})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_pets", r.CreatePet)
	api.Delete("/delete_pet/:id", r.DeletePet)
	api.Get("/get_pets/:id", r.GetPetByID)
	api.Get("/pet", r.GetPet)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.EstablishConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigratePets(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
