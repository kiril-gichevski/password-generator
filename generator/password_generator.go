package generator

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

func GeneratePassword(minLength int, numbers int, specialChars int) string {
	lettersList := make([]string, minLength)

	LettersLength := len(letters)
	for i := 0; i < minLength; i++ {
		Letter := string(letters[rand.Intn(LettersLength)])
		if convertVowelToNumber(numbers, Letter) {
			lettersList[i] = strconv.Itoa(rand.Intn(10))
			numbers--
		} else {
			lettersList[i] = Letter
		}
	}

	numbersList := make([]string, numbers)
	for i := 0; i < numbers; i++ {
		numbersList[i] = strconv.Itoa(rand.Intn(10))
	}

	symbolsList := make([]string, specialChars)
	symbolsLength := len(symbols)
	for i := 0; i < specialChars; i++ {
		symbolsList[i] = string(symbols[rand.Intn(symbolsLength)])
	}

	if len(numbersList) > 0 {
		LettersAndNumbersList := append(lettersList, numbersList...)
		return shuffleAndConvertToString(append(LettersAndNumbersList, symbolsList...))
	}
	return shuffleAndConvertToString(append(lettersList, symbolsList...))

}

func convertVowelToNumber(numbers int, letter string) bool {
	if rand.Intn(2) == 1 && numbers > 0 && isVowel(letter) {
		return true
	}
	return false
}

func isVowel(Letter string) bool {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	if vowels[Letter] {
		return true
	}
	return false
}

func shuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func shuffleAndConvertToString(passwordList []string) string {
	shuffle(passwordList)
	return strings.Join(passwordList, "")
}
