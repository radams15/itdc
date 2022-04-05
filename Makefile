libitdc:
	cd lib && go build -o libitdc.so -buildmode=c-shared && mv libitdc.h include/

libitdcmm: libitdc
	cd itdcmm && g++ -fPIC -shared PT.cpp -I include/ -I ../lib/include -L ../lib/ -o libitdcmm.so

test:
	g++ test.cpp -o test -Iitdcmm/include -Ilib/include -Litdcmm -Llib -litdcmm -litdc

clean:
	rm -f lib/include/libitdc.h *.so itdcmm/*.so test
