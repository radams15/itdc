#include <stdio.h>
#include <stdlib.h>

#include <libitdc.h>

int main(){
    PT client = pt_new();

    char* vers = pt_get_vers(client);

    printf("Version: %s\n", vers);

    pt_free(client);
}
