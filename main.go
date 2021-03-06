package main

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

const dbname = "PersonDb"
const collectionName = "Person"

func indexRoute(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("Bİ Hata Oluştu")

	}
	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res.Status(400).SendString("bir hata oluştu")
	}
	defer curr.Close(context.Background())
	var result []bson.M
	curr.All(context.Background(), &result)
	json, _ := json.Marshal(result)
	return res.Status(200).Send(json)
}
func addPerson(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("Hata oluştu Tekrar deneyin")
	}
	var newPoz Poz
	json.Unmarshal([]byte(res.Body()), &newPoz)
	curr, err := collection.InsertOne(context.Background(), newPoz)
	if err != nil {
		return res.Status(400).SendString("Bi Hata oluştu. Tekrar Dene")
	}
	response, err := json.Marshal(curr)
	return res.Status(200).Send(response)

}
func main() {
	app := fiber.New()
	app.Get("/", indexRoute)
	app.Post("/create", addPerson)
	app.Listen(":8000")
}
