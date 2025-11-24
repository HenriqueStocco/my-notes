#  Aula 02  [Curso: Programação Moderna em C - Mente Binária]

## Operador sizeof

O operador sizeof mostra o tamanho (na memória, em bytes), de um tipo.

```c
// ...code
char cr;

printf("Tamanho de CHAR: %llu\n byte", sizeof cr);
// >>> 1 

// Não é recomendado mas podemos verificar o tamanho do tipo diretamente
printf("Tamanho de CHAR: %llu bytes\n", sizeof (char));

// ...code
```

## Valores com e sem sinal

> O tipo `char` também permite utilizar números decimais.

O tipo `char` comporta apenas 1 _*byte*_ ou 8 _*bits*_ de informação.
Dentro destes 8 _*bits*_, em valores com **sinal**, o primeiro **número binário** indica o sinal do número, sendo ele **0 - positivo** e **1 - negativo**.
Isso implica no valor total que o tipo comporta, pois agora não temos mais **8** _*bist*_, e sim, apenas **7**.

> 7 _*bist*_ + o sinal, consegue armazenar o número **127** positivo e **-128** negativo.
>
> 8 _*bits*_ sem sinal armazena o valor de **255**.
>
> A metade de **255** é **127.50**, mas um binário não armazena números com casas decimais.

> [!NOTE]
>
> Os tipos (que permitem sinal), por padrão, vão colocar os valores com sinal
>
> Para definir que eles não terão sinal, adicionar à palavra-chave `unsigned` atrás do tipo
>
> ex: `unsigned char c;`
>
> O mesmo serve para valores com sinal explicito (desnecessário)
>
> ex: `signed char c;`
>

## Tipo Booleano (true | false)

> [!NOTE]
> Antigamente, antes de existir este "tipo", era utilizado 0 para **positivo** e 1 para **negativo**, em operações de comparação e etc.
>
> O que aconteceu, foi que, para deixar mais legivel e padronizado, criaram a biblioteca
> `stdbool.h`. Mas se entrarmos no código, veremos as diretivas de pré-processamento indicando que o valor de `false` é 0 e `true` é 1.
>
> Se imprimirmos o valor de `true` or `false`, veremos a prova disso também.
>

Os tipos `booleanos` em C não existem como tipo primitivo, para isso precisamos utilizar uma biblioteca padrão chamada `<stdbool.h>`.

```c
#include <stdbool.h>

bool b = true;
bool b = false;
```