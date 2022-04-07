%module PT

%{
#include <PT.h>
#include <Settings.h>
%}

%include "std_string.i"
%include "../itdcmm/include/PT.h"
%include "../lib/include/libitdc.h"
%include "../lib/include/Settings.h"