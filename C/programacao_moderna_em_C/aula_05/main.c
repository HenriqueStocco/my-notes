#include <stdio.h>

int main(void)
{
    char c[3];
    size_t len = sizeof c / sizeof c[0];

    printf("O tamanho de c eh: %zu bytes / %zu bits\n", sizeof c, sizeof c * 8);
    printf("O número de elementos de c eh: %zu\n", len);
    return 0;
}