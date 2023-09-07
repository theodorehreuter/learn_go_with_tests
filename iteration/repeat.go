package iteration

func Repeat(char string, rep int) string {
	var repeated string
	for i := 0; i < rep; i++ {
		repeated += char
	}
	return repeated
}
