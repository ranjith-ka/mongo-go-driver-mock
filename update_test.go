package mongo_go_driver_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFindOneAndUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		userData := user{
			ID:    primitive.NewObjectID(),
			Name:  "john",
			Email: "john.doe@test.com",
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "_id", Value: userData.ID},
				{Key: "name", Value: userData.Name},
				{Key: "email", Value: userData.Email},
			}},
		})

		updatedUser, err := update(userData)

		assert.Nil(t, err)
		assert.Equal(t, &userData, updatedUser)
	})
}
