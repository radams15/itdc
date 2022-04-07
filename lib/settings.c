//
// Created by rhys on 07/04/2022.
//

#include "include/Settings.h"

char* pt_fs_read(void* ptr, char* to_get);

struct SettingsData* pt_get_settings(void* client){
    return (struct SettingsData*) pt_fs_read(client, "/settings.dat");
}