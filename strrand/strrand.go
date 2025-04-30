package strrand

import (
	"math/rand/v2"
	"strconv"
	"time"
	"unsafe"
)

const (
	alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	otpSafe      = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var reg = map[rune][]rune{
	'0': {'4', 'c', 'l', 'u', 'F', 'U', 'Z'},
	'1': {'5', 'e', 's', 'B', 'E', 'V', 'Y'},
	'2': {'1', 'd', 'k', 'x', 'N', 'Q'},
	'3': {'7', 'b', 'q', 'D', 'H', 'X'},
	'4': {'6', 'h', 't', 'C', 'G', 'S'},
	'5': {'8', 'j', 'n', 'w', 'M', 'W'},
	'6': {'0', 'a', 'o', 'A', 'J', 'R'},
	'7': {'9', 'g', 'p', 'y', 'I', 'P'},
	'8': {'2', 'i', 'r', 'v', 'L', 'T'},
	'9': {'3', 'f', 'm', 'z', 'K', 'O'},
}

func RandomID() string {
	return randomTimeAlphaNumeric(alphaNumeric, 16)
}

func RandomOTP() string {
	return randomAlphaNumeric(otpSafe, 6)
}

func randomTimeAlphaNumeric(dictionary string, length int) string {
	s := randomAlphaNumeric(dictionary, length-10)
	t := []rune(strconv.FormatInt(time.Now().Unix(), 10))

	for i := 0; i < len(t); i++ {
		t[i] = reg[t[i]][rand.IntN(len(reg[t[i]]))]
	}

	return string(t) + s
}

func randomAlphaNumeric(dictionary string, length int) string {
	b := make([]byte, length)

	for i, cache, remain := length-1, rand.Uint64(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Uint64(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(dictionary) {
			b[i] = dictionary[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
