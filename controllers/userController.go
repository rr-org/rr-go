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

func GetUsers(c *fiber.Ctx) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://naflafadia:1234@resonance-riddle.hgccwvw.mongodb.net/"))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("Resonance-Riddle").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var users []entities.User

	for cursor.Next(ctx) {
		var user entities.User
		if err := cursor.Decode(&user); err != nil {
			return err
		}

		// if user.Avatar == nil {
		// 	user.Avatar = []entities.Avatar{}
		// }

		// avatarCollection := client.Database("Resonance-Riddle").Collection("avatar")
		// avatarCursor, err := avatarCollection.Find(ctx, bson.M{"_id": bson.M{"$in": user.Avatar}})
		// if err != nil {
		// 	return err
		// }
		// defer avatarCursor.Close(ctx)

		// var avatars []entities.Avatar
		// for avatarCursor.Next(ctx) {
		// 	var avatar entities.Avatar
		// 	if err := avatarCursor.Decode(&avatar); err != nil {
		// 		return err
		// 	}
		// 	avatars = append(avatars, avatar)
		// }

		// user.Avatar = avatars
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return err
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON("Invalid user ID")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://naflafadia:1234@resonance-riddle.hgccwvw.mongodb.net/"))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("Resonance-Riddle").Collection("users")

	var user entities.User
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON("User not found")
		}
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(user)
}
