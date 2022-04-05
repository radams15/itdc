#include <cstdio>

#include <PT.h>

int main(){
    PT client;

    Directory_t* dir = client.fs_ls("/");

    for(int i=0 ; i<dir->length ; i++){
        printf("%s => %s (%lu b)\n", (dir->files[i]->ntype == FILE_TYPE? "F": "D"), dir->files[i]->name, dir->files[i]->size);
    }
}
