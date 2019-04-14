#ifndef FIELDS_MERSENNE_H
#define FIELDS_MERSENNE_H

#include "../FastOperators/IterativePow.h"


unsigned Mersene_generator();

unsigned Mersene_generator() {
    unsigned prime;
    do {
        prime = binary_pow((unsigned)random() % 17 + 1) + 1;
    } while (!is_prime(prime) || prime == 0);
    return prime;
}


#endif //FIELDS_MERSENNE_H
