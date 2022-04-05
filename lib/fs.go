package main

/*
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

#ifdef __cplusplus
extern "C" {
#endif

void set_file(struct Node** nodes, struct Node* to_add, int index);

void free_node(Node_t* node);

void free_dir(struct Directory* dir);

#ifdef __cplusplus
}
#endif
*/
import "C"

import (
	"unsafe"

	"io/ioutil"
	"path/filepath"

	"go.arsenm.dev/itd/api"
)

//export pt_fs_ls
func pt_fs_ls(ptr uintptr, path *C.char) (*C.Directory_t){
	client := *(*api.Client)(unsafe.Pointer(ptr))
    listing, _ := client.ReadDir(C.GoString(path))

    out := (*C.Directory_t)(C.malloc(C.size_t(C.sizeof_Directory_t)))
    out.length = (C.ulong)(len(listing))

    out.files = (**C.Node_t)(C.malloc(out.length*C.size_t(C.sizeof_Nodes_t)))

    for i:=0 ; i<len(listing) ; i++ {
        node := (*C.Node_t)(C.malloc(C.size_t(C.sizeof_Node_t)))

        node.name = C.CString(listing[i].Name)
        node.size = (C.ulong)(listing[i].Size)

        if(listing[i].IsDir){
            node.ntype = (uint32)(C.DIR_TYPE)
        }else{
            node.ntype = (uint32)(C.FILE_TYPE)
        }

        C.set_file(out.files, node, (C.int)(i));

    }

    return out;
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