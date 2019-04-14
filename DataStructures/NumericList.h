#ifndef FIELDS_NUMERICLIST_H
#define FIELDS_NUMERICLIST_H

#include "malloc.h"

typedef struct {
    unsigned val;
    unsigned sequence;
}Pair;

typedef struct {
    Pair ** Node;
    unsigned ind;
}NumericList;


NumericList init_NL(NumericList * NL, unsigned size);
NumericList del_NL(NumericList * NL);
void add_NL_node(NumericList * NL, Pair * P);
void del_NL_node(NumericList * NL);

void add_NL_node(NumericList * NL, Pair * P){
    ++NL->ind;
    NL->Node[NL->ind] = malloc(sizeof(P));
};

#endif //FIELDS_NUMERICLIST_H
