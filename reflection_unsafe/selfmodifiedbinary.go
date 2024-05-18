package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

type persistent struct {
	Magic   [8]byte
	Content int64
}

var p = persistent{
	Magic:   [8]byte{0xBA, 0xDD, 0xFA, 0xCE, 0xBE, 0xEF, 0xCA, 0xCE},
	Content: 0,
}

func main() {
	fmt.Println(p.Content)

	const size = int(unsafe.Sizeof(p))
	currentBuf := bytes.NewBuffer(make([]byte, 0, size))
	err := binary.Write(currentBuf, binary.LittleEndian, p)
	if err != nil {
		panic(err)
	}

	newP := p
	newP.Content++

	newBuf := bytes.NewBuffer(make([]byte, 0, size))
	err = binary.Write(newBuf, binary.LittleEndian, newP)

	currentBytes, newBytes := currentBuf.Bytes(), newBuf.Bytes()
	self, err := os.OpenFile(os.Args[0], os.O_RDONLY, 0755)
	selfBytes, err := ioutil.ReadAll(self)

	i := bytes.Index(selfBytes, currentBytes)
	copy(selfBytes[i:i+size], newBytes)
	newSelf, err := ioutil.TempFile("", "selfmodifying")
	_, err = newSelf.Write(selfBytes)
	err = os.Rename(newSelf.Name(), self.Name())
	err = os.Chmod(self.Name(), 0755)
}
