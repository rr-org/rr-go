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

func GetAvatars(c *fiber.Ctx) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://naflafadia:1234@resonance-riddle.hgccwvw.mongodb.net/"))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("Resonance-Riddle").Collection("avatar")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if count == 0 {
		return c.JSON("No avatars found")
	}

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer cursor.Close(ctx)

	var avatars []entities.Avatar

	for cursor.Next(ctx) {
		var avatar entities.Avatar
		if err := cursor.Decode(&avatar); err != nil {
			return c.Status(500).JSON(err.Error())
		}
		avatars = append(avatars, avatar)
	}
	if err := cursor.Err(); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(avatars)
}

func GetAvatar(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON("Invalid avatar ID")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://naflafadia:1234@resonance-riddle.hgccwvw.mongodb.net/"))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("Resonance-Riddle").Collection("avatar")

	var avatar entities.Avatar
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&avatar)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON("Avatar not found")
		}
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(avatar)
}