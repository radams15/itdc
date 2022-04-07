//
// Created by rhys on 05/04/2022.
//

#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/stat.h>
#include <inttypes.h>
#include <fcntl.h>
#include <string.h>

enum NodeType{
    FILE_TYPE = 0,
    DIR_TYPE = 1
};

typedef struct Node{
    char* name;
    unsigned long size;
    enum NodeType ntype;
} Node_t;

typedef Node_t *Nodes_t;

typedef struct Directory{
    struct Node** files;
    unsigned long length;
} Directory_t;

void set_file(struct Node** nodes, struct Node* to_add, int index){
    nodes[index] = to_add;
}

void free_node(Node_t* node){
    free(node->name);
    free(node);
}

void free_dir(struct Directory* dir){
    for(int i=0 ; i<dir->length ; i++){
        free_node(dir->files[i]);
    }

    free(dir->files);

    free(dir);
}

char* read_file(const char* gopath){
    char* path = calloc(256, sizeof(gopath));
    strncpy(path, gopath, strlen(path));

    struct stat sb;
    sb.st_size = 0;

    while(sb.st_size == 0) {
        if (stat(gopath, &sb) == -1) {
            perror("stat");
            exit(EXIT_FAILURE);
        }
    }

    FILE* f = fopen(gopath, "rb");

    char* bytes = (char*) malloc(sb.st_size);

    fread(bytes, 1, sb.st_size, f);

    fclose(f);

    return bytes;
}

void write_file(char* data, int len, char* path){
    FILE* f = fopen(path, "wb");

    fwrite(data, 1, len, f);

    fclose(f);
}