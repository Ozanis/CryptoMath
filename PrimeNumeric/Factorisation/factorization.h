#ifndef RSAJOKE_FACTORIZATION_H
#define RSAJOKE_FACTORIZATION_H

#include "../SeekingAlg/primes.h"

typedef struct {
    unsigned *primes;
    unsigned *sequences;
    unsigned num;
    unsigned of_var;
} primes;

primes * prime_factorization(unsigned val);

primes * prime_factorization(unsigned val){};

#endif //RSAJOKE_FACTORIZATION_H
