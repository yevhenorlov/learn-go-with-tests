package iteration

func Repeat(character string, repeats int) (repeated string) {
	for i := 0; i < repeats; i++ {
		repeated += character
	}
	return
}
