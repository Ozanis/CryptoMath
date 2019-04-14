#ifndef FIELDS_ELIPTICCURVEPARAMS_H
#define FIELDS_ELIPTICCURVEPARAMS_H

#include "../Operators/IterativePow.h"
#include "Clousure.h"
#include "Lezhandr.h"
#include "malloc.h"
#include "math.h"

typedef struct {
    unsigned a;
    unsigned b;
    unsigned Field;
    unsigned Charact;
    unsigned D;
    unsigned Inv;
    unsigned Range[2];

}EllipticCurve;

EllipticCurve * init_curve(unsigned a, unsigned b, unsigned Field);
unsigned F_x(EllipticCurve * Curve, unsigned x);
void Discriminant(EllipticCurve * Curve);
void Hasse(EllipticCurve * Curve);
void Charact_iter(EllipticCurve * Curve);

unsigned F_x(EllipticCurve * Curve, unsigned x){
    return (Karatsuba_pw(x, 3) + Karatsuba_ml(x, Curve->a) + Curve->b) % Curve->Field;
};

void Discriminant(EllipticCurve * Curve){
    unsigned A = Karatsuba_ml(Karatsuba_pw(Curve->a, 3), 4) % Curve->Field;
    Curve->D = (A + Karatsuba_ml(27,Karatsuba_pw(Curve->b, 2))) % Curve->Field;
    unsigned args={1728, A, Karatsuba_pw(Curve->D, Curve->Field-2) % Curve->Field}; // {1728, 4*a^3, D^-1} % p = {1728, 4*a^3, D^(Phi(p)-1)%p} % p
    Curve->Inv = multi_mul(&args, Curve->Field, 4);
};

void Hasse(EllipticCurve * Curve){
    Curve->Range[0]=Curve->Range[1]=Curve->Field+1;
    unsigned defect = 2 * (unsigned)sqrt(Curve->Field);
    Curve->Range[0] -= defect;
    Curve->Range[1] += defect;
};

void Charact_iter(EllipticCurve * Curve){
    Curve->Charact = Curve->Field + 1;
    for(unsigned i = 0; i<Curve->Field; Curve->Charact+=Lezhander(F_x(Curve, i++), Curve->Charact));
};


EllipticCurve * init_curve(unsigned a, unsigned b, unsigned Field){
    EllipticCurve * Curve = malloc(sizeof(EllipticCurve));
    Curve->a = a;
    Curve->a = b;
    Curve->Field = Field;
    Discriminant(Curve);
    Hasse(Curve);
    Charact_iter(Curve);
    return Curve;
}


typedef struct {
    unsigned x;
    unsigned y;
}EllipticDot;


unsigned short dot_exist(EllipticCurve * Curve, unsigned x, unsigned y);
EllipticDot * create_dot(EllipticCurve * Curve, unsigned x, unsigned y);

unsigned short dot_exist(EllipticCurve * Curve, unsigned x, unsigned y){
    if((Karatsuba_sqr(y) % Curve->Field) == F_x(Curve, x)) return 1;
    else return 0;
}

EllipticDot * create_dot(EllipticCurve * Curve, unsigned x, unsigned y){
    EllipticDot * dot = malloc(sizeof(EllipticDot));
    if(dot_exist(Curve, x, y)) {
        dot->x = x;
        dot->y = y;
        return dot;
    }
    else{
        return 0;
    }
};


typedef struct{
    EllipticDot * r;
    unsigned Field;
    unsigned m;
}EllipticSum;


unsigned lambda(EllipticDot * p, EllipticDot * q, unsigned field);
EllipticSum * elliptic_sum(EllipticDot * p, EllipticDot * q, unsigned field);
EllipticSum * elliptic_mul(EllipticSum * Sum, EllipticCurve * Curve, unsigned seq);


unsigned lambda(EllipticDot * p, EllipticDot * q, unsigned field){
    return Karatsuba_ml((p->y - q->y), Karatsuba_pw((p->x - q->x), field-2));
};


EllipticSum * elliptic_sum(EllipticDot * p, EllipticDot * q, unsigned field){
    EllipticSum * Sum = malloc(sizeof(EllipticSum));
    Sum->Field = field;
    Sum->m = lambda(p, q, field);
    Sum->r->x = (Karatsuba_pw(Sum->m, 2) - p->x - q->x) % field;
    Sum->r->y = (q->y + Karatsuba_ml(Sum->r->x - q->x, Sum->m)) % field;
    return Sum;
}

EllipticSum * elliptic_mul(EllipticSum * Sum, EllipticCurve * Curve, unsigned seq){
    EllipticSum * mul = Sum;
    for(unsigned i =1; i <= seq; i++){
        mul = elliptic_sum(mul->r, Sum->r, Sum->Field);
    }
    return mul;
}

#endif //FIELDS_ELIPTICCURVEPARAMS_H
