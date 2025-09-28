# Aula 1.1 - Hello e Saida de dados STDOUT

## Introdução

Em assembly existem 3 seções que podemos utilizar para separar nosso código, essas são:

```asm
section .data ; Responsável por armazenar constantes, valores imutáveis.

section .bss ; Responsável por armazenar variáveis, que terão valores mutáveis.

section .text ; Responsável por armazenar o código, funções e etc, todo o restante.
```

Todo código assembly, igual ao código em linguagem C, vão ter um ponto de entrada, e este ponto é definido como:

```asm
global _start ; Esta é a parte obrigatória que todo código assembly deve conter para poder iniciar a execução do programa.
```
A seguir, podemos definir um label, onde irá começar o nosso código, a escrita dele, e o nome da label pode ser qualquer um.

```asm
_start: 
```
Comentários são com ponto e virgula, o que estiver a direita de um será ignorado na compilação:

```àsm
; isso é um comentário
```

## Registradores

Tabela contendo alguns registradores e suas equivalências em duas arquiteturas, registradores de 32 bits e 64 bits:

| 64 bits | 32 bits | Descrição                                                          |
| ------- | ------- | ------------------------------------------------------------------ |
| rax     | eax     | Valores que são retornados dos comandos em um registrador          |
| rbx     | ebx     | Registrador preservado. Cuidado ao utiliza-lo                      |
| rcx     | ecx     | Uso livre como por exemplo contador                                |
| rdx     | edx     | Uso livre em alguns comandos                                       |
| rsp     | esp     | Ponteiro de uma pilha                                              |
| rbp     | ebp     | Registrador preservado. Algumas vezes armazena ponteiros de pilhas |
| rdi     | edi     | Na passagem de argumentos, contém a quantidade desses              |
| rsi     | esi     | Na passagem de argumentos, contém os argumentos em si              |

### RAX e EAX

São registradores utilizados para dizer ao sistema operacional que é o fim do programa, que estamos finalizando ele, seguindo com o valor de 1:

```asm
mov eax, 0x1
```
### RBX e EBX

São registradores de retorno, indicam ao sistema operacional qual será o retorno, e isso segue a lógica de 0 para sem erros e 1 para com erros:

```asm
mov ebx, 0x0
```

## Instruções

**mov** - Move um valor/endereço de memória para *(geralmente)* outro registrador

```asm
mov eax, 1 ; movendo o valor 1 para o registrador de 32 bits eax
```

> [!NOTE]
> 
> Preferir sempre utilizar número em base hexadecimal, pois alguns registradores, por default, utilizam esta base. E em assembly (como em C e Javascript), dizemos que um valor está em hexadecimal, assim: 0xvalor.
>
> Base Hexadecimal:
>
> 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, A,  B,  C,  D,  E,  F
> 
> A = 10, B = 11, C = 12, D = 13, E = 14, F = 15


## Compilação e Link

No caso estamos utilizando o NASM, então ele será o nosso assembler, mas poderia ser o **as** que é nativo no ubunto + o linker **ld**:

```sh
nasm -f elf64 hello.asm # -f é o arquivo executavel + arquitetura, no caso é um exec linux na arch de 64 bits. hello.asm é o arquivo com o código.
```

Agora para que o sistema operacional entenda que o arquivo compilado será um executável, precisamos linkeditar ele. Linkeditar é fazer com que o nosso objeto binário se junte com as bibliotecas do SO, e aí sim, o SO entende que é um executável e nos permite de fazer iniciar a execução do programa:

```sh
ld -s -o hello hello.o # ld é o linkeditor padrão do linux, -s eu não sei, -o é o objetivo executável que será gerado e hello.o é o nosso arquivo compilado para linguagem de máquina (binário)

./hello # executa o nosso arquivo executável
```

## Usando a seção de constantes

Nesta seção criamos constantes, mas elas não são variáveis, são apontamentos para um endereço de memória:

```asm
section .data
    msg: db 'Hello World!' 
    ; msg é o label, um nome para nós utilizarmos (como se fosse um variavel)
    ; db é o tipo
    ; e o 'Hello World!' é o valor que estará na memória e será buscado quando chamado
```

> obs: os : depois de msg não são obrigatórios, pode ser com e sem.

O assembly tem uma sintaxe que pega o tamanho do que você para ele pegar o tamanho, no caso vamos pegar o tamanho do 'Hello World!', que tem 12 caracterés:

```asm
section .data
    msg db 'Hello World'
    tam equ $- msg 
    ; equ é igual, equivalente, parecido com atribuição com =
    ; $- é a sintaxe que pega o tamanho
    ; msg é o nome da nossa "variavel"
```

Muito cuidado ao colocar mais de um tipo db ou variaveis em geral, em sequência, antes de pegar o tamanho, se não ele vai entender que ainda faz parte da variavel anterior, e, no nosso caso, vai imprimir várias mensagens, já que a sintaxe pega um tamanho variavel:

```asm
section .data

msg0: db 'Hello 1', 0xA ; 0xA indica o line feed (LF), que é a quebra de linha, vem da tabela ASCII
msg1: db 'Hello 2', 0xA
msg2: db 'Hello 3', 0xA
msg3: db 'Hello 4', 0xA
tam: equ $- msg0 ; ele vai entender que até aqui é para ele pegar o tamanho, mesmo passando o msg0 como "argumento"

>>> stdout 
Hello 1
Hello 2
Hello 3
Hello 4
```

O correto seria:

```asm
section .data

msg0: db 'Hello 1', 0xA
tam: equ $- msg0
msg1: db 'Hello 2', 0xA
msg2: db 'Hello 3', 0xA
msg3: db 'Hello 4', 0xA

>>> stdout
Hello 1
```