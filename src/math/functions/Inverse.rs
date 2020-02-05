#ifndef FIELDS_MULTIPLICATIVEINVERSE_H
#define FIELDS_MULTIPLICATIVEINVERSE_H

#include "../multipling/Karatsuba.h"
#include "Euler.h"

unsigned short is_invert(unsigned a, unsigned b, unsigned mod);
unsigned invert(unsigned a, unsigned mod);

unsigned short is_invert(unsigned a, unsigned b, unsigned mod){
    if((a*b)%mod == 1) return 1;
    return 0;
};

unsigned invert(unsigned a, unsigned mod){
    return Karatsuba_pw(a, Euler(mod)) % mod;
}


#endif //FIELDS_MULTIPLICATIVEINVERSE_H
