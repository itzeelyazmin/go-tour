package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func(m MyReader) Read(b []byte) (i int, err error){
	for a := range b{
	b[a]='A'
	}
	return len (b), nil
}
func main() {
	reader.Validate(MyReader{})
}

