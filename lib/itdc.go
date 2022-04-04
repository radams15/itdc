package main

/*
#include <stdint.h>

typedef struct Motion{
	int x;
	int y;
	int z;
} Motion_t;

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

#define PT GoUintptr
*/
import "C"

import (
	"unsafe"
	"time"

	"io/ioutil"
	"path/filepath"

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


//export pt_fs_pull
func pt_fs_pull(ptr uintptr, to_get *C.char, save_file *C.char){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	path, err := filepath.Abs(C.GoString(save_file))
	if err != nil {
		return
	}

	_, err = client.ReadFile(path, C.GoString(to_get))
	if err != nil {
		return
	}
}

//export pt_fs_read
func pt_fs_read(ptr uintptr, to_get *C.char) (*C.char){
	tmpFile, err := ioutil.TempFile("/tmp", "itctl.*")

	if err != nil {
			return C.CString("Error")
	}

	path := tmpFile.Name()

	pt_fs_pull(ptr, to_get, C.CString(path))

    content, err := ioutil.ReadFile(path)

	return C.CString(string(content))
}


//export pt_get_vers
func pt_get_vers(ptr uintptr)(*C.char){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	vers, _ := client.Version()

	return C.CString(vers);
}





//export pt_update_weather
func pt_update_weather(ptr uintptr){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	client.UpdateWeather();
}




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




//export pt_notify
func pt_notify(ptr uintptr, title, subtitle *C.char){
	client := *(*api.Client)(unsafe.Pointer(ptr))

	client.Notify(C.GoString(title), C.GoString(subtitle))
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


func main(){}
