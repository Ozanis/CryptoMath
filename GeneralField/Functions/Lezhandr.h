#ifndef FIELDS_LEZHANDR_H
#define FIELDS_LEZHANDR_H

#include "../FastOperators/Karatsuba.h"

unsigned Lezhander(unsigned a, unsigned p);

unsigned Lezhander(unsigned a, unsigned p){
    return Karatsuba_pw(a, (p-1)>>2);
}
#endif //FIELDS_LEZHANDR_H
