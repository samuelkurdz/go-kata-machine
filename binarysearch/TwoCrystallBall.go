package binarysearch

import (
	"fmt"
	"math"
	"math/rand"
)

func GenerateRandomNumber(minBound, maxBound int) int {
	return rand.Intn(maxBound-minBound) + minBound
}

func LinearSearch(haystack []bool) {
	for i, value := range haystack {
		if value {
			fmt.Println(i)
			break
		}
	}
}

func TwoCrystalBall(haystack []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(haystack)))))

	i := 0
	for ; i < len(haystack); i += jumpAmount {
		if haystack[i] {
			break
		}
	}

	//if i == len(haystack) {
	//	return -1
	//}

	for j := i - jumpAmount; j < i && j < len(haystack); j++ {
		if haystack[j] {
			return j
		}
	}

	return -1
}
