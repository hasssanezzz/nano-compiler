package lexer

func contains(word string, words []string) int {
	for i, currWord := range words {
		if word == currWord {
			return i
		}
	}
	return -1
}
