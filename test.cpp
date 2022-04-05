#include <cstdio>

#include <PT.h>

int main(){
    GoUintptr client = pt_new();

    Directory_t* dir = pt_fs_ls(client, "/");

    for(int i=0 ; i<dir->length ; i++){
        printf("%s => %s (%lu b)\n", (dir->files[i]->ntype == FILE_TYPE? "F": "D"), dir->files[i]->name, dir->files[i]->size);
    }
}
