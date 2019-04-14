#ifndef FIELDS_ECCFIELD_H
#define FIELDS_ECCFIELD_H

unsigned substr(unsigned a, unsigned b, unsigned modulus){
    return  (a - b) % modulus;
};

unsigned adding(unsigned a, unsigned b, unsigned modulus){
    return  (a + b) % modulus;
};

unsigned multipling(unsigned a, unsigned b, unsigned modulus){
    return  (a * b) % modulus;
};

unsigned multiplicative_inverse(unsigned a, unsigned b){
    return a^(b-1) % b;
};

unsigned Lezhandr(unsigned a, unsigned p){
    return a^((p-1)/2) % p;
}

//class GeneralField{
//public:
//    unsigned N;
//
//    explicit GeneralField(unsigned power): N(power){};
//    ~GeneralField() = default;
//
//    virtual GeneralField operator +(unsigned &a){
//
//    };
//    virtual GeneralField operator -(unsigned &a){};
//    virtual GeneralField operator *(unsigned &a){};
//    virtual GeneralField operator /(unsigned &a){};
//};

#endif //FIELDS_ECCFIELD_H
