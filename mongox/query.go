package mongox

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeQueryByID(oid primitive.ObjectID) bson.D {
	return bson.D{
		{
			Key: "_id",
			Value: oid,
		},
	}
}
