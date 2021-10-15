package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(input []byte) string {
	md5hash := md5.New()
	md5hash.Write(input)
	return hex.EncodeToString(md5hash.Sum(nil))
}

func MD5Hash_String(input string) string {
	return MD5Hash([]byte(input))
}

//combine all string elements and hash
func MD5Hash_StringArray(input []string) string {
	var merge string
	for i := 0; i < len(input); i++ {
		merge = merge + input[i]
	}
	return MD5Hash_String(merge)
}
