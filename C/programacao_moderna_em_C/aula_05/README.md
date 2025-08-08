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