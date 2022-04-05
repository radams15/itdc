project(itdc)

cmake_minimum_required(VERSION 3.0)

set(ITDC_INCLUDE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/lib/include)
set(ITDCMM_INCLUDE_DIR ${CMAKE_CURRENT_SOURCE_DIR}/itdcmm/include)
set(ITDC_LINK_DIR ${CMAKE_BINARY_DIR}/lib)

add_custom_target(libitdc
        COMMAND go build -o libitdc.so -buildmode=c-shared && mkdir -p ${CMAKE_BINARY_DIR}/lib && mkdir -p include && mv libitdc.so ${CMAKE_BINARY_DIR}/lib && mv libitdc.h include/
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/lib
        VERBATIM)

add_library(itdcmm SHARED itdcmm/PT.cpp)
target_include_directories(itdcmm PUBLIC ${ITDC_INCLUDE_DIR})
target_link_directories(itdcmm PUBLIC ${ITDC_LINK_DIR})
target_link_libraries(itdcmm PUBLIC itdc)

add_executable(test test.cpp)
target_include_directories(test PRIVATE ${ITDCMM_INCLUDE_DIR})
target_link_libraries(test PRIVATE itdcmm)

install(FILES ${CMAKE_BINARY_DIR}/lib/libitdc.so DESTINATION lib)
install(FILES ${ITDC_INCLUDE_DIR}/libitdc.h DESTINATION include/itdc)