package DDATM_ES_0002

import (
	"errors"
)

func Exec(numbers []int) ([]int, error) {
	obtainMax := 0
	var fillNumbers []int

	for _, number := range numbers {
		if number < 0 {
			return nil, errors.New("invalid number, you need input positive numbers")
		}
		if number > obtainMax {
			obtainMax = number
		}
	}

	for i := 1; i <= obtainMax; i++ {
		fillNumbers = append(fillNumbers, i)
	}

	return fillNumbers, nil
}

func ExecAlternative(numbers []int) ([]int, error) {
	obtainMax := 0
	obtainMin := 0
	var fillNumbers []int
	
	for _, number := range numbers {
		if number < 0 {
			return nil, errors.New("invalid number, you need input positive numbers")
		}
		if number > obtainMax {
			obtainMax = number
		}
		if obtainMin == 0 || number < obtainMin{
			obtainMin = number
		}
	}

	for i := obtainMin; i <= obtainMax; i++ {
		fillNumbers = append(fillNumbers, i)
	}

	return fillNumbers, nil
}