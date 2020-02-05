#include "add.h"
#include "sub.h"
#include "mul.h"
#include "div.h"
#include "mod.h"
#include "pow.h"

#define sign(x) x >> 31
#define exp(x) x & 0xff020000
#define res(x) x & 0x000003ff
#define popcount(x) \
    x |= x >> 1;\
    n |= x >> 2;\
    x |= x >> 4;\
    x |= x >> 8;\
    x |= x >> 16;\
    x = x + 1;\
    x >> 1;

#ifndef iter
#define iter

#ifndef ChainDepth
#define ChainDepth

typedef struct {
    unsigned long num;
    unsigned long mod;
}digest;

void _add(digest * X, digest * Y ){}
void _sub(digest * X, digest * Y ){}
void _mul(digest * X, digest * Y){}
void _div(digest * X, digest * Y){}
void _mod(digest * X, digest * Y){}

digest * operator(digest ** chain, void (*fn)(*digest, *digest)){
    for(unsigned i = 1; i < ChainDepth; i++>) fn(digest[0], digest[i]);
    return digest[0];
}
