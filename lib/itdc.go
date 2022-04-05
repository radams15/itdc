package main

/*
#include <stdint.h>

enum ClockType { H24, H12 };
enum Notification { ON, OFF };
enum ChimesOption { None, Hours, HalfHours };

enum WakeUpMode {
    SingleTap = 0,
    DoubleTap = 1,
    RaiseWrist = 2,
    Shake = 3,
};

enum Colors {
    White,
    Silver,
    Gray,
    Black,
    Red,
    Maroon,
    Yellow,
    Olive,
    Lime,
    Green,
    Cyan,
    Teal,
    Blue,
    Navy,
    Magenta,
    Purple,
    Orange
};

struct PineTimeStyle {
    enum Colors ColorTime;
    enum Colors ColorBar;
    enum Colors ColorBG;
};

enum BrightnessLevels { Off, Low, Medium, High };

struct SettingsData {
        uint32_t version;
        uint32_t stepsGoal;
        uint32_t screenTimeOut;

        enum ClockType clockType;
        enum Notification notificationStatus;

        uint8_t clockFace;
        enum ChimesOption chimesOption;

        struct PineTimeStyle PTS;

        unsigned char wakeUpMode : 4;
        uint16_t shakeWakeThreshold;
        enum BrightnessLevels brightLevel;
};

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
