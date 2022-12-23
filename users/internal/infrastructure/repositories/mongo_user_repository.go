package repositories

// import mongo driver
import (
	"context"
	"errors"
	"time"

	"users/internal/contracts"
	"users/internal/domain/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName    string = "users"
	mongoQueryTimeout        = 10 * time.Second
)

type mongoUserRepository struct {
	collection *mongo.Collection
}

// NewMongoUserRepository returns a new instance of UserRepository.
func NewMongoUserRepository(db *mongo.Database) contracts.UserRepository {
	return &mongoUserRepository{
		collection: db.Collection("users"),
	}
}

// Get returns a user by id. Returns an error if the user does not exist.
func (r *mongoUserRepository) Get(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// GetAll returns all users.
func (r *mongoUserRepository) GetAll() []models.User {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil
	}

	return users
}

// Create creates a new user. Returns an error if the user already exists.
func (r *mongoUserRepository) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	user.ID = uuid.New()
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return errors.New("user already exists")
	}

	return nil
}

// Update updates an existing user. Returns an error if the user does not exist.
func (r *mongoUserRepository) Update(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": user})
	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

// Delete deletes a user by id. Returns an error if the user does not exist.
func (r *mongoUserRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return errors.New("user not found")
	}

	return nil
}
