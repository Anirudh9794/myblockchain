package hash

import (
	"strings"
	"testing"
)

func TestCreateHash(t *testing.T) {

	t.Run("Order doesn't matter", func(t *testing.T) {
		h1 := CreateHash("One", "Two", "Three")
		h2 := CreateHash("Two", "One", "Three")

		if h1 != h2 {
			t.Fail()
		}
	})

	t.Run("Correct hash is generated", func(t *testing.T) {
		sampleHash := "50D858E0985ECC7F60418AAF0CC5AB587F42C2570A884095A9E8CCACD0F6545C"
		sampleHash = strings.ToLower(sampleHash)

		createdHash := CreateHash("example")

		if createdHash != sampleHash {
			t.Errorf("created hash: %s, expected hash: %s", createdHash, sampleHash)
		}
	})
}
