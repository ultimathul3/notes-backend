package hash

import "crypto/sha256"

type sha256Hasher struct {
	salt []byte
}

func NewSHA256Hasher(salt []byte) *sha256Hasher {
	return &sha256Hasher{
		salt: salt,
	}
}

func (h *sha256Hasher) Hash(data []byte) ([]byte, error) {
	data = append(data, h.salt...)
	hash := sha256.Sum256(data)
	return hash[:], nil
}
