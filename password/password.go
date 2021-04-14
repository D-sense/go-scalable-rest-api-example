package password

import (
	"crypto/rand"
	"crypto/subtle"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

// Password constants
const (
	SaltLen        = 32
	HashLen        = 64
)


type Hash struct {
	Hash []byte `json:"hash"`
	Salt []byte `json:"salt"`
}

// Value get value of Jsonb
func (h Hash) Value() (driver.Value, error) {
	j, err := json.Marshal(h)
	return j, err
}

// Scan scan value into Hash
func (h *Hash) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, h)
}

func NewHashedPassword(password string) (*Hash, error) {
	salt := salt()
	hash, err := createPasswordHash(password, salt)
	if err != nil {
		return nil, err
	}

	return &Hash{Hash: hash, Salt: salt}, nil
}

func (self *Hash) IsEqualTo(password string) bool {
	return VerifyPassword(password, self.Hash, self.Salt)
}

// Generate a random salt of suitable length
func salt() []byte {
	salt := make([]byte, SaltLen)

	_, _ = rand.Read(salt)
	return salt
}

// Create a hash of a password and salt
func createPasswordHash(password string, salt []byte) ([]byte, error) {
	password = strings.TrimSpace(password)

	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, HashLen)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// VerifyPassword checks that a password matches a stored hash and salt
func VerifyPassword(password string, hash []byte, salt []byte) bool {
	hs, err := createPasswordHash(password, salt)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(hs, hash) == 1
}
