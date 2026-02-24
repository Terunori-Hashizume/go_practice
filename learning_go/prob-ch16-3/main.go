package main

import "fmt"

/*
#include <string.h>

int mini_calc(char *op, int a, int b) {
    if (strcmp(op, "+") == 0) {
        return a + b;
    }
    if (strcmp(op, "*") == 0) {
        return a * b;
    }
    if (strcmp(op, "-") == 0) {
        return a - b;
    }
    if (strcmp(op, "/") == 0) {
        if (b == 0) {
            return 0;
        }
        return a / b;
    }
    return 0;
}
*/
import "C"

func main() {
	op_add := C.CString("+")
	op_sub := C.CString("-")
	op_mul := C.CString("*")
	op_div := C.CString("/")

	fmt.Printf("Mini_calc add: 6 + 2 = %d\n", C.mini_calc(op_add, 6, 2))
	fmt.Printf("Mini_calc sub: 6 - 2 = %d\n", C.mini_calc(op_sub, 6, 2))
	fmt.Printf("Mini_calc mul: 6 * 2 = %d\n", C.mini_calc(op_mul, 6, 2))
	fmt.Printf("Mini_calc div: 6 / 2 = %d\n", C.mini_calc(op_div, 6, 2))
}
