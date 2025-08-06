#include <stdio.h>

int main(void) {
    float f = 1;
    float f1 = 3e2; // matematicamente: 3 * 10 ^ 2 = 300.00

    printf("Float sizeof: %zu bytes\n", sizeof f);
    printf("Float value: %f\n", f);
    printf("Float value: %.2f\n", f1);
    
    return 0;
}