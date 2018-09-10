package pangram

import (
	// "log"
	// "os"
	"strings"
	"testing"
)

var pangramTests = []struct {
	input    string // input
	expected bool   // expected result
}{
	{"", false},                              //empty
	{"abcdefghijklmnopqrstuvwxyz", true},     //full
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", true},     //allcaps
	{"abcdefghijklmNOPQRSTUVWXYZ", true},     //mixedCase
	{"abcdefghijklmnopqrstuvwxy", false},     //missing one
	{"abcdefgrstuvwxyz", false},              //missing many
	{"abcdefghijklmn0000opqrstuvwxyz", true}, //numbers characters
	{"abcde平仮名fgrstuvwxyz", false},           //special chars missing
	{"平仮名abcdefghijklmNOPQRSTUVWXYZ", true},  //special chars contains all
	{"zyxwvuabcdefghijklmnopqrstuv", true},   //reversed order
	{"\x00\x01abs", false},
}

var pangramCustomAlphabetTests = []struct {
	input string //input
	expected bool // output
	alphabet string //al {phabet to test against
}{
	{"平仮名abcdefghijklmNOPQRSTUVWXYZ", true, "仮名"},
	{"平仮abcdefghijklmNOPQRSTUVWXYZ", false, "仮名"},
}


func TestIsPangram(t *testing.T) {
	for _, tt := range pangramTests {
		handle := strings.NewReader(tt.input)
		actual, _ := IsPangram(handle, "abcdefghijklmnopqrstuvwxyz", 64)
		if actual != tt.expected {
			t.Errorf("IsPangram(%s): expected %t, actual %t", tt.input, tt.expected, actual)
		}
	}
	for _, tt := range pangramCustomAlphabetTests {
		handle := strings.NewReader(tt.input)
		actual, _ := IsPangram(handle, tt.alphabet, 64)
		if actual != tt.expected {
			t.Errorf("IsPangram(%s): expected %t, actual %t", tt.input, tt.expected, actual)
		}
	}
}





// func BenchmarkIsPangram(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		handle, err := os.Open("testFileAlphabetAtBeginning.txt")
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		defer handle.Close()
// 		IsPangram(handle)
// 	}
// }

// func BenchmarkIsPangramReadAll(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		handle, err := os.Open("testFileAlphabetAtBeginning.txt")
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		defer handle.Close()

// 		IsPangramReadAll(handle)
// 	}
// }
