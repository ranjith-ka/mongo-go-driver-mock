package mongo_go_driver_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestDeleteOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}})
		err := delete(primitive.NewObjectID())
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {
		userCollection = mt.Coll
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}})
		err := delete(primitive.NewObjectID())
		assert.NotNil(t, err)
	})
}
