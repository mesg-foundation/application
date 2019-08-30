package hash

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash"

	"github.com/cnf/structhash"
	"github.com/mr-tron/base58"
)

// DefaultHash is a default hashing algorithm - "sha256".
var DefaultHash = sha256.New

// size is a default size for hashing algorithm.
var size = DefaultHash().Size()

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
	return d.Write(structhash.Dump(v, 0))
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
		return nil, fmt.Errorf("hash: invalid length string")
	}
	return Hash(hash), nil
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

// Marshal marshals hash into slice of bytes. It's used by protobuf.
func (h Hash) Marshal() ([]byte, error) {
	return h, nil
}

// MarshalTo marshals hash into slice of bytes. It's used by protobuf.
func (h Hash) MarshalTo(data []byte) (int, error) {
	return copy(data, h), nil
}

// Unmarshal unmarshals slice of bytes into hash. It's used by protobuf.
func (h *Hash) Unmarshal(data []byte) error {
	*h = data
	return nil
}

// Size retruns size of hash. It's used by protobuf.
func (h Hash) Size() int {
	if len(h) == 0 {
		return 0
	}
	return size
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
