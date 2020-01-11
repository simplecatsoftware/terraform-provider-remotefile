package helpers

import (
	"terraform-provider-remotefiles/remotefiles/fetch"
	"testing"
)

func TestHashSha256(t *testing.T) {
	knownFile := fetch.LocalFile{Path: "hash_known.txt"}
	knownHash := "df2e8c587f36dee996ace336501547402d7ab6529dacb1e9e8003d5a1974efc8"
	computedHash, err := HashSha256(knownFile)

	if err != nil {
		t.Fatal("Could not generate computed hash", err)
	}

	if knownHash != computedHash {
		t.Fatal("Computed hash does not match known hash", knownHash, "!=", computedHash)
	}
}
