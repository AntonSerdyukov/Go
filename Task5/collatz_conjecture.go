package collatzconjecture

import "errors"

func CollatzConjecture(num int) (int, error) {
	if num <= 0 {
		return 0, errors.New("Wrong input")
	}
	numOfSteps := 0
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		numOfSteps += 1
	}
	return numOfSteps, nil
}
