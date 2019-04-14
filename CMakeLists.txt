cmake_minimum_required(VERSION 3.10.2)

project(Fields)

set(CMAKE_C_STANDARD 11)

set(CMAKE_C_COMPILER /usr/bin/gcc)

include_directories(DiscreteLog ECC GeneralField PrimeNumeric)

add_executable(test_log DiscreteLog/test_log.c)
add_executable(test_ECC ECC/test_ECC.c)
add_executable(test_GeneralField GeneralField/test_field.c)
add_executable(test_factor PrimeNumeric/Factorisation/test_factor.c)
add_executable(test_prime PrimeNumeric/PrimesTests/test_primes.c)
add_executable(test_seeking PrimeNumeric/SeekingAlg/test_seeking.c)

add_executable(Fields main.c)


#SET_TARGET_PROPERTIES(RSAjoke PROPERTIES LINKER_LANGUAGE CXX)