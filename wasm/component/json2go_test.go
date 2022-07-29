package component

import (
	"fmt"
	"testing"
)

func TestJson2Gorm(t *testing.T) {
	json := "{\"x\": 123, \"y\": \"test\", \"z\": false}"
	s := Json2Go(json)
	fmt.Println(s)
}
