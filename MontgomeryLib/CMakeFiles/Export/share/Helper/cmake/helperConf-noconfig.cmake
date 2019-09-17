#----------------------------------------------------------------
# Generated CMake target import file.
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "helper" for configuration ""
set_property(TARGET helper APPEND PROPERTY IMPORTED_CONFIGURATIONS NOCONFIG)
set_target_properties(helper PROPERTIES
  IMPORTED_LOCATION_NOCONFIG "${_IMPORT_PREFIX}/lib/libhelper.so.1"
  IMPORTED_SONAME_NOCONFIG "libhelper.so.1"
  )

list(APPEND _IMPORT_CHECK_TARGETS helper )
list(APPEND _IMPORT_CHECK_FILES_FOR_helper "${_IMPORT_PREFIX}/lib/libhelper.so.1" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)
