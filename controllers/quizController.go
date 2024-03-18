package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/naflafadia/go-trivia-app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetQuizzes(c *fiber.Ctx) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://naflafadia:1234@resonance-riddle.hgccwvw.mongodb.net/"))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("Resonance-Riddle").Collection("quizzes")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer cursor.Close(ctx)

	var quizzes []entities.Quiz

	for cursor.Next(ctx) {
		var quiz entities.Quiz
		if err := cursor.Decode(&quiz); err != nil {
			return c.Status(500).JSON(err.Error())
		}
		quizzes = append(quizzes, quiz)
	}
	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(quizzes)
}

func GetQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON("Invalid quiz ID")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://naflafadia:1234@resonance-riddle.hgccwvw.mongodb.net/"))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("Resonance-Riddle").Collection("quizzes")

	var quiz entities.Quiz
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&quiz)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON("Quiz not found")
		}
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(quiz)
}