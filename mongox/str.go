package mongox

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func ConvertSort(sortedFieldList ...string) bson.D {
	result := bson.D{}
	for i := 0; i < len(sortedFieldList); i++ {
		if strings.HasPrefix(sortedFieldList[i], "-") {
			result = append(result, bson.E{Key: sortedFieldList[i][1:], Value: -1})
		} else {
			result = append(result, bson.E{Key: sortedFieldList[i], Value: 1})
		}
	}
	return result
}
