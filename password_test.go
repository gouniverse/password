package password

import (
	"testing"
)

// Hash Creates a password hash
func TestHash(t *testing.T) {
	hash, err := Hash("TESTtest")
    if err != nil {
        t.Error(err.Error())
	}
	if len(hash)!=60 {
         t.Error(len(hash))
	}
}

// Verify verifies that a password matches a hash
func TestVerify(t *testing.T) {
	hash, err := Hash("TEST")
    if err != nil {
		t.Error(err.Error())
	}
	if Verify("TEST",hash)==false {
         t.Error("Verification failed")
	}
}
