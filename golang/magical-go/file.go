package main

import (
	"fmt"
)

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ICloser interface {
	Close() error
}

type File struct {
	filename string
	fid int64
}

func (f *File) Read(buf []byte) (n int, err error) {
	fmt.Println("File opened")
	return 0, nil
}

func (f *File) Write(buf []byte) (n int, err error) {
	fmt.Println("File writed")
	return 0, nil
}

func (f *File) Close() error {
	fmt.Println("File closed")
	return nil
}

func main() {
	bytes := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var file1 IFile = new(File)
	file1.Read(bytes)
	file1.Write(bytes)
	file1.Close()

	var file2 IReader = new(File)
	file2.Read(bytes)

	var file3 IWriter = new(File)
	file3.Write(bytes)

	var file4 ICloser = new(File)
	file4.Close()
}