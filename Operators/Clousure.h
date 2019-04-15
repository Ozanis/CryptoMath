#ifndef FIELDS_MONAD_H
#define FIELDS_MONAD_H

#include "Karatsuba.h"


unsigned multi_mul(unsigned mod, ...);
unsigned multi_sum(unsigned mod, ...);
unsigned Closure(void (*fn)(unsigned, ... ), unsigned mod);

unsigned multi_mul(unsigned mod, ...){
    unsigned var =1;
    for(unsigned * p = &mod; ){
        var = Karatsuba_ml(var, ) % mod;
    }
    return var;
}

unsigned multi_sum(unsigned mod, ...){
    unsigned var = 0;
    for(arg_num; arg_num > 0; arg_num--) {
        var += args[arg_num]%= mod;
    }
    return var;
}

unsigned Closure(void (*fn)(unsigned, ... ), unsigned mod){

}

//Karatsuba_ml(Karatsuba_ml(Karatsuba_ml(Curve->D, Curve->Charact-1), A), 1728)
#endif //FIELDS_MONAD_H
