libitdc:
	cd lib && go build -o libitdc.so -buildmode=c-shared

test:
	gcc test.c -o test -Ilib -Llib -litdc

clean:
	rm -f libitdc.h *.so
