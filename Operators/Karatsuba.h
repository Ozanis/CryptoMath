#ifndef FIELDS_KARATSUBA_H
#define FIELDS_KARATSUBA_H

#include "malloc.h"
#include "pthread.h"

typedef struct {
    int * var;
    unsigned length;
}vector;

typedef struct{
    vector ** vectors;
    unsigned width;
    unsigned height;
}matrix;

/**
 * vector -> matrix -> vector -> int
 * exactly: _ml_(vector * args) -> gen_vector(intnum) -> as_bytes(int per_arg) ->
 * gen_matrix(vector ** vector_of_bytes) -> Karatsuba_ml(matrix * of_byte_vectors) -> recombine(vector)
**/

int inverse(int num, unsigned mod);
vector * gen_vector(int num);
vector * as_bytes(vector * num);
matrix * gen_matrix(vector ** vector_of_bytes, unsigned count);

int recombine(vector * bytes);
vector * Karatsuba_ml(matrix * bytes);
int _ml_(vector * args);

int Karatsuba_pw(int base, int pow);
int Karatsuba_sqr(int a);

vector * gen_vector(int num){
    vector * new_vector = malloc(sizeof(vector));
    int temp = num;
    for(new_vector->length=0; temp>0; temp /= 10, new_vector->length++);
    new_vector->var = malloc(new_vector->length*sizeof(int));
    new_vector->length = new_vector->length;
    new_vector->var[0] = num;
    return new_vector;
}

vector * as_bytes(vector * num){
    int temp = num->var[0];
    for(unsigned i = 0; i < num->length; i++){
        num->var[i] = temp % 10;
        temp /= 10;
    }
    return num;
}

matrix * gen_matrix(vector ** vector_of_bytes, unsigned count){
    matrix * M = malloc(sizeof(matrix));
    M->vectors = vector_of_bytes;
    M->width = count;
    return M;
};


int Karatsuba_ml(matrix * bytes){



    return 0;
};

int _ml_(vector * args){
    for(args->length; args->length > 0; args->length--){

    };


    for(unsigned i = 0; i < args->length; i++){
        as_bytes(args->var[i]);
    }


};



int Karatsuba_pw(int base, int pow){


};










#endif //FIELDS_KARATSUBA_H
