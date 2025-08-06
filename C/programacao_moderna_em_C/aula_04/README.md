#  Aula 04  [Curso: Programação Moderna em C - Mente Binária]

## Tipo float

Tipo que representa os números reais, com ponto flutuante e/ou números com casas decimais.
Dependendo da arquitetura (novamente), ele ocupa 4 bytes ou 32 bits.

```c
float num = 1;
// Pela notação cientifica também é possivel, basta passar o número que queremos
// ele ja é base 10 por padrão
float f1 = 3e2 // matemáticamente: 3 * 10^2

printf("Float size: %zu bytes\n", sizeof num);
printf("Float value: %f\n", num);

// Imprimindo somente 2 casas decimais (por padrão são 6)
printf("2 Casas: %.2f\n", num);
```