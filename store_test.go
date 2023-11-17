package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "mombestpicture"
	pathname := CASPathTransformFunc(key)
	expectedPathName := "cf5d4/b01c4/d9438/c22c5/6c832/f83bd/3e8c6/304f9"
	fmt.Println(pathname)
	if pathname != expectedPathName {
		t.Errorf("have %s want %s", pathname, expectedPathName)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpeg"))
	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
