#ifndef FIELDS_ITERATIVEPOW_H
#define FIELDS_ITERATIVEPOW_H

unsigned modular_pow(int base, unsigned pow, unsigned modulus);
unsigned doubling_pow(unsigned base, unsigned pow);


unsigned modular_pow(int base, unsigned pow, unsigned modulus) {
    unsigned a = base % modulus;
    for (--pow; pow > 1; pow >>= 2) {
        base %= modulus;
        base *= base;
    }
    if(pow%2==1) return (base * a) % modulus;
    else return base % modulus;
};

unsigned doubling_pow(unsigned base, unsigned pow) {
    if (pow == 0) return 1;
    else {
        return base << pow;
    }
};



#endif //FIELDS_ITERATIVEPOW_H
