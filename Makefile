libitdc:
	cd lib && go build -o libitdc.a -buildmode=c-archive && sed -E -i 's/typedef (.*) _Complex GoComplex.*;//gm' libitdc.h && mv libitdc.h include/

libitdcmm: libitdc
	cd itdcmm && g++ -fPIC -shared PT.cpp -I include/ -I ../lib/include -L ../lib/ -litdc -o libitdcmm.so

python: libitdcmm
	cd swig && swig -Wall -python -o PT.cxx -c++ PT.i && g++ -shared -o _PT.so -I../itdcmm/include -I../lib/include -L../itdcmm -L../lib -fPIC PT.cxx `pkg-config --libs --cflags python3` -litdcmm

test:
	g++ test.cpp -o test -Iitdcmm/include -Ilib/include -Litdcmm -Llib -litdcmm -litdc -lpthread

clean:
	rm -f lib/include/libitdc.h *.so itdcmm/*.so test swig/*.cxx swig/PT.py swig/*.so
