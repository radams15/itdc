project(itdc)

cmake_minimum_required(VERSION 3.0)

set(ITDC_INCLUDE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/lib/include)
set(ITDC_LINK_DIR ${CMAKE_BINARY_DIR}/lib)

add_custom_target(libitdc
        COMMAND go build -o libitdc.so -buildmode=c-shared itdc.go && mkdir -p ${CMAKE_BINARY_DIR}/lib && mkdir -p include && mv libitdc.so ${CMAKE_BINARY_DIR}/lib && mv libitdc.h include/
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/lib
        VERBATIM)


add_executable(test test.c)

add_dependencies(test libitdc)
target_include_directories(test PRIVATE ${ITDC_INCLUDE_DIR})
target_link_directories(test PRIVATE ${ITDC_LINK_DIR})
target_link_libraries(test PRIVATE itdc)