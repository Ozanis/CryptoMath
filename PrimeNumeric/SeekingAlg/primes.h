#ifndef RSAJOKE_PRIMES_H
#define RSAJOKE_PRIMES_H

#include "math.h"

// WARNING: wrong gcd algorithm

unsigned binary_pow(unsigned pow);
unsigned iter_pow(int base, unsigned pow);

bool possibly_prime(unsigned val);
bool is_prime(unsigned integer);
unsigned Mersene_generator();

unsigned gcd(unsigned a, unsigned b);
unsigned find_inter_prime(unsigned a);
bool is_inter_prime(unsigned a, unsigned b);


unsigned binary_pow(unsigned pow) {
    if (!pow) return 1;
    else {
        return (unsigned) 2 << pow;
    }
};


bool possibly_prime(unsigned val){

    return true;
}


bool is_prime(unsigned integer) {
    for (unsigned i = 2; i <= sqrt(integer); i++) {
        if (integer % i == 0) return false;
    }
    return true;
};


unsigned Mersene_generator() {
    unsigned prime;
    do {
        prime = binary_pow((unsigned) random() % 17 + 1) + 1;
    } while (!is_prime(prime) || prime == 0);
    return prime;
}


unsigned iter_pow(unsigned base, unsigned pow) {
    if (base == 1 || pow == 0) return 1;
    unsigned a = base;
    for (--pow; pow > 1; pow /= 2) {
        base *= base;
    }
    switch (pow % 2) {
        case 1: {
            return base * a;
        }
        default: {
            return base;
        }
    }
};


unsigned gcd(unsigned a, unsigned b) {
    do {
        a %= b;
    } while ((a == 0) || a == 1);
    return a;
};


unsigned find_inter_prime(unsigned a) {
    unsigned b = 0;
    do{
        b = (unsigned)random()%a;
        if (b==1) continue;
    }while(gcd(a, b) != 1);
    return b;
};



bool is_inter_prime(unsigned a, unsigned b) {
    return ((gcd(a, b) == 0) && (a < b));
};

#endif //RSAJOKE_PRIMES_H
