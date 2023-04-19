package repository

import (
	"context"
	"github.com/hongdangcseiu/go-back-end/domain"
	"github.com/hongdangcseiu/go-back-end/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postRepository struct {
	database   mongo.Database
	collection string
}

func NewPostRepository(db mongo.Database, collection string) domain.PostRepository {
	return &postRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *postRepository) Create(c context.Context, post *domain.Post) error {
	collection := pr.database.Collection(pr.collection)
	_, err := collection.InsertOne(c, post)

	return err
}

func (pr *postRepository) GetPost(c context.Context) ([]domain.Post, error) {
	collection := pr.database.Collection(pr.collection)

	var posts []domain.Post

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &posts)
	if posts == nil {
		return []domain.Post{}, err
	}

	return posts, err
}

func (pr *postRepository) GetPostByUserID(c context.Context, userID string) ([]domain.Post, error) {
	collection := pr.database.Collection(pr.collection)

	var posts []domain.Post

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return posts, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &posts)
	if posts == nil {
		return []domain.Post{}, err
	}

	return posts, err
}

func (pr *postRepository) GetPostByID(c context.Context, postID string) (domain.Post, error) {

	collection := pr.database.Collection(pr.collection)
	var post domain.Post
	idHex, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return post, err
	}
	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&post)
	return post, err

}
