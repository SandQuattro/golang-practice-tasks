package bruteforce

import (
	"crypto/md5"
)

var alphabet = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func RecoverPassword(h []byte) string {
	maxLen := 8 // max pwd len
	base := len(alphabet)

	for length := 1; length <= maxLen; length++ {
		maxNum := pow(base, length)

		for num := 0; num < maxNum; num++ {
			candidate := numToString(num, length)
			hashedCandidate := md5.Sum([]byte(candidate))

			if string(hashedCandidate[:]) == string(h) {
				return candidate
			}
		}
	}

	return ""
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func numToString(num, length int) string {
	result := make([]rune, length)
	for i := length - 1; i >= 0; i-- {
		result[i] = alphabet[num%len(alphabet)]
		num /= len(alphabet)
	}
	return string(result)
}

// nothing change here !
func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
