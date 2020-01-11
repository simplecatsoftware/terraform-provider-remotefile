package helpers

import (
	"terraform-provider-remotefiles/remotefiles/data"
	"testing"
)

func TestHashSha256(t *testing.T) {
	knownFile := data.LocalFile{Path: "../../test/fixtures/known_file_hash"}
	knownHash := "df2e8c587f36dee996ace336501547402d7ab6529dacb1e9e8003d5a1974efc8"
	computedHash, err := hashSha256(knownFile)

	if err != nil {
		t.Fatal("Could not generate computed hash", err)
	}

	if knownHash != computedHash {
		t.Fatal("Computed hash does not match known hash", knownHash, "!=", computedHash)
	}
}