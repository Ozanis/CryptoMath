#ifndef FIELDS_ELIPTICCURVEPARAMS_H
#define FIELDS_ELIPTICCURVEPARAMS_H

#include "../Operators/IterativePow.h"
#include "Clousure.h"
#include "Lezhandr.h"
#include "malloc.h"
#include "math.h"

typedef struct {
    unsigned x;
    unsigned y;
}EllipticDot;


typedef struct {
    unsigned a;
    unsigned b;
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

int dot_exist(EllipticCurve * Curve, unsigned x, unsigned y);
EllipticDot * create_dot(EllipticCurve * Curve, unsigned x, unsigned y);
unsigned lambda(EllipticDot * p, EllipticDot * q);
EllipticDot * elliptic_sum(EllipticDot * p, EllipticDot * q);
EllipticDot * elliptic_mul(EllipticDot * R, EllipticCurve * Curve, unsigned seq);


typedef struct {
    unsigned N;
    EllipticCurve * Curve;
    EllipticDot ** Dots;
} Field;

Field * CreateField();
void attach_Curve(EllipticCurve * Curve);
void attach_Dot(EllipticDot ** Dots);


unsigned F_x(EllipticCurve * Curve, unsigned x){
    return (Karatsuba_pw(x, 3) + Karatsuba_ml(x, Curve->a) + Curve->b);
};

void Discriminant(EllipticCurve * Curve){
    unsigned A = Karatsuba_ml(Karatsuba_pw(Curve->a, 3), 4);
    Curve->D = (A + Karatsuba_ml(27,Karatsuba_pw(Curve->b, 2)));
    unsigned args={1728, A, Karatsuba_pw(Curve->D, Curve->Field-2)}; // {1728, 4*a^3, D^-1} % p = {1728, 4*a^3, D^(Phi(p)-1)%p} % p
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
    for(unsigned i = 0; i < Curve->Field; Curve->Charact+=Lezhander(F_x(Curve, i++), Curve->Field));
};

EllipticCurve * init_curve(unsigned a, unsigned b, unsigned Field){
    EllipticCurve * Curve = malloc(sizeof(EllipticCurve));
    Curve->a = a;
    Curve->a = b;
    Discriminant(Curve);
    Hasse(Curve);
    Charact_iter(Curve);
    return Curve;
}


int dot_exist(EllipticCurve * Curve, unsigned x, unsigned y){
    return Karatsuba_sqr(y) == F_x(Curve, x);
}

EllipticDot * create_dot(EllipticCurve * Curve, unsigned x, unsigned y){
    if(dot_exist(Curve, x, y)) {
        EllipticDot * dot = malloc(sizeof(EllipticDot));
        dot->x = x;
        dot->y = y;
        return dot;
    }
    else{
        return 0;
    }
};


unsigned lambda(EllipticDot * p, EllipticDot * q){
    return Karatsuba_ml((p->y - q->y), Karatsuba_pw((p->x - q->x), field-2));
};

EllipticDot * elliptic_sum(EllipticDot * p, EllipticDot * q){
    EllipticDot * R = malloc(sizeof(EllipticDot));
    unsigned m = lambda(p, q);
    R->x = (Karatsuba_pw(m, 2) - p->x - q->x);
    R->y = q->y + Karatsuba_ml(R->x - q->x, m);
    return R;
}

EllipticDot * elliptic_mul(EllipticDot * R, EllipticCurve * Curve, unsigned seq){
    EllipticDot * Mul = R;
    for(unsigned i =1; i <= seq; i++){
        Mul = elliptic_sum(Mul, R);
    }
    return Mul;
}



#endif //FIELDS_ELIPTICCURVEPARAMS_H
