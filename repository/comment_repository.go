package repository

import (
	"context"
	"github.com/hongdangcseiu/go-back-end/domain"
	"github.com/hongdangcseiu/go-back-end/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentRepository struct {
	database   mongo.Database
	collection string
}

func (c2 *commentRepository) GetCommentByID(c context.Context, commentID string) (domain.Comment, error) {
	collection := c2.database.Collection(c2.collection)
	var comment domain.Comment
	idHex, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return comment, err
	}
	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&comment)
	return comment, err
}

func (c2 *commentRepository) Create(c context.Context, comment *domain.Comment) error {
	collection := c2.database.Collection(c2.collection)
	_, err := collection.InsertOne(c, comment)
	return err
}

func (c2 *commentRepository) Edit(c context.Context, commentID string, comment *domain.Comment) error {
	collection := c2.database.Collection(c2.collection)

	commentObjectID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": commentObjectID}
	update := bson.M{
		"$set": bson.M{
			"content": comment.Content,
		},
	}

	_, err = collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c2 *commentRepository) GetCommentByPostID(c context.Context, postID string) ([]domain.Comment, error) {
	collection := c2.database.Collection(c2.collection)

	var comments []domain.Comment

	idHex, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return comments, err
	}

	cursor, err := collection.Find(c, bson.M{"postID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &comments)
	if comments == nil {
		return []domain.Comment{}, err
	}

	return comments, err
}

func (c2 *commentRepository) Delete(c context.Context, comment *domain.Comment) error {
	collection := c2.database.Collection(c2.collection)

	_, err := collection.DeleteOne(c, comment)
	if err != nil {
		return err
	}
	return nil
}

func NewCommentRepository(db mongo.Database, collection string) domain.CommentRepository {
	return &commentRepository{
		database:   db,
		collection: collection,
	}
}
