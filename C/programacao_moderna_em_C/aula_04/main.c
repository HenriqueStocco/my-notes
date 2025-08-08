#include <stdio.h>
#include <limits.h>
#include <stdint.h>
#include <stdlib.h>

int main(void){
    float f = 1;
    double d = 1;

    printf("O tamanho  de f (float): %zu bytes / %zu bits\n", sizeof f, sizeof f * 8);
    printf("O tamanho  de d (double): %zu bytes / %zu bits\n", sizeof f, sizeof f * 8);
    printf("Valor de f: %.2f\n", f);

    return 0;
}