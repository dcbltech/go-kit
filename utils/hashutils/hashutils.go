package hashutils

import (
	"crypto/md5"
	"fmt"
)

func MD5Hex(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
