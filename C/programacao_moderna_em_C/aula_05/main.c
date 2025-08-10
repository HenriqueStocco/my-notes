#include <stdio.h>

int main(void)
{
    // unsigned char uc[3];
    // char c[3];
    long long int c[3];
    // size_t len = sizeof c / sizeof c[0];

    // printf("O tamanho de c eh: %zu bytes / %zu bits\n", sizeof c, sizeof c * 8);
    // printf("O numero de elementos de c eh: %zu\n", len);

    // c[0] = 'J';
    // c[1] = 0x42;
    // c[2] = 67;
    c[0] = 1;
    c[1] = 2;
    c[2] = 3;

    // printf("c[0]: %c\nc[1]: %c\nc[2]: %c\n", c[0], c[1], c[2]);
    // printf("O endereco do array c em memoria: %p\n", c);
    // printf("O endereco do array &c em memoria: %p\n", &c);
    // printf("O endereco do primeiro elemento de c em memoria: %p\n", &c[0]);
    // printf("O endereco do primeiro elemento de c em memoria: %c\n", *c + 1);
    // printf("Valor do indice 0: %c\n", *c);
    // printf("Valor do indice 0: %c\n", c[0] + 1);
    // printf("Valor do indice 1: %c\n", *(c + 1));
    // printf("Valor do indice 1: %c\n", c[0 + 1]);
    printf("Endereco de memoria do indice 0: %p\n", &c);
    printf("Endereco de memoria do indice 1: %p\n", &c[1]);
    printf("Endereco de memoria do indice 2: %p\n", &c[2]);

    return 0;
}