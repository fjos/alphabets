package pangram

import (
	"fmt"
)

var alphabetMap = map[string]string{
	"latin": "abcdefghijklmnopqrstuvwxyz",
}

func GetAlphabet(unicodeName string) (string, error) {
	val, ok := alphabetMap[unicodeName]
	if !ok {
		err := fmt.Errorf("Selected alphabet '%s' does not exist.", unicodeName)
		return val, err
	}
	return val, nil
}
