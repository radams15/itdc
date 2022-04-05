package main

import "C"

import (
	"unsafe"

	"time"

	"go.arsenm.dev/itd/api"
)

//export pt_set_time
func pt_set_time(ptr uintptr){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	client.SetTimeNow()
}

//export pt_set_time_cust
func pt_set_time_cust(ptr uintptr, secs int64){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	time_obj := time.Unix(secs, 0);

	client.SetTime(time_obj)
}

//export pt_set_weather
func pt_set_weather(ptr uintptr){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	client.UpdateWeather();
}