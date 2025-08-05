#include <stdio.h>
#include <limits.h>
#include <stdint.h>

int main(void)
{
    // i = UINT_MAX;
    // i++; // i = i + 1 ou i += 1
    int simx = INT_MAX;
    unsigned int uimx = UINT_MAX;
    short si = SHRT_MAX; // int de 16 bits
    long li = LONG_MAX; // int de 64 bits (cabe 8 bytes mas ele me mostra apenas 4 bytes, investigarei posteriormente, vi que depende da arquitetura, então cabem apenas 32 bits na arquitetura do meu processador) 
    long long lli = LONG_LONG_MAX; // agora sim os 64 bits

    printf("O tamanho de i (int): %zu bytes / %zu bits\n", sizeof(int), sizeof(int) * 8);
    printf("SHORT INT: %llu bytes / %llu bits\n", sizeof si, sizeof si * 8);
    printf("LONG INT: %llu bytes / %llu bits\n", sizeof li, sizeof li * 8);
    printf("LONG LONG INT: %llu bytes / %llu bits\n", sizeof lli, sizeof lli * 8);
    printf("O valor de simx: %d\n", simx);
    printf("O valor de uimx: %u\n", uimx);
    printf("UIMX em Octal: %o\n", uimx);
    printf("UIMX em Hexadecimal: %X\n", uimx);
    return 0;
}