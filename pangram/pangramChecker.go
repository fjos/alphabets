package pangram

import (
	"io"
	"log"
	"strings"
)

func IsLatinPangram(input io.Reader) (bool, error) {
	return IsPangram(input, alphabetMap["latin"], 64)
}

func IsPangram(input io.Reader, alphabet string, readSize int) (bool, error) {
	//initialize a map containing rune -> bool
	alphabetMap := make(map[rune]bool)
	for _, value := range alphabet {
		alphabetMap[value] = true
	}
	totalBytesRead := 0

	buffer := make([]byte, readSize)

	for {
		bytesRead, err := input.Read(buffer)
		totalBytesRead += bytesRead
		if err != nil {
			//exit if eof reached
			if err == io.EOF {
				RemoveRunesFromMap(string(buffer[:bytesRead]), alphabetMap)
				break
			}
			return false, err
		}

		RemoveRunesFromMap(string(buffer[:bytesRead]), alphabetMap)
		//early escape if allowable
		if len(alphabetMap) == 0 {
			log.Printf("alphabet found after %d bytes\n", totalBytesRead)

			return true, nil
		}
	}
	if len(alphabetMap) == 0 {
		log.Printf("alphabet found after %d bytes\n", totalBytesRead)

		return true, nil
	}
	return false, nil
}

//remove runes from map on a per chunk basis, delete does nothing when value doesn't exist
func RemoveRunesFromMap(chunk string, alphabetMap map[rune]bool) {
	lowerChunk := strings.ToLower(chunk)
	for _, value := range lowerChunk {
		delete(alphabetMap, value)
	}
}
