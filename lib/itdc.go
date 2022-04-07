package main

/*
#ifndef __cplusplus
#define PT GoUintptr
#endif
*/
import "C"

import (
	"unsafe"

	"go.arsenm.dev/itd/api"
)

const DefaultAddr = "/tmp/itd/socket"

func newClientWithAddrGo(addr string) (uintptr){
	client, err := api.New(addr)

	if err != nil {
		return 0
	}

	return uintptr(unsafe.Pointer(client))
}

//export pt_new_with_addr
func pt_new_with_addr(addr *C.char) (uintptr){
	return newClientWithAddrGo(C.GoString(addr))
}

//export pt_new
func pt_new() (uintptr){
	return newClientWithAddrGo(DefaultAddr)
}

//export pt_free
func pt_free(ptr uintptr){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	client.Close()
}






//export pt_notify
func pt_notify(ptr uintptr, title, subtitle *C.char){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	client.Notify(C.GoString(title), C.GoString(subtitle))
}





func main(){}
