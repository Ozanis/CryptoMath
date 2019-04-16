#ifndef FIELDS_KARATSUBA_H
#define FIELDS_KARATSUBA_H

#include "malloc.h"
#include "pthread.h"

typedef struct {
    int * var;
    unsigned len;
}vector;

typedef struct{
    vector ** vectors;
    unsigned len;
}matrix;

/**
 * vector -> matrix -> vector -> int
 * exactly: _ml_(vector * args) -> gen_vector(int num, int size) -> as_bytes(int per_arg) ->
 * parallel_ml(matrix * bytes) -> Karatsuba_ml(matrix * of_byte_vectors) -> recombine(vector)
**/

int inverse(int num, unsigned mod);
unsigned def_size(vector * num);
vector * gen_vector(int num, unsigned size);
vector * as_bytes(vector * num);
matrix * parallel_ml(matrix * bytes);
int recombine(matrix * bytes);
vector * Karatsuba_ml(vector * A, vector * B);

int _ml_(vector * args);
int Karatsuba_pw(vector * base, int pow);
int _pw_(int base, unsigned pow);
int Karatsuba_sqr(int a);
int _sqr_(int a);

unsigned def_size(vector * num){
    int temp = 0;
    unsigned max_size = 0, size = 0;
    for(num->len; num->len > 0; temp = num->var[num->len--]) {
        for (size; temp > 0; temp /= 10, size++);
        if(size>max_size) max_size = size;
        size = 0;
    }
    free(num);
    return max_size;
}

vector * gen_vector(int num, unsigned size){
    vector * new_vector = malloc(sizeof(vector));
    int temp = num;
    new_vector->var = malloc(new_vector->len*sizeof(int));
    new_vector->len = new_vector->len;
    new_vector->var[0] = num;
    return new_vector;
}

vector * as_bytes(vector * num){
    int temp = num->var[0];
    for(unsigned i = 0; i < num->len; i++){
        num->var[i] = temp % 10;
        temp /= 10;
    }
    return num;
}

vector * Karatsuba_ml(vector * A, vector * B){

};

matrix * parallel_ml(matrix * bytes) {
    switch (bytes->len) {
        case 1:
            return bytes;
        default: {
            matrix *low_oreder = malloc(sizeof(matrix));
            low_oreder->len = bytes->len / 2;
            low_oreder->vectors = malloc(sizeof(vector) * low_oreder->len);
            for (unsigned i = 0, j = 0; i < bytes->len; i += 2, j++) {
                low_oreder->vectors[j] = Karatsuba_ml(bytes->vectors[i], bytes->vectors[i + 1]);
            };
            free(bytes);
            parallel_ml(low_oreder);
        }
    }
}

int recombine(matrix * bytes){
    int result = 0;
    vector * V = bytes->vectors[0];
    V->len = bytes->len;
    free(bytes);
    for(unsigned i = 0; i < V->len; i++){
        result += V->var[i]*10^i;
    }

    return result;
}

int _ml_(vector * args){
    matrix * bytes = malloc(args->len);
    bytes->vectors = malloc(args->len);
    bytes->len = def_size(args);
    for(unsigned i = 0; i < args->len; i--){
        bytes->vectors[i] = as_bytes(gen_vector(args->var[i], bytes->len));
    };
    if(bytes->len % 2 == 1) Karatsuba_ml(bytes->vectors[bytes->len], bytes->vectors[bytes->len--]);
    return recombine(parallel_ml(bytes));
};

int _pw_(int base, unsigned pow){

    return 0;
};


int Karatsuba_pw(vector * base, int pow){


};






#endif //FIELDS_KARATSUBA_H
