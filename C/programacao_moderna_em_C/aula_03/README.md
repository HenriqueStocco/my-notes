#  Aula 03  [Curso: Programação Moderna em C - Mente Binária]

## INT_MAX

**INT_MAX** vem da biblioteca `<limits.h>`, onde ele informa o maior número inteiro com sinal. Mas esta biblioteca também mostra o tamanho de outros tipos com o char (referente a aula passada), tanto com sinal quanto sem sinal, muito interessante.

```sh
# Em distribuições baseado no Debian, para encontrar alguma biblioteca
# utilizar o comando locate (é necessário instalar o pacote plocate e atualizar a base de dados do OS)

# Para instalar o pacote plocate
sudo apt install plocate

# Atualizar a base de dados
sudo updatedb

# Encontrar a biblioteca padrão `limits.h`
locate limits.h

# Copiar o caminho do arquivo e utilizar o visualiador baseado em vim, mas é read-only
view /usr/include/limits.h # este caminho pode mudar de distro para distro
```

## Modificadores para tipos sem sinal (unsigned)

```c
// Para números sem sinal usamos o modificador %u
printf("unsigned int - %u\n", ui);

// Para números em octal usamos o modificador %o
// Este modificador converte o números decimais em octal
printf("Octal - %o\n", oct);

// Para hexadecimais com letras minusculas usamos o modificador %x
// Para hexadecimais com letras maiúsculas usamos o modificador %X
// Este modificador converte números decimais em hexadecimais
printf("Hexadecimal - %x ou %X\n", hex);

```

## Tamanhos de inteiros e formas de declaração

Até agora, nossos estudos mostram que o int por padrão permite alocar até 4 bytes (32 bits) de memória.
Mas isso depende da arquitetura e do compilador. Aqui estamos utilizando o GCC, mas o Turbo C define que
um int vai alocar no máximo 2 bytes (16 bits) de memória. Isso em alguns casos pode ser ruim por que, na teoria nunca vamos saber extamente o tamanho daquele int em um código, e para facilitar a vida do sofredor, existe uma biblioteca padrão (header), chamada, <stdint.h>, que define o tamanho de tipos, então nós conseguimos utilizar os tipos com tamanho definido sem dor de cabeça.

> Toda biblioteca padrão vem com a extensão `.h`, isso significa _**header**_
> devido a toda pré-compilação ele "copiar" o código desta biblioteca para o código a ser compilado.

```c
// Biblioteca
// ...
#include <limits.h>
#include <stdint.h>

int main(void) {
    uint32_t i32 = UINT_MAX; // variavel sem sinal de 32 bits (4 bytes) com o valor máximo de 32 bits
    uint16_t i16 = SHRT_MAX; // variavel sem sinal de 16 bits (2 bytes) com o valor máximo de 16 bits
    uint8_t i8 = UCHAR_MAX; // variavel sem sinal de 8 bits (1 byte) com o valor máximo de 8 bits (char tem 8 bits, por isso CHAR_MAX)
}
```

Também existe os "limitadores" _*(não sei como chamam)*_, de tamanho dos tipos

```c
short int i; // short reduz pela metade o tamanho do inteiro, neste caso para 2 bytes (16 bits)
short i; // também pode ser escrito sem o int que o compilador entende

long int i; // long aumenta 4 bytes no tamanho do inteiro, nesse caso para 8 bytes (64 bits)
// (Na minha arquitetura cabem apenas 32 bits em long int).
long i; // mesma coisa do short

long long int i; // long long (na minha arquitetura) cabem 8 bytes (64 bits)
long long i; // mesma coisa do long

// Esses "limitadores" também possuem o seu valor sem sinal, apenas adicionar o unsigned atrás
unsigned long long batata;
```

## size_t

Este é o tipo exato do retorno do operador `sizeof`. Um tipo sem sinal de 8 bytes ou 64 bits (dependendo da arquitetura).

```c
#include <stdlib.h>

// ...
size_t len;

printf("SIZE_T: %zu bytes\n", sizeof len);

// ...
```

## Modificador `register`

Modificador register faz com que a variável seja alocada em um registrador, que é um espaço pequeno e mais rápido de momória, dentro do processador, onde ele não precisa acessar a memória RAM para salvar, buscar e manipular um dado, o que leva muito tempo.

> Normalmente utilizado em contadores de loops, justamente pela sua velocidade.

```c
register int i = 10;
```