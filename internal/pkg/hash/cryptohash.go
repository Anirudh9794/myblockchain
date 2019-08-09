package hash

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
	"encoding/hex"
)

func CreateHash(vals ...string) string {
	sort.Strings(vals)

	v := strings.Join(vals, " ")

	hasher := sha256.Sum256([]byte(v))

	return fmt.Sprintf("%x",hasher)
}

func HexToBinary(h string) string {
	bytearr, err := hex.DecodeString(h)
	if err != nil {
		panic(err.Error())
	}

	return bytesToBinary(bytearr)
}

func bytesToBinary(buffer []byte) string {
	str := ""

	for _, b := range buffer {
		str += fmt.Sprintf("%08b", b)
	}

	return str
}
