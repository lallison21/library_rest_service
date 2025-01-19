package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/lallison21/library_rest_service/internal/config/config"
	"golang.org/x/crypto/argon2"
)

type Utils struct {
	password *config.Password
}

func NewPassword(cfg *config.Password) *Utils {
	return &Utils{
		password: cfg,
	}
}

func (u *Utils) GeneratePassword(password string) (string, error) {
	salt := make([]byte, u.password.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, u.password.Iterations, u.password.Memory, u.password.Parallelism,
		u.password.KeyLength)

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
	const hashPartsCount = 6

	vals := strings.Split(hash, "$")
	if len(vals) != hashPartsCount {
		//nolint: err113 // can't wrap error
		err := errors.New("invalid hash parts")

		return false, err
	}

	var version int
	if _, err := fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return false, fmt.Errorf("scan hash: %w", err)
	}

	if version != argon2.Version {
		//nolint: err113 // can't wrap error
		err := errors.New("invalid argon version")

		return false, err
	}

	pass := &config.Password{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &pass.Memory, &pass.Iterations, &pass.Parallelism); err != nil {
		return false, fmt.Errorf("invalid scan hash: %w", err)
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return false, fmt.Errorf("invalid decode salt: %w", err)
	}

	//nolint:gosec // salt len not overflow uint32
	pass.SaltLength = uint32(len(salt))

	decodedHash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return false, fmt.Errorf("invalid decode hash: %w", err)
	}

	//nolint:gosec // key len not overflow uint32
	pass.KeyLength = uint32(len(decodedHash))

	otherHash := argon2.IDKey([]byte(password), salt, pass.Iterations, pass.Memory, pass.Parallelism, pass.KeyLength)

	return subtle.ConstantTimeCompare(decodedHash, otherHash) == 1, nil
}
