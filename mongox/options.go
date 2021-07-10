package mongox

import "go.mongodb.org/mongo-driver/mongo/options"

func MakeFindPageOpt(opt *options.FindOptions, pageIndex, pageSize int64) *options.FindOptions {
	if opt == nil {
		opt = options.Find()
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 10 {
		pageSize = 10
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
