package hamming

import "errors"

func Distance(firstDNA string, secondDNA string) (int, error) {
	if len(firstDNA) != len(secondDNA) {
		return 0, errors.New("Error")
	} else {
		hamDistance := 0
		for i := 0; i < len(firstDNA); i++ {
			if firstDNA[i] != secondDNA[i] {
				hamDistance++
			}
		}
		return hamDistance, nil
	}
}
