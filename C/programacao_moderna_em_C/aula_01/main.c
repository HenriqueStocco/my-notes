#include <stdio.h>
#include "soma.h"

int main(void)
{
    puts("Hello, World!");
    // Função do meu arquivo de modulo
    // int resultado = soma(2, 4);
    // ret = printf("Minha soma de 2 + 4 = %d\n", resultado);
    int ret = 0;
    // Exibindo o retorno da função printf
    // ret = printf("Hello, World!");
    // printf("Valor do printf: %d\n", ret);

    // Converte os valores para decimal
    printf("%d - %d - %d - %i - %d - %d\n", ret, 10, 0xa, 'A', 0b1010, 012);
    // Converte os valores para hexadecimal
    printf("%x - %x - %x - %x - %x - %x\n", ret, 10, 0xa, 'A', 0b1001, 020);
    // Converte os valores para octal
    printf("%o - %o - %o - %o - %o - %o", ret, 10, 0xa, 'A', 0b1001, 020);
    // Usando o caractere de quebra de linha, igual o \n
    printf("%c", 0xa);
    printf("%c", 0xa);
    puts("Teste");
    printf("%c", 0xa);

    return 0;
}