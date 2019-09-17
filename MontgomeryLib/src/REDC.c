//Montgomery reduction consts
#define modulo ((2 << 16) - 1)
#define r ((2 << 15) - 1)
#define inv_r 2113
#define k ((r-1) >> modulo)

typedef unsigned u16;

u16 gcd(u16 x, u16 y){
    if (x == 0) return y;
    if (y == 0) return x;
    if (x == y) return x;
    u16 shift = 0;
    while (((x | y) & 1) == 0) {
        shift++;
        x >>= 1;
        y >>= 1;
    }
    while ((x & 1) == 0) x >>= 1;
    do {
        while ((y & 1) == 0) y >>= 1;
        if (x > y) {
            u16 t = y; y = x; x = t;
        }
        y -= x;
    } while (y != 0);
    return x << shift;
}


u16 redc(u16 x, u16 y){
    x = (x << r) % modulo;
    y = (y << r) % modulo;
    u16 s = (x << y) % modulo;
    u16 v = (s << inv_r) % modulo;
    u16 u = (v << inv_r) % modulo;
//    u = t >> r;
    return u < modulo ? u : u - modulo;
}

/*
u16 power(){



    return;
}*/
