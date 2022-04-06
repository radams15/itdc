//
// Created by rhys on 05/04/2022.
//

#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/stat.h>
#include <inttypes.h>
#include <fcntl.h>

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
    // [3, 0, 0, 0, 16, 39, 0, 0, 136, 19, 0, 0, 1, 0, 2, 0, 11, 11, 3, 0, 10, 0, 0, 0, 250, 4, 0, 0, 3, 0, 0, 0]

    char* path = calloc(256, sizeof(gopath));
    strncpy(path, gopath, strlen(path));

    printf("File: '%s'\n", gopath);

    struct stat sb;
    sb.st_size = 0;

    while(sb.st_size == 0) {
        if (stat(gopath, &sb) == -1) {
            perror("stat");
            exit(EXIT_FAILURE);
        }
    }

    printf("File size: %d\n", sb.st_size);

    FILE* f = fopen(gopath, "rb");

    char* bytes = (char*) malloc(sb.st_size);

    fread(bytes, 1, sb.st_size, f);

    fclose(f);

    return bytes;
}