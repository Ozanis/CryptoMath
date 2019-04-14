#ifndef FIELDS_GCD_H
#define FIELDS_GCD_H

// WARNING: wrong gcd algorithm
unsigned gcd(unsigned a, unsigned b);
unsigned find_inter_prime(unsigned a);
unsigned gen_inter_prime();
unsigned short is_inter_prime(unsigned a, unsigned b);


unsigned gcd(unsigned a, unsigned b) {
    do {
        a %= b;
    } while ((a == 0) || a == 1);
    return a;
};

unsigned gen_inter_prime(){

    return 0;
}

unsigned find_inter_prime(unsigned a) {
    unsigned b = 0;
    do{
        b = gen_inter_prime()%a;
        if (b==1) continue;
    }while(gcd(a, b) != 1);
    return b;
};



unsigned short is_inter_prime(unsigned a, unsigned b) {
    return ((gcd(a, b) == 0) && (a < b));
};




#endif //FIELDS_GCD_H
