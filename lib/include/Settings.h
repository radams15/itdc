//
// Created by rhys on 07/04/2022.
//

#ifndef ITDC_SETTINGS_H
#define ITDC_SETTINGS_H

#include <stdint.h>

#ifdef __cplusplus
#include <bitset>
#endif

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
#ifdef __cplusplus
    std::bitset<4> wakeUpMode;
#else
    long wakeUpMode; // Emulate a std::bitset<4> in C
#endif
    uint16_t shakeWakeThreshold;
    enum BrightnessLevels brightLevel;
};

#ifdef __cplusplus
extern "C" {
#endif

struct SettingsData *pt_get_settings(void *client);

#ifdef __cplusplus
}
#endif


#endif //ITDC_SETTINGS_H
