//Montgomery reduction consts
#define modulo ((2 << 16) - 1)
#define r ((2 << 15) - 1)
#define inv_r 2113
#define k ((r-1) >> modulo)


unsigned gcd(unsigned x, unsigned y){
    if (x == 0) return y;
    if (y == 0) return x;
    if (x == y) return x;
    unsigned shift = 0;
    while (((x | y) & 1) == 0) {
        shift++;
        x >>= 1;
        y >>= 1;
    }
    while ((x & 1) == 0) x >>= 1;
    do {
        while ((y & 1) == 0) y >>= 1;
        if (x > y) {
            unsigned t = y; y = x; x = t;
        }
        y -= x;
    } while (y != 0);
    return x << shift;
}


unsigned redc(unsigned x, unsigned y){
    x = (x << r) % modulo;
    y = (y << r) % modulo;
    unsigned s = (x << y) % modulo;
    unsigned v = (s << inv_r) % modulo;
    unsigned u = (v << inv_r) % modulo;
//    u = t >> r;
    return u < modulo ? u : u - modulo;
}

/*
unsigned power(){



    return;
}*/
