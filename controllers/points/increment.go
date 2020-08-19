package points

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/vpassanisi/Project-S/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Increment //
// @desc adds a point to a post or comment
// @route POST /api/v1/points/increment/?type&id
// @access Private
func Increment(c *fiber.Ctx) {

	targetID, idErr := primitive.ObjectIDFromHex(c.Query("id"))
	if idErr != nil {
		c.Status(400).JSON(respondM{
			Success: false,
			Message: "please include a valid target id",
		})
		return
	}

	userID, idErr := primitive.ObjectIDFromHex(c.Locals("id").(string))
	if idErr != nil {
		c.Status(400).JSON(respondM{
			Success: false,
			Message: "Bad user id",
		})
		return
	}

	var collection *mongo.Collection

	if c.Query("type") == "comment" {
		collection = config.GetCollection("Comments")
	} else if c.Query("type") == "post" {
		collection = config.GetCollection("Posts")
	} else {
		c.Status(400).JSON(respondM{
			Success: false,
			Message: "please give a type (comment or post)",
		})
		return
	}

	opts := options.Find().SetLimit(1)

	targetCur, findErr := collection.Find(c.Context(), bson.M{"_id": targetID}, opts)
	if findErr != nil {
		c.Status(500).JSON(respondM{
			Success: false,
			Message: "There was a problem with the query",
		})
		return
	}

	if targetCur.RemainingBatchLength() == 0 {
		c.Status(400).JSON(respondM{
			Success: false,
			Message: "There is no target with that id",
		})
		return
	}

	pointsCollection := config.GetCollection("Points")

	pointCur, findPointErr := pointsCollection.Find(c.Context(), bson.M{
		"user":   userID,
		"target": targetID,
	}, opts)
	if findPointErr != nil {
		c.Status(400).JSON(respondM{
			Success: false,
			Message: "There was a problem querying for the point",
		})
		return
	}

	point := point{}

	if pointCur.RemainingBatchLength() == 0 {
		point.Active = true
		point.Target = targetID
		point.User = userID
		point.UpdatedAt = time.Now().Unix()
		_, InsertOneErr := pointsCollection.InsertOne(c.Context(), point)
		if InsertOneErr != nil {
			c.Status(400).JSON(respondM{
				Success: false,
				Message: "There was a problem adding the point to the database",
			})
			return
		}

		if c.Query("type") == "post" {
			incrementPost(c, collection, targetID)
		} else {
			incrementComment(c, collection, targetID)
		}
		return
	}

	defer pointCur.Close(c.Context())
	if pointCur.TryNext(c.Context()) {
		if err := pointCur.Decode(&point); err != nil {
			c.Status(500).JSON(respondM{
				Success: false,
				Message: "there was a problem decoding the point cursor",
			})
			return
		}
	}

	if err := pointCur.Err(); err != nil {
		c.Status(500).JSON(respondM{
			Success: false,
			Message: "there was a problem decoding the point cursor",
		})
		return
	}

	if !point.Active {
		pointErr := pointsCollection.FindOneAndUpdate(c.Context(), bson.M{"_id": point.ID}, bson.M{"$set": bson.M{"active": true}}).Decode(&point)
		if pointErr != nil {
			c.Status(500).JSON(respondM{
				Success: false,
				Message: "there was a problem updating the point",
			})
			return
		}

		if c.Query("type") == "post" {
			incrementPost(c, collection, targetID)
		} else {
			incrementComment(c, collection, targetID)
		}
		return
	}

	c.Status(400).JSON(respondM{
		Success: false,
		Message: "you already up voted that",
	})
	return
}

func incrementPost(c *fiber.Ctx, postsCollection *mongo.Collection, targetID primitive.ObjectID) {

	post := post{}

	decrementErr := postsCollection.FindOneAndUpdate(c.Context(), bson.M{"_id": targetID}, bson.M{"$inc": bson.M{"points": 1}}).Decode(&post)
	if decrementErr != nil {
		c.Status(500).JSON(respondM{
			Success: false,
			Message: "there was a problem decrementing the target",
		})
		return
	}

	post.Points = post.Points + 1

	c.Status(200).JSON(respondP{
		Success: true,
		Data:    post,
	})
}

func incrementComment(c *fiber.Ctx, commentsCollection *mongo.Collection, targetID primitive.ObjectID) {

	comment := comment{}

	decrementErr := commentsCollection.FindOneAndUpdate(c.Context(), bson.M{"_id": targetID}, bson.M{"$inc": bson.M{"points": 1}}).Decode(&comment)
	if decrementErr != nil {
		c.Status(500).JSON(respondM{
			Success: false,
			Message: "there was a problem decrementing the target",
		})
		return
	}

	comment.Points = comment.Points + 1

	c.Status(200).JSON(respondC{
		Success: true,
		Data:    comment,
	})
}
