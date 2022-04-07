#include <cstdio>
#include <string.h>

#include <PT.h>
#include <iostream>

int main(){
    PT client;

    char* data = "hello world!";

    pt_fs_write((GoUintptr) client, data, strlen(data), "hi.dat");

    /*auto bytes = pt_fs_read((GoUintptr) client, (char*) "/settings.dat");

    for(int i=0 ; i<32 ; i++){
        printf("%d ", bytes[i]);
    }
    printf("\n");

    auto* dat = (SettingsData*) bytes;
    printf("Timeout: %d\n", dat->screenTimeOut);*/
}
