package iteration

func Repeat(character string, numOfIterations int) string {
	var finalString string
	for i := 0; i < numOfIterations; i++ {
		finalString += character
	}
	return finalString
}
