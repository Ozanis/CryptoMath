#ifndef FIELDS_KARATSUBA_H
#define FIELDS_KARATSUBA_H

#include "malloc.h"

unsigned _bytes_num(unsigned val);

typedef struct {
    unsigned * var;
    unsigned len;
} bytes;

typedef struct{
    bytes * A;
    bytes * B;
} Karatsuba;

unsigned Karatsuba_ml(unsigned left, unsigned right);
unsigned Karatsuba_pw(unsigned base, unsigned pow);
unsigned Karatsuba_sqr(unsigned a);






unsigned Karatsuba_pw(unsigned base, unsigned pow){

    return _bytes_num(base);
};


unsigned Karatsuba_ml(unsigned left, unsigned right){

    return 0;
};








#endif //FIELDS_KARATSUBA_H
