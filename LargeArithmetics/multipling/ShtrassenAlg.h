#ifndef FIELDS_SHTRASSENALG_H
#define FIELDS_SHTRASSENALG_H

#include "pthread.h"
#include "malloc.h"

typedef struct {
    int * var;
    unsigned len;
}vector;

typedef struct{
    int ** vectors;
    unsigned len;
}matrix;


unsigned def_size(int * args, unsigned num){
    unsigned max_size = 0;
    for(unsigned i = 0, size = 0; i < num; i++){
    for(int temp = args[i]; temp > 0; temp /= 10) {
        ++size;
        }
    if(size>max_size) max_size = size;
    size = 0;
    }
    free(args);
    return max_size;
}

int * as_bytes(int num, unsigned size){
    int * bytes = malloc(size*sizeof(int));
    for(int temp = 0, i =0; i<size ; i++){
        bytes[i] = temp % 10;
    }
    return bytes;
}

int recombine(int * bytes, unsigned size){
    int result = 0;
    for(int j = 0, d = 1; j < size; j++){
        result += bytes[j]*d;
        d *= 10;
    }
    free(bytes);
    return result;
}

int decomposition(vector * V){

    return 0;
}

int primitive(int a, int mod){

}

int DTF(int a, ){

    return 0;
}


int IDFT(vector * V){

    return 0;
}


int _ml_(int * args, unsigned num){
    int ** bytes = malloc(sizeof(int)*num);
    unsigned size = def_size(args, num);
    for(unsigned i =0, temp = 0; i<0; i++){
        bytes[i] = as_bytes(args[i], size);   // exec int stream

    }
    for(unsigned i  =0; i < ){

        }

    }

    return recombine();
}


#endif //FIELDS_SHTRASSENALG_H
