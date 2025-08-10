#  Aula 05  [Curso: Programação Moderna em C - Mente Binária]

## Arrays

Um array é definido assim:

```c
char a[3];
a[0] = { 'a' };
a[1] = { 'b' };
a[2] = { 'c' };

int b[4] = { 1, 2, 3, 4 };
```

## Tamanho do array

O tamanho do array (em memória depende do tipo e da quantidade de elementos)

Se for um int por exemplo, ficaria:

```c
int a[3];

printf("%zu\n", sizeof a);

// Resultado: 4 (tamanho de int em bytes) * 3 (quantos elementos) = 12 bytes
```

Mas e se quisermos saber o comprimento em relação aos elementos, fazemos:

```c
char array[4];

// Forma 1
size_t length = sizeof array / sizeof array[0];

// OU
size_t length = sizeof array / sizeof (int);
```

## Conversão implicita de modificadores

Quando utilizamos o tipo `char` _*(int também, como vi anteriormente)*_, o compilador faz uma conversão automaticamente, quando o valor não é um char, ele busca caracteres _*(no caso de valores hexa e decimais)*_, na tabela **ASCII**.

```c
c[0] = 'A';
c[1] = 0x42;
c[2] = 67;

printf("c[0]: %c\nc[1]: %c\nc[2]: %c\n", c[0], c[1], c[2]);
/*
c[0]: A
c[1]: B
c[2]: B
*/ 

// Ou os valores númericos
printf("c[0]: %d\nc[1]: %d\nc[2]: %d\n", c[0], c[1], c[2]);

/* 
c[0]: 65
c[1]: 66
c[2]: 67
*/
```

## Como funciona o Array

O array nada mais é que um espaço na memória, em sequência, onde todos os valores de um array são armazenados um do lado do outro.

O array é um ponteiro, basicamente sempre aponta para o primeiro indice na memoria, e quando você utiliza uma estrutura de loop, por exemplo, como é uma sequencia na memória, ele pega o indice atual (o endereço), e incrementa, soma +1 para buscar o próximo.

E a diferença entre um ponteiro para um array é que, o array sempre aponta para endereços fixos de memória, já o ponteiro, para endereços variáveis.

> [!NOTE]
>
> O modificador para ponteiros é `%p`
>
> O operador `&` é o operador de endereço de memória

```c
// Pegando o endereço de memória do array (sempre aponta para o primeiro indice)
char c[3];

printf("Address: %p\n", c);
// >>> Address: 0x000000058273627
```

Uma prova de que e o array é um ponteiro e sempre aponta para o primeiro elemento, é fazer com que o endereço some 1 e pegue o endereço de memória dessa nova posição.

> [!NOTE]
>
> O `*` na variavel c, quer dizer buscar o valor de um endereço de memoria, no caso o ponteiro, valor do ponteiro.
>
> Este também é um operador _*(operador indireto como é chamado na documentação GNU C - Pointers)*_, ele é responsável por criar um ponteiro, um armazenador de endereços de memóra, mas quando temos um endereço de memória e queremos o valor, utilizar ele no printf, por exemplo, faz com que ele pegue o valor, já que o ponteiro já um endereço de memória, isso se chama desreferenciar um ponteiro.

```c
char c[3];
c[0] = 'A';
c[1] = 0x42 // que é 66 que é o caracter B na ASCII
c[2] = 67 // que é C em ASCII

printf("Valor de A: %c\n", *c);
printf("Valor de B: %c\n", *(c + 1));
```

Aqui é importante o comentário, o `*` é um operador e um operador com a precedência maior do que o operador de soma _*(+)*_ e de subtração _*(-)*_, então, se apenas somar a desreferencia do array com um número, por exemplo:

```c
printf("Próximo indice: %c\n", *c + 1);
```

Ele vai pegar o valor do primeiro indice e vai identificar o seu numerico (valor na tabela ASCII), em seguida ele vai somar o número a este valor, no caso se for A, o próximo é B, e de assim em diante.

Obs: Isso `*c` é o mesmo que isso `c[0]`, e isso `*c + 1` é o mesmo que `c[0] + 1` e isso `(*c) + 1`.

Se quisermos pegar o próximo valor do array, manualmente, fazemos assim:

```c
printf("Indice 0: %c\n", *c);
printf("Indice 1: %c\n", *(c + 1));
```

Agora sim, ele está somando o primeiro indice com 1 e desreferenciando o endereço de memória para obter o valor armazenado.

Obs: Isso `*(c + 1)` é o mesmo que `c[0 + 1]`

É interessante, também, observar que os elementos são armazenados em sequência mas não de 1 byte para 1 byte (no caso pode ser se o tipo ocupar 1 byte), mas sim de acordo com o tipo do array, se for um int, será de 4 em 4, se for um long int, de 8 em 8, e assim sucessivamente, prova disso é observar os endereços em sequeência, através do operador de endereço de memória, veja:

```c
long int c[3]; // No meu caso tive que utilizar long long devido ao data model do windows (long int é o mesmo que int em tamanho.)
c[0] = 1;
c[1] = 2;
c[3] = 3;

printf("Endereco de memoria do indice 0: %p\n", &c);
printf("Endereco de memoria do indice 1: %p\n", &c[1]);
printf("Endereco de memoria do indice 2: %p\n", &c[2]);

/* Resultado:
Endereco de memoria do indice 0: 000000b42e5ffb10  8 bytes
Endereco de memoria do indice 1: 000000b42e5ffb18 +8 bytes
Endereco de memoria do indice 2: 000000b42e5ffb20 +8 bytes
*/
```