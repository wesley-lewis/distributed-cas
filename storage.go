package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := (i * blocksize), (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Original: hashStr,
	}

}

type PathTransformFunc func(string) PathKey

type PathKey struct {
	Pathname string
	Original string
}

func (p PathKey) Filename() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Original)
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(storeOpts StoreOpts) *Store {
	return &Store{
		StoreOpts: storeOpts,
	}
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathKey.Pathname, os.ModePerm); err != nil {
		return err
	}

	pathAndFileName := pathKey.Filename()

	f, err := os.Create(pathAndFileName)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s\n", n, pathAndFileName)

	return nil
}
