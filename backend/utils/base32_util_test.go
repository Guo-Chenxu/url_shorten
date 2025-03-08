package utils

import (
	"testing"
)

func TestBase32(t *testing.T) {
	input := uint64(79695888)
	encoded := EncodeBase62Uint64(input)
	t.Log(encoded)
	decoded, err := DecodeBase62Uint64(encoded)
	if err != nil {
		t.Error(err)
	}
	t.Log(decoded)
	if input != decoded {
		t.Errorf("want %d, got %d", input, decoded)
	}
}
