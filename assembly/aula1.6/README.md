# Aula 1.5 - Vetores

Em section .data, as "variáveis" não são variáveis, são labels que apontam para algum lugar na memória para o primeiro caractere.

De alguma forma, por mais que venha a ter uma outra "variavel" abaixo, mas o tamanho que for definido na saida para imprimir o valor de msg1, ele ira imprimir tudo que estiver dentro do tamanho definido, seja ele definido ou um lixo qualquer, exemplo:

```asm
section .data
    msg1 db "Hello", 0xA, 0x0
    len equ $- msg1
    msg2 db "World", 0xA, 0x0

section .text
global _start
_start:

    mov eax, 0x4
    mov ebx, 0x1
    mov ecx, msg1
    mov edx, 0x24
    int 0x80

;   ...
; Resultado: 
; Hello
; World
; @% _ 0 3 - simboliza lixo na memória  
;  %%2....
```

## Acessando indices de uma sequência de valores

Mover a linha de memória (buscando por endereço) de um valor armazenado, como um valor dentro de uma variável, para um registrador, somando a posição dele, no caso abaixo, a variavel vai ter um tamanho de 4 bytes (32 bits), ou seja, a cada 4 bytes estará um dos valores definidos, exemplo:

```asm
section .data
    ;        0x0, 0x4, 0x8, 0x16, 0x24 - simulando endereços de memória
    array: dd 10,  20,  30,   40,   50

section .text
    global _start
    _start:
        mov eax, [array + 4] ; o indice começa em 0
        ; >>> 20

```

Dado o exemplo, é possivel associar o motivo pelo qual o primeiro indice sempre é zero, sempre teremos o endereço de memória da primeira posição, mas somando pelo tamanho de bytes de cada item no array e multiplicando pela posição que desejamos, obtemos o indice esperado, segue:

```asm
section .data
    array: dd 10, 20, 30, 40, 50, 60

section .text
global _start
_start:
    mov eax, [array + 4 * 0] ; [array + tamanho_tipo * indice] ->
```