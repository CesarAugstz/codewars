package roman_numerals_decoder

var romansMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1_000,
}

func Decode(roman string) (res int) {

	decrement := 0
	increment := 0

	if roman == "" {
		return 0
	}

	for i, v := range roman {
		var next byte

		if i < (len(roman) - 1) {
			next = roman[i+1]
		}

		nextValue, exists := romansMap[next]
		value := romansMap[byte(v)]

		if !exists {
			nextValue = 0
		}

		if nextValue > value {

			decrement += value
		} else {

			increment += value
		}

	}

	return increment - decrement

}
