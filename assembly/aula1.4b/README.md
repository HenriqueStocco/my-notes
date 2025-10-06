# Aula 1.4 b - Conversores

## movzx

Mover dados de um operando para outro, enquanto converte um valor menor em um valor maior, preenchendo a parte superior com zeros. o **zx** significa *zero extend*.

## sub

Subtrai um operando de outro, armazenando o resultado no primeiro operando, exemplo:

```asm
;   destino,   destino
mov EAX,       10 ; EAX = 10
sub EAX,        4  ; EAX = EAX - 4 (10 - 4 = 6)
```

Mas não somente isso, este tipo de operação altera algumas flags do registrador EFLAGS *(ou RFLAGS se for 64 bits)*, essas são:

- zero flag **(ZF)** - se o resultado é zero
- sign flag **(SF)** - se o resultado é negativo
- carry flag **(CF)** -  se ocorreu um "empréstimo" durante a subtração, ou seja, quando o operando de destino é menor que o operando de origem.
- overflow flag **(OF)** -  se houve uma condição de estouro, que ocorre em operações com números que excedem o espaço disponível (em operações com sinal).

## inc

Como o nome sugere, ela incrementa um valor em 1.

## dec 

Como o nome sugere, ela decrementa um valor em 1.

## imul

Realiza a multiplicação.

## loop

Cria um loop imediato, decrementando o registrador `ECX` até que ele seja igual a 0.


## div

Como o nome sugere, divide o valor do registrador `EAX` pelo operando fornecido, e o resultado é armazenado em `EAX` e o resto em `EDX`, exemplo: 

```asm
mov EAX, 20
mov EBX, 4
div EBX, ; EAX = EAX / EBX, EDX = (EAX % EBX)
```

## test

Esta instrução realiza uma operação **AND** entre dois operandos, mas não armazena o resultado; apenas atualiza as flags, exemplo:

```asm
mov EAX, 3
test EAX, 1 ; Verifica se o bit menos significativo é 1
; A zero flag (ZF) não será setada porque EAX não é 0
```

## jnz

Salto condicional *(jump if not equal)*. Se a **ZF** não estiver setada *(se o resultado anterior não foi 0)*, a execução salta para o endereço especificado *(marcador)*.