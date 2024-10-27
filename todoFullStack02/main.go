package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"isCompleted"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello world! Once again!")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	mongodbURI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongodbURI)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas successfully!")

	collection = client.Database("go_react_todo").Collection("todos")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Server is up!"})
	})

	app.Get("/api/todos", getAllTodos)

	app.Post("/api/todos", createTodo)

	app.Patch("/api/todos/:todoId", updateTodo)

	app.Delete("/api/todos/:todoId", deleteTodo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	app.Listen("0.0.0.0:" + port)
}

func getAllTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(fiber.Map{"msg": "Success!", "todos": todos, "total": len(todos)})
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required!"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	todoId := c.Params("todoId")
	objectId, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid Todo id!",
		})
	}

	filter := bson.M{"_id": objectId}

	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func deleteTodo(c *fiber.Ctx) error {
	todoId := c.Params("todoId")
	objectId, err := primitive.ObjectIDFromHex(todoId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid Todo id!",
		})
	}

	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
