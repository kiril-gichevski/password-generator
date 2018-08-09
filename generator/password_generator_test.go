package generator

import (
	"strconv"
	"strings"
	"testing"
)

func TestShuffle(t *testing.T) {
	list := []string{"1", "2", "3", "4"}
	tmp := make([]string, len(list))
	copy(tmp, list)

	// check length
	shuffle(list)
	if len(list) != 4 {
		t.Errorf("Shuffle changed list size")
	}

	// check if all elements are in the shuffled list
	for _, element := range list {
		if contains(list, element) != true {
			t.Errorf("Element not in shuffled array")
		}
	}

	// check the order of the elements
	if strings.Join(list, "") == strings.Join(tmp, "") {
		t.Errorf("List has not been shuffled")
	}
}

func TestShuffleAndConvertToString(t *testing.T) {
	list := []string{"1", "2", "3", "4"}
	shuffledString, _ := shuffleAndConvertToString(list)
	if len(shuffledString) != 4 {
		t.Errorf("Shuffle string is not complete")
	}
}

func TestIsVowel(t *testing.T) {
	if isVowel("a") != true {
		t.Errorf("Vowel detecting is wrong")
	}
	if isVowel("E") != true {
		t.Errorf("Vowel detecting is wrong")
	}

	if isVowel("K") != false {
		t.Errorf("Vowel detecting is wrong")
	}
}

func TestConvertVowelToNumber(t *testing.T) {
	if convertVowelToNumber(1, 1, "A") != true {
		t.Errorf("Converting Vowel to number condition is wrong")
	}
	if convertVowelToNumber(1, 1, "K") == true {
		t.Errorf("Converting Vowel to number condition is wrong")
	}
}

func TestGeneratePassword(t *testing.T) {
	//check password length
	password, _ := GeneratePassword(8, 2, 2)
	if len(password) < 8 {
		t.Errorf("The password doesn't contain the specific length")
	}

	// check if there are the specified number of integers
	numbers := 0
	for _, char := range password {
		if isNumber(string(char)) {
			numbers++
		}
	}
	if numbers < 2 {
		t.Errorf("The password doesn't contain the specific number of integers")
	}

	// check if there are the specified number of symbols
	numSymbols := 0
	splitSymbols := strings.Split(symbols, "")
	for _, char := range symbols {
		if contains(splitSymbols, string(char)) {
			numSymbols++
		}
	}
	if numSymbols < 2 {
		t.Errorf("The password doesn't contain the specific number of symbols")
	}
}

func contains(arr []string, element string) bool {
	for _, a := range arr {
		if a == element {
			return true
		}
	}
	return false
}

func isNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}
	return false
}
