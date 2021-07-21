package mongox

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeFindPageOpt(opt *options.FindOptions, pageIndex, pageSize int64) *options.FindOptions {
	if opt == nil {
		opt = options.Find()
	}
	if pageIndex <= 0 || pageSize <= 0 {
		return opt
	}
	opt.SetLimit(pageSize)
	if pageIndex > 1 {
		skip := (pageIndex - 1) * pageSize
		opt.SetSkip(skip)
	}
	return opt
}

func MakeSortedFieldOpt(opt *options.FindOptions, sortedField string) *options.FindOptions {
	if opt == nil {
		opt = options.Find()
	}
	sortedField = strings.Trim(sortedField, " ")
	if sortedField == "" {
		sortedField = "-_id"
	}
	if sortedField != "" {
		opt.SetSort(ConvertSort(sortedField))
	}
	return opt
}

func MakeReturnAfter(opt *options.FindOneAndUpdateOptions) *options.FindOneAndUpdateOptions {
	if opt == nil {
		opt = options.FindOneAndUpdate()
	}
	opt.SetReturnDocument(options.After)
	return opt
}

func MakeRegexFilter(fieldName, queryStr string) bson.E {
	return bson.E{
		Key:   fieldName,
		Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: queryStr, Options: "i"}}},
	}
}
