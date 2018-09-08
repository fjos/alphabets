package pangram

import (
	"log"
	"os"
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

func TestIsPangram(t *testing.T) {
	for _, tt := range pangramTests {
		handle := strings.NewReader(tt.input)
		actual := IsPangram(handle, "abcdefghijklmnopqrstuvwxyz", 64)
		if actual != tt.expected {
			t.Errorf("IsPangram(%s): expected %t, actual %t", tt.input, tt.expected, actual)
		}
	}
}

// func TestIsPangramReadAll(t *testing.T) {
// 	for _, tt := range pangramTests {
// 		handle := strings.NewReader(tt.input)
// 		actual := IsPangramReadAll(handle)
// 		if actual != tt.expected {
// 			t.Errorf("IsPangramReadAll(%s): expected %t, actual %t", tt.input, tt.expected, actual)
// 		}
// 	}
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
