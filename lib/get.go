package main

/*
typedef struct Motion{
	int x;
	int y;
	int z;
} Motion_t;
*/
import "C"

import (
	"unsafe"

	"go.arsenm.dev/itd/api"
)

//export pt_get_vers
func pt_get_vers(ptr uintptr)(*C.char){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	vers, _ := client.Version()

	return C.CString(vers);
}

//export pt_get_mac_addr
func pt_get_mac_addr(ptr uintptr) *C.char {
	client := *(*api.Client)(unsafe.Pointer(ptr))

	address, err := client.Address()
	if err != nil {
		return C.CString("")
	}

	return C.CString(address)
}

//export pt_get_battery_level
func pt_get_battery_level(ptr uintptr) int {
	client := *(*api.Client)(unsafe.Pointer(ptr))

	battLevel, err := client.BatteryLevel()
	if err != nil {
		return -1
	}

	return (int)(battLevel)
}

//export pt_get_heart_rate
func pt_get_heart_rate(ptr uintptr) int {
	client := *(*api.Client)(unsafe.Pointer(ptr))

	heartRate, err := client.HeartRate()

	if err != nil {
		return -1;
	}

	return (int)(heartRate)
}

//export pt_get_motion
func pt_get_motion(ptr uintptr) *C.Motion_t {
	client := *(*api.Client)(unsafe.Pointer(ptr))

	motionVals, err := client.Motion()
	if err != nil {
		return nil
	}

	p := (*C.Motion_t)(C.malloc(C.size_t(C.sizeof_Motion_t)))

	p.x = C.int(motionVals.X);
	p.y = C.int(motionVals.Y);
	p.z = C.int(motionVals.Z);


	return p
}

//export pt_get_step_count
func pt_get_step_count(ptr uintptr) int {
	client := *(*api.Client)(unsafe.Pointer(ptr))

	stepCount, err := client.StepCount()
	if err != nil {
		return -1;
	}

	return (int)(stepCount)
}