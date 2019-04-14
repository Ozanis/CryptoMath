#ifndef FIELDS_KARATSUBA_H
#define FIELDS_KARATSUBA_H

unsigned Karatsuba_pw(unsigned base, unsigned pow);
unsigned Karatsuba_ml(unsigned left, unsigned right);
unsigned Karatsuba_sqr(unsigned a);


unsigned Karatsuba_pw(unsigned base, unsigned pow){

    return 0;
};


unsigned Karatsuba_ml(unsigned left, unsigned right){

    return 0;
};

unsigned Karatsuba_sqr(unsigned a){
    return Karatsuba_ml(a, a);
};

#endif //FIELDS_KARATSUBA_H
