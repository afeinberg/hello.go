package sayhello

import (
	// #include <stdlib.h>
	// #include "hello.h"
	"C"
	"unsafe"
	"errors"
)

type hello struct {
	hello *C.Hello_T
}

func NewHello(name string) (*hello, error) {
	h := new(hello)
	h.hello = new(C.Hello_T)
	if !C.Hello_create(C.CString(name), h.hello) {
		defer h.Dispose()
		return nil, errors.New("unable to create hello!")
	}
	return h, nil
}

func (h *hello) SayHello() (string, error) {
	b := make([]byte, 128)
	p := unsafe.Pointer(&b)
	s := (*C.char)(p)
	if !C.Hello_sayHello(h.hello, s, C.size_t(len(b))) {
		return "", errors.New("insufficient buffer space!")
	}
	return C.GoString(s), nil
}

func (h *hello) Dispose() {
	C.free(unsafe.Pointer(C.Hello_getName(h.hello)))
}
