package password_utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Utils struct {
	password *config.Password
}

func New(cfg *config.Password) *Utils {
	return &Utils{
		password: cfg,
	}
}

func (u *Utils) GeneratePassword(password string) (string, error) {
	salt := make([]byte, u.password.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, u.password.Iterations, u.password.Memory, u.password.Parallelism, u.password.KeyLength)

	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		u.password.Memory,
		u.password.Iterations,
		u.password.Parallelism,
		b64salt,
		b64hash,
	)

	return encodedHash, nil
}

func (u *Utils) ComparePassword(password, hash string) (bool, error) {
	vals := strings.Split(hash, "$")
	if len(vals) != 6 {
		return false, fmt.Errorf("invalid hash format")
	}

	var version int
	if _, err := fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return false, fmt.Errorf("invalid hash format")
	}
	if version != argon2.Version {
		return false, fmt.Errorf("invalid hash format")
	}

	p := &config.Password{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism); err != nil {
		return false, fmt.Errorf("invalid hash format")
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return false, fmt.Errorf("invalid hash format")
	}
	p.SaltLength = uint32(len(salt))

	decodedHash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return false, fmt.Errorf("invalid hash format")
	}
	p.KeyLength = uint32(len(decodedHash))

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	return subtle.ConstantTimeCompare(decodedHash, otherHash) == 1, nil
}
