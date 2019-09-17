cmake_minimum_required(VERSION 3.10.2)

project(Montgomery)

include(GNUInstallDirs)

set(CMAKE_CXX_FLAGS_DEBUG "${CMAKE_CXX_DEBUG} -Wall")

add_library(${PROJECT_NAME} SHARED src/REDC.c )

target_include_directories(${PROJECT_NAME} PRIVATE .)

set_target_properties(${PROJECT_NAME} PROPERTIES
    VERSION 1
    SOVERSION 1
    PUBLIC_HEADER REDC.c)

install(TARGETS ${PROJECT_NAME} EXPORT MontgomeryLibConfig
    ARCHIVE  DESTINATION ${CMAKE_INSTALL_LIBDIR}
    LIBRARY  DESTINATION ${CMAKE_INSTALL_LIBDIR}
    RUNTIME  DESTINATION ${CMAKE_INSTALL_BINDIR})
install(DIRECTORY include/ DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}/${PROJECT_NAME})

install(EXPORT MontgomeryLibConfig DESTINATION share/MontgomeryLib/cmake)

export(TARGETS ${PROJECT_NAME} FILE MontgomeryLibConfig.cmake)


