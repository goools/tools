package mongox

import "go.mongodb.org/mongo-driver/bson/primitive"

func EmptyOid(oid primitive.ObjectID) bool {
	return oid == primitive.NilObjectID
}
