package isbn

func IsValidISBN(isbnExample string) bool {
	var numArr []int

	for i := 0; i < len(isbnExample); i++ {
		switch isbnExample[i] {
		case '-':
			continue
		case 'X':
			if i == len(isbnExample)-1 {
				numArr = append(numArr, 10)
			} else {
				return false
			}
		case '0':
			numArr = append(numArr, 0)
		case '1':
			numArr = append(numArr, 1)
		case '2':
			numArr = append(numArr, 2)
		case '3':
			numArr = append(numArr, 3)
		case '4':
			numArr = append(numArr, 4)
		case '5':
			numArr = append(numArr, 5)
		case '6':
			numArr = append(numArr, 6)
		case '7':
			numArr = append(numArr, 7)
		case '8':
			numArr = append(numArr, 8)
		case '9':
			numArr = append(numArr, 9)
		default:
			return false
		}
	}

	if len(numArr) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < len(numArr); i++ {
		sum += numArr[i] * (10 - i)
	}

	if sum%11 == 0 {
		return true
	} else {
		return false
	}

}
