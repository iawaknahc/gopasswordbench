package passwordbench

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

func Plaintext(numBytes int) string {
	b := make([]byte, numBytes)
	_, err := rand.Reader.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func Bcrypt(plaintext string, cost int) (string, error) {
	bcryptStr, err := bcrypt.GenerateFromPassword([]byte(plaintext), cost)
	if err != nil {
		return "", err
	}
	return string(bcryptStr), nil
}

func PBKDF2_HMAC_SHA256(plaintext string, iterations int) (salt []byte, hash []byte) {
	// RFC8018 says the salt must be at least 8 octets.
	// https://www.rfc-editor.org/rfc/rfc8018#section-4.1
	saltSize := 16
	salt = make([]byte, saltSize)
	_, err := rand.Reader.Read(salt)
	if err != nil {
		panic(err)
	}

	keyLen := sha256.Size
	hash = pbkdf2.Key([]byte(plaintext), salt, iterations, keyLen, sha256.New)

	return
}

func Argon2id(plaintext string, time uint32, memoryInKiB uint32, threads uint8) (salt []byte, hash []byte) {
	// RFC9106 recommends a salt size of 128 bit = 16 bytes.
	// https://www.rfc-editor.org/rfc/rfc9106.html#name-parameter-choice
	saltSize := 16
	salt = make([]byte, saltSize)
	_, err := rand.Reader.Read(salt)
	if err != nil {
		panic(err)
	}

	// RFC9106 recommends a tag length of 256 bit = 32 bytes.
	keyLen := uint32(32)
	hash = argon2.IDKey([]byte(plaintext), salt, time, memoryInKiB, threads, keyLen)

	return
}

func Scrypt(plaintext string, N int, r int, p int) (salt []byte, hash []byte) {
	// RFC7914 does not say what size the salt should have.
	// Let's use a common salt size of 16 bytes.
	// https://www.rfc-editor.org/rfc/rfc7914.html
	saltSize := 16
	salt = make([]byte, saltSize)
	_, err := rand.Reader.Read(salt)
	if err != nil {
		panic(err)
	}

	// RFC7914 does not say what size the key length should have.
	// Let's follow Argon2 and use 32 bytes.
	keyLen := 32
	hash, err = scrypt.Key([]byte(plaintext), salt, N, r, p, keyLen)
	if err != nil {
		panic(err)
	}

	return
}
