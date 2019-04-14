#ifndef FIELDS_ELIPTICCURVEPARAMS_H
#define FIELDS_ELIPTICCURVEPARAMS_H

#include "../GeneralField/ECCfield.h"
#include "malloc.h"
#include "math.h"

typedef struct{
    unsigned P;
}Field;

unsigned Monad(unsigned p);

typedef struct {
    unsigned a;
    unsigned b;
    unsigned Charact;
    unsigned D;
}EllipticCurve;

unsigned F(unsigned x);
unsigned Discriminant(unsigned a, unsigned b);
unsigned Inv();
unsigned * Hasse(unsigned p);
unsigned Sequence(unsigned x);
unsigned Lezhander(void * f);

unsigned F(unsigned x){
    return
};









typedef struct {
    unsigned x;
    unsigned y;
}EllipticDot;

typedef struct{
    EllipticDot * p;
    EllipticDot * q;
    unsigned m;
}EllipticSum;



//class EllipticCurve: public Field{
//private:
//    unsigned A;
//    unsigned B;
//    unsigned Charact;
//public:
//    explicit EllipticCurve(unsigned power, unsigned a, unsigned b) : Field(power), A(a), B(b), Charact(0) {};
//    ~EllipticCurve() = default;
//    unsigned F(unsigned x);
//    unsigned def_sequence_Lezhandr();
//    unsigned Hasse();
//};
//
//unsigned EllipticCurve :: F(unsigned x){
//    return (x^3+x*this->A+this->B) % this->N;
//};
//
//unsigned EllipticCurve :: def_sequence_Lezhandr(){
//    unsigned x;
//    for(x=0; x<=this->N--; x++){
//        x+=Lezhandr(F(x), this->N);
//    }
//    return x+1+this->N;
//};
//
//unsigned EllipticCurve :: Hasse(){
//    return 1 + this->N + 2*((unsigned)ceil(sqrt(this->N)));
//}
//
//
//class EllipticDot: public GeneralField{
//private:
//    unsigned X;
//    unsigned Y;
//public:
//    unsigned m;
//    explicit EllipticDot(unsigned power, unsigned x, unsigned y): GeneralField(power), Y(y), X(x), m(0){};
//    ~EllipticDot() = default;
//    unsigned lambda(EllipticDot * dot1, EllipticDot * dot2);
//    unsigned lambda(EllipticDot * dot1, EllipticDot * dot2, unsigned a);
//    unsigned dot_sequence_iterable();
//    EllipticDot add_x(EllipticDot * dot1, EllipticDot * dot2);
//    EllipticDot add_y(EllipticDot * dot1, EllipticDot * dot2);
//};
//
//unsigned EllipticDot :: lambda(EllipticDot * dot1, EllipticDot * dot2){
//    return multipling(substr(dot_1->y, dot_2->y, dot_1->n), multiplicative_inverse(substr(dot_1->x, dot_2->x, dot_1->n), dot_1->n), dot_1->n);
//};
//
//unsigned EllipticDot :: lambda(EllipticDot * dot1, EllipticDot * dot2, unsigned a){
//    return multipling(adding(multipling(3, dot->x^2, dot->n), a, dot->n), multiplicative_inverse(multipling(2, dot->y, dot->n), dot->n), dot->n);
//};
//
//unsigned EllipticDot :: add_x(EllipticDot * dot1, EllipticDot * dot2){
//    return (dot1->X^2-2*dot_2->x) % dot_1->n;
//};
//
//unsigned EllipticDot :: add_y(EllipticDot * dot1, EllipticDot * dot2){
//    return adding(dot_2->y, multipling(m_different(dot_1, dot_2), substr(add_x(dot_1, dot_2), dot_2->x, dot_1->n), dot_1->n), dot_1->n);
//};
//
//unsigned EllipticDot ::  def_dot_sequence_iterable(eliptic_dot * dot){
//    dot->m = m_different(dot, dot);
//    eliptic_dot * temp = dot;
//    unsigned sequence;
//    for(sequence = 1; ((dot->x != 0) && (dot->y != 0)); sequence++){
//        temp->m = m_different(temp, temp);
//        temp->x = add_x(temp, dot);
//        temp->y = add_y(temp, dot);
//        printf("x= %u\n", temp->x);
//        printf("y= %u\n", temp->y);
//        sequence +=1;
//    }
//    return sequence;
//};

#endif //FIELDS_ELIPTICCURVEPARAMS_H
