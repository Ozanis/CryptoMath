#ifndef FIELDS_ELIPTICCURVEPARAMS_H
#define FIELDS_ELIPTICCURVEPARAMS_H

#include "../FastOperators/IterativePow.h"
#include "Monad.h"
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
    return Karatsuba_pw(x, 3) + Karatsuba_ml(x, Curve->a) + Curve->b;
};

void Discriminant(EllipticCurve * Curve){
    unsigned A = Karatsuba_pw(Curve->a, 3);
    Curve->D = A + 27*Karatsuba_pw(Curve->b, 2);
    unsigned args={Curve->D, Curve->Field-1, A, 1728};
    Curve->Inv = monad(&Karatsuba_mul, &args, 4);
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


unsigned short dot_exist(EllipticCurve * Curve, EllipticDot * dot);

unsigned short dot_exist(EllipticCurve * Curve, EllipticDot * dot){
    if((Karatsuba_sqr(dot->y) % Curve->Field) == (F_x(Curve, dot->x) % Curve->Field)) return 1;
    else return 0;
}


typedef struct{
    EllipticDot * p;
    EllipticDot * q;
    unsigned m;
}EllipticSum;


EllipticSum * elliptic_sum(EllipticDot * p, EllipticDot * q);


#endif //FIELDS_ELIPTICCURVEPARAMS_H
