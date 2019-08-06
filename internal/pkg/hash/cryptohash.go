package hash

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

func CreateHash(vals ...string) string {
	sort.Strings(vals)

	v := strings.Join(vals, " ")

	hasher := sha256.Sum256([]byte(v))

	hash := fmt.Sprintf("%x", hasher)

	return hash
}
