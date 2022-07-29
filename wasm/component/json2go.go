package component

import "github.com/m-zajac/json2go"

func Json2Go(str string) string {
	parser := json2go.NewJSONParser("NameXXX")
	parser.FeedBytes([]byte(str))

	res := parser.String()
	return res
}
