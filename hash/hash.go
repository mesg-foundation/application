package hash

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"hash"

	"github.com/mesg-foundation/engine/hash/structhash"
	"github.com/mr-tron/base58"
)

// DefaultHash is a default hashing algorithm - "sha256".
var DefaultHash = sha256.New

// size is a default size for hashing algorithm.
var size = DefaultHash().Size()

var errInvalidLen = errors.New("hash: invalid length")

// Digest represents the partial evaluation of a checksum.
type Digest struct {
	hash.Hash
}

// Sum appends the current checksum to b and returns the Hash.
func (d *Digest) Sum(b []byte) Hash {
	return Hash(d.Hash.Sum(b))
}

// WriteObject  adds an interface data to the running hash.
// It never retruns an error.
func (d *Digest) WriteObject(v interface{}) (int, error) {
	return d.Write(structhash.Dump(v))
}

// A Hash is a type for representing common hash.
type Hash []byte

// New returns new hash from a given integer.
func New() *Digest {
	return &Digest{
		Hash: DefaultHash(),
	}
}

// Dump takes an interface and returns its hash representation.
func Dump(v interface{}) Hash {
	d := New()
	d.WriteObject(v)
	return d.Sum(nil)
}

// Int returns a new hash from a given integer.
// NOTE: This method is for tests purpose only.
func Int(h int) Hash {
	hash := make(Hash, size)
	binary.PutUvarint(hash, uint64(h))
	return hash
}

// Random returns a new random hash.
func Random() (Hash, error) {
	hash := make(Hash, size)
	n, err := rand.Reader.Read(hash)
	if err != nil {
		return nil, fmt.Errorf("hash generate random error: %s", err)
	}
	if n != size {
		return nil, fmt.Errorf("hash generate random error: invalid hash length")
	}
	return hash, nil
}

// Decode decodes the base58 encoded hash. It returns error
// when a hash isn't base58,encoded or hash length is invalid.
func Decode(h string) (Hash, error) {
	hash, err := base58.Decode(h)
	if err != nil {
		return nil, fmt.Errorf("hash: %s", err)
	}
	if len(hash) != size {
		return nil, errInvalidLen
	}
	return Hash(hash), nil
}

// DecodeFromBytes decodes hash and checks it length.
// It returns empty hash on nil slice of bytes.
func DecodeFromBytes(data []byte) (Hash, error) {
	if len(data) != size {
		return nil, errInvalidLen
	}
	return Hash(data), nil
}

// IsZero reports whethere h represents empty hash.
func (h Hash) IsZero() bool {
	return len(h) == 0
}

// String returns the hash hex representation.
func (h Hash) String() string {
	return base58.Encode(h)
}

// Equal returns a boolean reporting whether h and h1 are the same hashes.
func (h Hash) Equal(h1 Hash) bool {
	return bytes.Equal(h, h1)
}

// Valid checks if service hash length is valid.
// It treats empty hash as valid one.
func (h Hash) Valid() bool {
	return len(h) == 0 || len(h) == size
}

// MarshalJSON mashals hash into encoded json string.
func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(h))
}

// UnmarshalJSON unmashals hex encoded json string into hash.
func (h *Hash) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if str == "" {
		return nil
	}
	h1, err := base58.Decode(str)
	if err != nil {
		return err
	}
	*h = Hash(h1)
	return nil
}
