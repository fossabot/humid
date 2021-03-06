package wordlist_test

import (
	"testing"

	"github.com/kscarlett/humid/wordlist"
)

func TestAnimalsDuplicate(t *testing.T) {
	err, problems := duplicate(wordlist.Animals)

	if err != nil {
		t.Errorf("wordlist contains duplicates: %v", problems)
	}
}

func TestAnimalsEmpty(t *testing.T) {
	err := empty(wordlist.Animals)

	if err != nil {
		t.Error("wordlist is empty")
	}
}

func TestAnimalsIllegalCharacters(t *testing.T) {
	err, problems := illegalCharacters(wordlist.Animals)

	if err != nil {
		t.Errorf("wordlist contains illegal characters: %v", problems)
	}
}

func TestAnimalsLowercase(t *testing.T) {
	err, problems := lowercase(wordlist.Animals)

	if err != nil {
		t.Errorf("wordlist contains non-lowercase letters: %v", problems)
	}
}

func TestAnimalsComplexDuplicate(t *testing.T) {
	err, problems := duplicate(wordlist.AnimalsComplex)

	if err != nil {
		t.Errorf("wordlist contains duplicates: %v", problems)
	}
}

func TestAnimalsComplexEmpty(t *testing.T) {
	err := empty(wordlist.AnimalsComplex)

	if err != nil {
		t.Error("wordlist is empty")
	}
}

func TestAnimalsComplexIllegalCharacters(t *testing.T) {
	err, problems := illegalCharacters(wordlist.AnimalsComplex)

	if err != nil {
		t.Errorf("wordlist contains illegal characters: %v", problems)
	}
}

func TestAnimalsComplexLowercase(t *testing.T) {
	err, problems := lowercase(wordlist.AnimalsComplex)

	if err != nil {
		t.Errorf("wordlist contains non-lowercase letters: %v", problems)
	}
}
