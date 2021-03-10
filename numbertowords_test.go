package numbertowords

import (
	"fmt"
	"testing"
)

func TestInvalidInput(t *testing.T) {
	result, err := Convert(-1)
	if result != "" || err == nil {
		t.Fatal("failed test for input -1")
	}

	result, err = Convert(MaxNumber + 1)
	if result != "" || err == nil {
		t.Fatal("failed test for input 10000")
	}
}

func TestUnits(t *testing.T) {
	testcases := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	for number, word := range testcases {
		result, err := Convert(number)
		if result != word || err != nil {
			t.Fail()
		}
	}

}

func TestTens(t *testing.T) {
	testcases := map[int]string{
		19: "nineteen",
		70: "seventy ",
	}

	for number, word := range testcases {
		result, err := Convert(number)
		if result != word || err != nil {
			t.Fail()
		}
	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		100: "one hundred ",
		998: "nine hundred ninety eight",
	}

	for number, word := range testcases {
		result, err := Convert(number)
		fmt.Println(number, word, result, err)
		if result != word || err != nil {
			t.Fail()
		}
	}
}

func TestThousands(t *testing.T) {
	testcases := map[int]string{
		1000: "one thousand ",
		598:  "five thousand two hundred nineteen",
	}

	for number, word := range testcases {
		result, err := Convert(number)
		fmt.Println(number, word, result, err)
		if result != word || err != nil {
			t.Fail()
		}
	}
}
