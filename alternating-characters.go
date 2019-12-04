package main

func alternatingCharacters(s string) int32 {
	var deletions int32
	sLength := int32(len(s))

	for i := int32(1); i < sLength; i++ {
		if s[i-1] == s[i] {
			deletions++
		}
	}

	return deletions
}

func main() {}
