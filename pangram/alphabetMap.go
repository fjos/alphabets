package pangram

import (
	"errors"
	"fmt"
	"strings"
	"expvar"
)
var (
	AlphabetsRequested = expvar.NewMap("alphabets")
)

var alphabetMap = map[string]string{
	"latin": "abcdefghijklmnopqrstuvwxyz",
}


// func init() {
	
// }
type Alphabet struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
}

func (alphabet *Alphabet) SetAlphabetContents() error {
	AlphabetsRequested.Add(alphabet.Name, 1)

	if alphabet.Name == "custom" {
		if alphabet.Contents != strings.ToLower(alphabet.Contents) {
			err := errors.New("Custom alphabet is required to only contain lowercase characters.")
			return err
		}
		return nil
	}
	val, ok := alphabetMap[alphabet.Name]

	if !ok {
		err := fmt.Errorf("Selected alphabet '%s' does not exist.", alphabet.Name)
		return err
	}
	alphabet.Contents = val
	return nil
}
