#ifndef LFSR_DISTRIBUTIONS_H
#define LFSR_DISTRIBUTIONS_H

#include "additionalMath.h"

ull64 onlyOdd(ull64 num, ull64 param){
    return (num & 1)?num:0;
}

ull64 onlyEven(ull64 num, ull64 param){
    return (num & 0)?num:0;
}

ull64 FermatNum(ull64 num, ull64 param){ // Fermat`s test
    return power_r(num, --num, param)==1?num:0;
}

ull64 inRange(ull64 num, ull64 param){
    return num % param;
}

ull64 aliquotBy(ull64 num, ull64 param){
    return (num & param)?num:0;
}

#endif //LFSR_DISTRIBUTIONS_H
