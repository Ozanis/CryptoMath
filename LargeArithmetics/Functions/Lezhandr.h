#ifndef FIELDS_LEZHANDR_H
#define FIELDS_LEZHANDR_H

#include "../multipling/Karatsuba.h"

unsigned Lezhander(unsigned a, unsigned p);

unsigned Lezhander(unsigned a, unsigned p){
    unsigned l = p-1;
    int res = Karatsuba_pw(a, (l)>>2) % p;
    if (res==l) return -1;
    else return res;
}
#endif //FIELDS_LEZHANDR_H
