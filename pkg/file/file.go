package file

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd   int    //文件描述符
	name string //文件名
}

type matrix struct {
	test int
}

func NewFile(fd int, name string) (file *File) {
	if fd < 0 {
		file = nil
		return
	}

	file = &File{fd, name}
	sizeof := unsafe.Sizeof(file)
	fmt.Println("The File struct use ", sizeof, " byte")
	return
}

func NewMatrix() *matrix {
	m := new(matrix)
	m.test = 0
	return m
}
