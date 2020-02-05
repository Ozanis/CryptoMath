#ifndef LFSR_CONFIG_H
#define LFSR_LFSR_H

#include "Pool.h"

typedef struct {
    unsigned num_of_conditions;
    ull64 (*DistrFN[])(ull64);
}Sieve;

typedef struct {
    ull64 mask;
    ull64 scrembler;
}Config;

ull64 LFSR_Galois(ull64 State, Config * Prop){
    switch (State & 0x00000001) {
        case 1:
            return ((State ^ Prop->mask) >> 1) | Prop->scrembler;
        default:
            return State >> 1;
    }
}

ull64 Generate(ull64 number, Sieve * laticce, Config * prop){

}

#endif //LFSR_CONFIG_H
