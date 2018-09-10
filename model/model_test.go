package model
import (
	"strings"
	"testing"
	"github.com/fjos/alphabets/pangram"
)


var checkIfPangramTests = []struct {
	input    string // input
	alphabet pangram.Alphabet
	expectedStatus int   // expected 400 status
	expectedResponseState bool 
}{
	{"", pangram.Alphabet{"latin", ""}, 200, false},                              //empty
	{"abcdefghijklmnopqrstuvwxyz", pangram.Alphabet{"latin", ""}, 200, true},                              //empty
	{"abcdefghijklmnopqrstuvwxyz", pangram.Alphabet{"egypt", ""}, 404, false},                              //empty
	{"abcdefghijklmnopqrstuvwxyz", pangram.Alphabet{"custom", "ab"}, 200, true},                              //custom alphabet
}

func TestCheckIfPangram(t *testing.T) {
	for _, tt := range checkIfPangramTests {
		handle := strings.NewReader(tt.input)
		response := CheckIfPangram(tt.alphabet, handle)
		if response.Status != tt.expectedStatus {
			t.Errorf("CheckIfPangram( alphabet:'%s', input : '%s'): STATUS expected %d, actual %d", tt.alphabet.Name, tt.input, tt.expectedStatus, response.Status)
		}
		if response.Data.Pangram != tt.expectedResponseState{
			t.Errorf("CheckIfPangram( alphabet:'%s', input : '%s'): RESPONSE expected %t, actual %t", tt.alphabet.Name, tt.input, tt.expectedResponseState, response.Data.Pangram)

		}
	}

}