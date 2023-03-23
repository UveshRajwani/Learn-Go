package main

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var objectIDFromHex = func(hex string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		log.Fatal(err)
	}
	return objectID
}

type singeNamazModel struct {
	Id    primitive.ObjectID `bson:"_id" json:"_id"`
	Name  string             `json:"name"`
	Start time.Time          `json:"start"`
	End   time.Time          `json:"end"`
	City  string             `json:"city"`
}
type singeMasjidModel struct {
	Id            string `bson:"_id" json:"_id"`
	Name          string `json:"name"`
	MaplLnk       string `json:"map_link" bson:"map_link"`
	Img           string `json:"img"`
	BackgroundImg string `json:"background_img" bson:"background_img"`
	City          string `json:"city"`
	Area          string `json:"area"`
	Timing        map[string]struct {
		AzanTime   time.Time `json:"azan_time" bson:"azan_time"`
		JammatTime time.Time `json:"jammat_time" bson:"jammat_time"`
	} `json:"timing"`
}
type singeConfigModel struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	City       string             `json:"city"`
	IsRamzan   bool               `json:"isRamzan"`
	IsLocation bool               `json:"isLocation"`
	IsEid      bool               `json:"isEid"`
	ChangeDate int                `json:"changeDate"`
	TimeTable  string             `json:"timeTable"`
}

func main() {
	app := fiber.New(
		fiber.Config{
			Prefork:     true,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)
	app.Use(logger.New())
	app.Get("/getNamaz/:city/:namaz", getNamaz)
	app.Get("/getMasjid/:id/", getMasjid)
	app.Get("/getConfig/:city/", getConfig)
	err := app.Listen(":3000")
	checkNilErr(err)
	//getNamaz()
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getNamaz(c *fiber.Ctx) error {
	namazHandler, err := getMongoDbCollection("Namazdb", "namaz_time")
	var namaz singeNamazModel
	err = namazHandler.FindOne(context.TODO(), bson.D{
		{"name", c.Params("namaz")},
		{"city", c.Params("city")},
	}).Decode(&namaz)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Namaz Not Found"})
	}
	checkNilErr(err)
	return c.JSON(namaz)
}

func getConfig(c *fiber.Ctx) error {
	namazHandler, err := getMongoDbCollection("Namazdb", "config")
	var config singeConfigModel
	err = namazHandler.FindOne(context.TODO(), bson.D{
		{"city", c.Params("city")},
	}).Decode(&config)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Config Not Found"})
	}
	checkNilErr(err)
	return c.JSON(config)
}
func getMasjid(c *fiber.Ctx) error {

	namazHandler, err := getMongoDbCollection("Namazdb", "masjids")
	var masjid singeMasjidModel

	err = namazHandler.FindOne(context.TODO(), bson.D{
		{"_id", objectIDFromHex(c.Params("id"))},
	}).Decode(&masjid)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "Masjid Not Found"})
	}
	checkNilErr(err)
	return c.JSON(masjid)
}

func GetMongoDbConnection() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://uvesh:uvesh786@167.86.86.55:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
