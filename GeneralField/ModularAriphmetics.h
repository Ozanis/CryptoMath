#ifndef RSAJOKE_RAND_H
#define RSAJOKE_RAND_H

#include "../PrimeNumeric/SeekingAlg/primes.h"

bool is_inverse(unsigned a, unsigned a_, unsigned modulus);
unsigned modular_pow(int base, unsigned pow, unsigned modulus);
unsigned multiplicative_inverse(unsigned number, unsigned modulus);
unsigned Euler(unsigned long long n);
unsigned Euler_pq(unsigned p, unsigned q);
bool is_integer(double val);
unsigned quadra_pow(unsigned base, unsigned pow);

bool is_inverse(unsigned a, unsigned a_, unsigned modulus) {
    return ((a * a_) % modulus == 0);
};


unsigned modular_pow(int base, unsigned pow, unsigned modulus) {
    unsigned a = base % modulus;
    for (--pow; pow > 1; pow /= 2) {
        base %= modulus;
        base *= base;
    }
    return (base * a) % modulus;
};


unsigned multiplicative_inverse(unsigned a, unsigned modulus) {
    return
}

bool is_integer(double val) {
        return val == floor(val * 1000000) / 1000000;
};


unsigned Euler_pq(unsigned p, unsigned q) {
    return (p--) * (q--);
};


unsigned Euler(unsigned long long n) {
 }


unsigned quadra_pow(unsigned base, unsigned pow) {
    if (pow == 0) return 1;
    else {
        return base << pow;
    }
};

#endif //RSAJOKE_RAND_H
