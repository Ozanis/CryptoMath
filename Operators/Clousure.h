#ifndef FIELDS_MONAD_H
#define FIELDS_MONAD_H

#include "Karatsuba.h"

unsigned multi_mul(unsigned *args, unsigned mod, unsigned arg_num);

unsigned multi_sum(unsigned *args, unsigned mod);

unsigned multi_mul(unsigned *args, unsigned mod, unsigned arg_num){
    unsigned var = 1;
    for(arg_num; arg_num > 0; arg_num--) {
        var = Karatsuba_ml(var, args[arg_num]) % mod;
    }
    return var;
}

unsigned multi_sum(unsigned *args, unsigned mod){
    unsigned var = 0;
    for(arg_num; arg_num > 0; arg_num--) {
        var += args[arg_num];
        var %= mod;
    }
    return var;
}

//Karatsuba_ml(Karatsuba_ml(Karatsuba_ml(Curve->D, Curve->Charact-1), A), 1728)
#endif //FIELDS_MONAD_H
