#ifndef FIELDS_EULER_H
#define FIELDS_EULER_H

#include "../PrimeNumeric/Factorisation/factorization.h"
#include "Karatsuba.h"

unsigned Euler(unsigned p);
unsigned recombine(NumericList * factors);

unsigned recombine(NumericList * factors){
    unsigned  number = 1;
    for(unsigned i = 1; i<=factors->ind; i++){
        number =Karatsuba_ml(number, Karatsuba_pw(factors->Node[i]->prime), factors->Node[i]->sequence);
    }
    return number;
};

unsigned Euler(unsigned p){
    if (is_prime){
        return --p;
    }
    else{
        return recombine(prime_factorization(p));
    };
};

#endif //FIELDS_EULER_H