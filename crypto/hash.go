package crypto

import (
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ripemd160"
)

func Blake256(data []byte) ([]byte, error) {
	b, _ := blake2b.New(32, nil)
	if _, err := b.Write(data); err != nil {
		return nil, err
	}
	bs := b.Sum(nil)
	return bs, nil
}

func Blake256WithKey(data []byte, key []byte) ([]byte, error) {
	b, _ := blake2b.New(32, key)
	if _, err := b.Write(data); err != nil {
		return nil, err
	}
	bs := b.Sum(nil)
	return bs, nil
}

func Blake128(data []byte) ([]byte, error) {
	b, _ := blake2b.New(16, nil)
	if _, err := b.Write(data); err != nil {
		return nil, err
	}
	bs := b.Sum(nil)
	return bs, nil
}

func Blake128WithKey(data []byte, key []byte) ([]byte, error) {
	b, _ := blake2b.New(16, key)
	if _, err := b.Write(data); err != nil {
		return nil, err
	}
	bs := b.Sum(nil)
	return bs, nil
}

func Blake160(data []byte) ([]byte, error) {
	b, _ := blake2b.New(20, nil)
	if _, err := b.Write(data); err != nil {
		return nil, err
	}
	bs := b.Sum(nil)
	return bs, nil
}

func Blake160WithKey(data []byte, key []byte) ([]byte, error) {
	b, _ := blake2b.New(20, key)
	if _, err := b.Write(data); err != nil {
		return nil, err
	}
	bs := b.Sum(nil)
	return bs, nil
}

func Ripemd160(data []byte) ([]byte, error) {
	r := ripemd160.New()
	if _, err := r.Write(data); err != nil {
		return nil, err
	}
	rs := r.Sum(nil)
	return rs, nil
}
