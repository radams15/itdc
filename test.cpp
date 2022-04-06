#include <cstdio>

#include <PT.h>
#include <iostream>

int main(){
    PT client;

    /*std::vector nodes = client.fs_ls((char*) "/");

    for(Node_t* node : nodes){
        printf("%s => %s (%lu b)\n", (node->ntype == FILE_TYPE? "F": "D"), node->name, node->size);

        if(node->ntype == FILE_TYPE) {
            const char* data = pt_fs_read((GoUintptr) client, node->name);

            std::cout << data << std::endl;
        }
    }*/

    auto bytes = pt_fs_read((GoUintptr) client, (char*) "/settings.dat");

    /*for(int i=0 ; i<32 ; i++){
        printf("%d ", raw[i]);
    }*/


    for(int i=0 ; i<32 ; i++){
        printf("%d ", bytes[i]);
    }
    printf("\n");

    auto* dat = (SettingsData*) bytes;
    printf("Timeout: %d\n", dat->screenTimeOut);
}
