#----------------------------------------------------------------
# Generated CMake target import file.
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "Montgomery" for configuration ""
set_property(TARGET Montgomery APPEND PROPERTY IMPORTED_CONFIGURATIONS NOCONFIG)
set_target_properties(Montgomery PROPERTIES
  IMPORTED_LOCATION_NOCONFIG "${_IMPORT_PREFIX}/lib/libMontgomery.so.1"
  IMPORTED_SONAME_NOCONFIG "libMontgomery.so.1"
  )

list(APPEND _IMPORT_CHECK_TARGETS Montgomery )
list(APPEND _IMPORT_CHECK_FILES_FOR_Montgomery "${_IMPORT_PREFIX}/lib/libMontgomery.so.1" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)
