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

## Usando GDB

No vídeo, o Fernando mostra como o float se diferencia de um int (sendo os dois 32 bits), através do código de máquina, usando o GDB para disassemblar a função main.

```ps1
# Compilar o código C permitindo que ela insira os sinais de depuração
# -g habilita esses sinais de depuração do executavel
gcc -g -o a4 main.c 

# Iniciando o gdb com o executavel e desabilitando os comentario de copyright inicial com -q
gdb -q a4

# disassemblando a função main
disas main

# Se o código já tiver sido executado, precisamos voltar para a primeira linha, usar o comando abaixo antes do disas main
list .
```

O código de máquina que retorna é esse, e eu comentei onde está o valor da variavel e seu endereço

```ps1      
   0x0000000140001450 <+0>:     push   %rbp      
   0x0000000140001451 <+1>:     mov    %rsp,%rbp 
   0x0000000140001454 <+4>:     sub    $0x30,%rsp
   0x0000000140001458 <+8>:     call   0x140001590 <__main>
   0x000000014000145d <+13>:    movss  0x8bdf(%rip),%xmm0        # 0x14000a044 -- endereço da var f
   0x0000000140001465 <+21>:    movss  %xmm0,-0x4(%rbp)
   0x000000014000146a <+26>:    lea    0x8b8f(%rip),%rax
   0x0000000140001471 <+33>:    mov    $0x20,%r8d
   0x0000000140001477 <+39>:    mov    $0x4,%edx
   0x000000014000147c <+44>:    mov    %rax,%rcx
   0x000000014000147f <+47>:    call   0x1400025c0 <__mingw_printf>
   0x0000000140001484 <+52>:    pxor   %xmm0,%xmm0
   0x0000000140001488 <+56>:    cvtss2sd -0x4(%rbp),%xmm0
   0x000000014000148d <+61>:    movapd %xmm0,%xmm1
   0x0000000140001491 <+65>:    movapd %xmm1,%xmm0
   0x0000000140001495 <+69>:    movq   %xmm1,%rdx
   0x000000014000149a <+74>:    lea    0x8b8e(%rip),%rax
   0x00000001400014a1 <+81>:    movapd %xmm0,%xmm1
   0x00000001400014a5 <+85>:    mov    %rax,%rcx
   0x00000001400014a8 <+88>:    call   0x1400025c0 <__mingw_printf>
   0x00000001400014ad <+93>:    mov    $0x0,%eax
   0x00000001400014b2 <+98>:    add    $0x30,%rsp
   0x00000001400014b6 <+102>:   pop    %rbp
   0x00000001400014b7 <+103>:   ret
```

O Fernando diz como ver o valor deste endereço, dessa forma:

```ps1
p/x *0x14000a044

# Meu teste e resultado
(gdb) p/x *0x14000a044
$1 = 0x3f800000 # valor de 1 em número real (na forma hexadecimal)
```

## Dissecando a forma como os 32 bits do número com ponto flutuante é quebrado

> O padrão internacional que representa este formato de números reais (com pontos flutuantes) em C, é a IEEE 754
>
> Link: [IEEE 754](https://en.wikipedia.org/wiki/IEEE_754)

Com o linux é mais fácil de fazer as conversões e calculos com valores de outras bases, devido a sua calculadora nativa, que é o `bc`.

```sh
# Usando o bc para coverter o valor hexadecimal, que é o 1 com ponto flutuante que vemos no código de máquina
# obase=2; é o formato de saída que queremos, no caso base 2 (binário)
# ibase=16; é o formato de entrada, no caso é base 16 (hexadecimal)
# echo para mostrar na saida padrão o valor convertido
# um pipe (|) para bc, para que ele use o que está dentro da string como parâmetros do comando bc e já devolva os valores no echo
echo 'obase=2; ibase=16; 3F800000' | bc

# Retorno
111111100000000000000000000000

# bc não retorna valores a esquerda, mas realmente, o valor em binário é:
00111111100000000000000000000000

# Se a variável fosse um int32, esse seria o valor (no caso o mesmo valor que utilizamos)
00000000000000000000000000000001

# Dos 32 bits, 23 são reservados para a mantíssia, que é a parte fracionária, o número a ser multiplicado pela base ao expoente.
# No meu caso, ele seria:
00000000000000000000000 # mantissia

# São 8 bits reservados para o expoente
01111111

# E o ultimo bit é o que determina o sinal, como no inteiro
0

# Resumindo
  0      01111111  00000000000000000000000
# SINAL  EXPOENTE  MANTISSIA
# 1 bit  8 bits    23 bits

# Estes 8 bits na verdade são subtraidos ao valor 127 que aí sim vai resultar no expoente na base 2.
# A fórmula, fornecida pelo mestre GPT, seria: valor = (-1)^S * 1.M * 2^(E - 127)
# Onde: S = SINAL (1 NEGATIVO ou 0 POSITIVO)
#       M = MANTISSA (fracionário)
#       E = EXPOENTE (com bias ou viés de 127)
```