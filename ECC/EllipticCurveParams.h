#ifndef FIELDS_ELIPTICCURVEPARAMS_H
#define FIELDS_ELIPTICCURVEPARAMS_H

#include "../Functions/Clousure.h"
#include "../Functions/Lezhandr.h"
#include "math.h"

typedef struct {
    int x;
    int y;
}EllipticDot;


typedef struct {
    EllipticDot ** dots;
    int Charact;
    int a;
    int b;
    int Seq;
    int D;
    int Inv;
    int Range[2];
}EllipticCurve;

EllipticCurve * init_curve(int a, int b, int Field);
int F_x(EllipticCurve * Curve, int x);                                    // %N
void Discriminant(EllipticCurve * Curve);                                           // %N
void Hasse(EllipticCurve * Curve);
void Charact_iter(EllipticCurve * Curve);

int dot_exist(EllipticCurve * Curve, int x, int y);                       // %N
EllipticDot * create_dot(EllipticCurve * Curve, int x, int y);            // %N
int lambda(EllipticDot * p, EllipticDot * q);                                  // %N
EllipticDot * elliptic_sum(EllipticDot * p, EllipticDot * q);                       // %N
EllipticDot * elliptic_mul(EllipticDot * R, EllipticCurve * Curve, unsigned seq); // %N


void attach_Curve(EllipticCurve * Curve);
void attach_Dot(EllipticDot ** Dots);


int F_x(EllipticCurve * Curve, int x){
    return (_pw_(x, 3) + Karatsuba_ml(x, Curve->a) + Curve->b);     // %N
};

void Discriminant(EllipticCurve * Curve){
    int A = _ml_(_pw_(Curve->a, 3), 4);
    Curve->D = (A + Karatsuba_ml(27, Karatsuba_pw(Curve->b, 2)));
    unsigned args={1728, A, Karatsuba_pw(Curve->D, Curve->Field-2)};    // %N     // {1728, 4*a^3, D^-1} % p = {1728, 4*a^3, D^(Phi(p)-1)%p} % p
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
}

EllipticCurve * init_curve(int a, int b, int Field){
    EllipticCurve * Curve = malloc(sizeof(EllipticCurve));
    Curve->a = a;
    Curve->a = b;
    Discriminant(Curve);
    Hasse(Curve);
    Charact_iter(Curve);
    return Curve;
}


int dot_exist(EllipticCurve * Curve, int x, int y){
    return _sqr_(y) == F_x(Curve, x);
}

EllipticDot * create_dot(EllipticCurve * Curve, int x, int y){
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


int lambda(EllipticDot * p, EllipticDot * q){
    return Karatsuba_ml((p->y - q->y), Karatsuba_pw((p->x - q->x), field-2));       // %N
};

EllipticDot * elliptic_sum(EllipticDot * p, EllipticDot * q){
    EllipticDot * R = malloc(sizeof(EllipticDot));
    int m = lambda(p, q);
    R->x = (Karatsuba_pw(m, 2) - p->x - q->x);      // %N
    R->y = q->y + _ml_(R->x - q->x, m);     // %N
    return R;
}

EllipticDot * elliptic_mul(EllipticDot * R, EllipticCurve * Curve, unsigned seq){
    EllipticDot * Mul = R;
    for(unsigned i =1; i <= seq; i++){
        Mul = elliptic_sum(Mul, R);                 // %N
    }
    return Mul;
}



#endif //FIELDS_ELIPTICCURVEPARAMS_H
