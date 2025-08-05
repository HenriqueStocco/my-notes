# Aula 01 - Curso: Programação Moderna em C - Mente Binária

>[!NOTE]
> Comando para a compilação + execução do objeto
>
> Powershell
>
> `clear ; gcc -o <obj_name> -Wall -Werror --pedantic-errors <filename.c> ; ./<obj_name>`
>
> Bash
> 
> `clear && gcc -o <obj_name> -Wall -Werror --pedantic-errors <filename.c> && ./<obj_name>`

## Importação de módulos externos

> `#include` é uma diretiva de pré-processamento, onde recebe um argumento de caminho de arquivo.
> O conteúdo do arquivo é "copiado" para o programa na compilação, para utilizar as funções deste
> que foi importado no programa atual
>
> `<stdlib.h>` indica o caminho da biblioteca padrão, e está entre `<>` por que este caminho é 
> padrão.

Importando a biblioteca padrão de C, onde podemos utilizar a função ``printf`

```c
#include <stdio.h>

int main(void) {
    ...
    printf("Hello, World!");
    return 0;
}
```

Importando a minha própria biblioteca/módulo

> Neste caso o arquivo está no mesmo caminho que o meu código.

```c
#include <soma.h>

int main(void) {
    int resultado = soma(2, 2);

    printf("Soma de 2 + 2 = %d\n", resultado);
    return 0;
}
```

## Função `printf`

A função `printf` vem da biblioteca padrão de C, ela recebe 1 ou mais argumento/parâmetros e retorna um valor do tipo `int`, referente à quantidade de caractéres imprimidos, ex:

```c
int main(void) {
    int retorno = printf("Hello, World!");

    printf("Retorno de printf: %d\n", retorno); // >>> 13
    return 0;
}
```

A função `printf` converte alguns valores de acordo com os identificadores de tipo, como o `%d`.

> `%d` permite receber apenas valores inteiros, mas se forem hexadecimal, octal e caractéres ASCII
> ele faz uma conversão para decimal (o identificador `%i` também faz).

```c
int main(void) {
    printf("Decimal: %d - Hexa: %d - ASCII: %i - Binário: %d - Octal: %d", 10, 0xa, 'A', 0b1010, 012);
    /* 
    Decimal: 10 - Hexa: 10 - ASCII: 64 - Binario: 10 - Octal: 10
    */
    return 0;
}
```

Isso ocorre com hexadecimais utilizando o identificador `%x`, ele converte tudo para valores hexadecimais, ex:

> Isso vale para octal também com o identificador `%o`
```c
int main(void) {
    printf("Decimal: %x - Hexa: %x - ASCII: %x - Binário: %x - Octal: %x", 10, 0xa, 'A', 0b1010, 012);
    /* 
    Decimal: 10 - Hexa: 10 - ASCII: 64 - Binario: 10 - Octal: 10
    */
    return 0;
}
```

Interessante também, consigo imprimir os caractéres `ASCII`, utilizando a estratégia de conversão da função `printf`, por exemplo, imprimindo o caractére de `break-line`, em hexa é `0xa` que é 10 em decimal e 012 em octal. Passamos o identificador de caracter para a função `printf`, e isso é convertido para o `ASCII` correspondente, ex:

```c
int main(void) {
    printf("Hello, World!");
    printf("%c", 0xa); // ele insere o mesmo caracter de \n
    ...
}
```