#ifndef FIELDS_MONAD_H
#define FIELDS_MONAD_H

#include "../multipling/Karatsuba.h"


int multi_mul(unsigned mod, ...);
int multi_sum(unsigned mod, ...);
int Closure(void (*fn)(unsigned, ... ), unsigned mod);

int multi_mul(unsigned mod, ...){
    unsigned var =1;
    for(unsigned * p = &mod; ){
        var = Karatsuba_ml(var, ) % mod;
    }
    return var;
}

int multi_sum(int mod, ...){
    int var = 0;
    for(arg_num; arg_num > 0; arg_num--) {
        var += args[arg_num]%= mod;
    }
    return var;
}

int Closure(void (*fn)(unsigned, ... ), unsigned mod){

}

#endif //FIELDS_MONAD_H
