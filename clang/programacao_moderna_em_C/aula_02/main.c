#include <stdio.h>
#include <stdbool.h>

int main(void)
{
    unsigned c = 129;
    bool b;

    printf("Tamanho do tipo CHAR: %llu byte\n", sizeof c);
    printf("Tamanho do tipo INT: %llu bytes\n", sizeof(int));
    printf("Tamanho do tipo FLOAT: %llu bytes\n", sizeof(float));
    printf("Tamanho do tipo DOUBLE: %llu bytes\n", sizeof(double));
    printf("Tamanho do tipo U INT: %llu bytes\n", sizeof(unsigned int));
    printf("Tamanho do tipo L INT: %llu bytes\n", sizeof(long int));
    printf("Tamanho do tipo LL INT: %llu bytes\n", sizeof(long long int));
    printf("Tamanho do tipo U LL INT: %llu bytes\n", sizeof(unsigned long int));
    printf("Tamanho do tipo NULL: %llu bytes\n", sizeof(NULL));
    printf("Tamanho do tipo NULL: %llu bytes\n", sizeof(0));
    printf("%i\n", c);

    printf("Tamanhho do tipo BOOLEANO: %llu byte\n", sizeof b);
    b = true;
    printf("Booleano: %i\n", b);

    return 0;
}