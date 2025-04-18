package util

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// --- Configuration parameters for Argon2id ---
const (
	// These parameters can be tuned according to your environment and security needs.
	timeCost   uint32 = 3         // Number of iterations
	memoryCost uint32 = 64 * 1024 // Memory usage in KiB (e.g. 64 MB)
	threads    uint8  = 4         // Number of threads (parallelism)
	keyLength  uint32 = 32        // Length of the generated key
	saltLength        = 16        // Salt length in bytes
)

// generateSalt creates a random salt of the given length.
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

// HashPassword hashes a plaintext password using Argon2id and returns an encoded hash.
// The encoded hash includes the algorithm, version, parameters, salt, and hash so that it can be verified later.
func HashPassword(password string) (string, error) {
	// Generate a random salt
	salt, err := generateSalt(saltLength)
	if err != nil {
		return "", err
	}

	// Generate the Argon2id hash of the password
	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, threads, keyLength)

	// Base64-encode the salt and hash for easy storage
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Format: $argon2id$v=19$m=65536,t=1,p=4$<salt>$<hash>
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memoryCost, timeCost, threads, b64Salt, b64Hash)

	return encodedHash, nil
}

// ComparePasswordAndHash verifies whether the provided password matches the encoded Argon2id hash.
// It extracts the parameters and salt from the encoded hash, recalculates the hash for the provided password,
// and uses a constant-time comparison to check for equality.
func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
	// The encoded hash should be in the format: $argon2id$v=19$m=65536,t=1,p=4$<salt>$<hash>
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	// Parse version information (e.g., "v=19")
	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, err
	}
	if version != argon2.Version {
		return false, errors.New("incompatible argon2 version")
	}

	// Parse parameters (e.g., "m=65536,t=1,p=4")
	var memory uint32
	var time uint32
	var threads uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads); err != nil {
		return false, err
	}

	// Decode the salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	// Decode the stored hash
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	// Re-compute the hash with the same parameters and salt using the provided password
	computedHash := argon2.IDKey([]byte(password), salt, time, memory, threads, uint32(len(hash)))

	// Compare the computed hash with the stored hash in constant time
	if subtle.ConstantTimeCompare(hash, computedHash) == 1 {
		return true, nil
	}
	return false, nil
}
