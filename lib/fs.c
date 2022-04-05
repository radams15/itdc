//
// Created by rhys on 05/04/2022.
//

#include <stdlib.h>

#include "include/libitdc.h"

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
