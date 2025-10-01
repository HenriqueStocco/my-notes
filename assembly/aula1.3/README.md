# Aula 1.3 - Comparação de valores

## Definição de bytes

Em assembly não um "tipo" diretamente, mas, temos a definição de espaços que serão utilizados ou separados para a futura utilização em bytes.
Para isso utilizamos os tipos de dados:

- db - define byte (1 byte = 8 bits)
- dw - define word (2 bytes = 16 bits)
- dd - define double word  (4 bytes = 32 bits)
- dq - define quad word (8 bytes = 64 bits)
- dt - define ten word (10 bytes = 80 bits)

## Operador de comparação

Em assembly, a forma de comparação é diferente de linguagens de mais alto nível, fazemos uma comparação direta entre dois registradores (no caso os valores para qual esses registradores apontam na memória).

O operador é o **CMP** de comparison, exemplo:

```asm
section .data
    x dd 10
    y dd 20

section .text
    ;...

    mov EAX, DWORD[x]
    mov EBX, DWORD[y]
    cmp EAX, EBX ; compara EAX com EBX
```

De acordo com a explicação do GPT, a comparação na verdade é uma subtração entre, por exemplo, o EAX com EBX, e o resultado altera as FLAGS do registrador RFLAGS.
Este registrador possui as flags:

- **ZF** *(zero flag)*     -> se o resultado foi 0 *(x == y)*
- **SF** *(sign flag)*     -> se o resultado foi negativo *(x < y>)*
- **OF** *(overflow flag)* -> usado para detectar overflow aritmético em signed *(divisão por zero, acredito)*
- **CF** *(carry flag)*    -> usado em *unsigned* comparações

### Operadores lógicos

Existem 3 operadores lógicos, estes são: AND, OR e XOR.

#### AND

O AND irá procurar valores que tenham seus binários iguais a 1, em ambos os números, exemplo:

7 - 0111
5 - 0101
R - 0101 -> resultado, todos os números que não forem iguais a um em ambos os valores, será convertido para 0, caso contrário, convertido para 1.

#### OR

O OR irá procurar a ocorreência onde o número 1 apareça pelo menos uma vez, exemplo:

7 - 0111
5 - 0101
R - 0111 -> resultado, toda comparação que conter um único um ou dois 1s, será 1, caso contrário, será 0.

#### XOR

O XOR é o contrário de AND, toda ocorrência que ele encontrar ambos iguais a 1, ele converterá para 0, exemplo:

7 - 0111
5 - 0101
R - 0010

O XOR é utilizado para retornar o valor 0 na finalização do programa, indicando que ocorreu tudo de acordo. O XOR zera todos os valores por que os bits são iguais, veja:

```àsm
xor EBX, EBX ; significa (1 == 1) = 0 | (0 == 0) = 0
```

## Jumps

O Jump é o que direciona o fluxo do programa, ele **PULA** para uma parte do código, se queremos pular diretamente ou de acordo com uma condição.

Normalmente, quando utilizamos o operador de comparação, como dito acima, ele compara (subtrai os valores) e altera alguma/algumas FLAGS do registrador RFLAGS, e isso impacta na forma como vamos direcionar o que vamos fazer depois da comparação.

Para direcionar o que vamos fazer, usamos os chamados *jumps*, que há duas categorias de jumps, condicionais e incondicionais.

### Jumps Condicionais

Os Jumps condicionais, como o própria sub-titulo sugere, ele vão para uma direção de acordo com a condição das FLAGS do registrador RFLAGS.
Os são esses:

- **JE** *(jump if equal)* -> ele pula para uma label quando os valores são iguais. (se a FLAG ZF do registrador RFLAGS for igual a 1)
- **JG** *(jump if greater)* -> ele pula para uma label quando o valor é maior que outro. (se a FLAG ZF for igual a zero e se a FLAG SF for igual a FLAG OF)
- **JGE** *(jump if greater or equal)* -> ele pula para uma label quando o valor é maior ou igual a outro. (se a FLAG SF for igual a FLAG OF)
- **JL** *(jump if less)* -> ele pula para uma label quando o valor é menor que outro. (se a FLAG SF for diferente da FLAG OF)
- **JLE** *(jump if less or equal)* -> ele pula para uma label quando o valor é menor ou igual a outro. (se a FLAG ZF for igual a 1 e a FLAG SF for diferente da FLAG OF)
- **JNE** *(jump if not equal)* -> ele pula para uma label quando o valor é diferente de outro. (se a FLAG ZF for igual a zero)

### Jump Incondicional

Este é o *JMP*, é o famoso *GOTO* de linguagens como **C**, ele pulam quando eu definir que deve pular, sem nenhuma condição, ele não depende de nenhuma FLAG do registrador RFLAGS, vai direto para a *label* que eu definir.

## Curiosidade

O nome da seção de constantes que terão valores atribuidos no futuro, a `.bss`, significa: **Block Starting Symbol**

## "Tipo" de variáveis não inicializadas

As variáveis **não** inicializadas, são aquelas que terão o seu valor atribuidos ao decorrer do programa, abaixo uma tabela de cada um com seu tipo *(segundo o livro do Fernando Anselmo)*:

| Sigla | Tipo        | Significado          |
| ----- | ----------- | -------------------- |
| resb  | byte        | variavel de 8 bits   |
| resw  | word        | variavel de 16 bits  |
| resd  | double      | variavel de 32 bits  |
| resq  | quad        | variavel de 64 bits  |
| resdq | double quad | variavel de 128 bits |

## Conversão relativa

No código, fazemos uma conversão relativa, no x e no y, antes de mandar para os registradores eax e ebx, usando DWORD, pois não é possível mover o conteúdo de um DD diretamente para estes registradores.