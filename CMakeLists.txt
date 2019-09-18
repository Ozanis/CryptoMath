cmake_minimum_required(VERSION 3.10.2)

project(Cryptomath)

set(CMAKE_C_STANDARD 11)

set(CMAKE_C_COMPILER /usr/bin/gcc)

include_directories(DiscreteLog ECC RSA LargeArithmetics RandomGenerators  MontgomeryLib)


#SET_TARGET_PROPERTIES(RSAjoke PROPERTIES LINKER_LANGUAGE CXX)
