# Aula 1.4 - Biblioteca externa & Função

## include

A clausula `%include`, permite que o código de outro arquivo seja adicionado ao executável final, após a linkedição do objeto compilado.

## call & ret

O comando `ret` serve para desviar o fluxo de volta para a próxima instrução onde o marcador que "agrupa" as instruções (que fica parecendo uma função, com escopo e tudo), volte para onde foi chamado (com o call).

O comando `call` faz a chamada do marcador e permite usar o `ret` para continuar na próxima linha onde o `call` foi chamado, dando sequência no fluxo principal.

## Instruções

### lea

A instrução `lea` *(Load Effective Address)*, carrega o endereço de memória calculado de uma variavel, no nosso caso, da variavel **BUFFER** e salva no registrador `ESI`.

O endereço calculado é o endereço efetivo de uma variavel, isso significa que a instrução `lea` não acessa a memória da variavel, e sim, apenas pega o endereço dela e salva no registrador.

Se usassemos o `mov`, estariamos pegando o valor contido naquele endereço, devido os colchetes, que são basicamente os ponteiros, eles indicam que aquele é o endereço de memória da variavel, e o `mov` faria o acesso a este endereço e pegaria o valor. Mesmo retirando os colchetes, ele teria o endereço de memória da variavel mas teria que acess-la para obte-lo. Essa é a diferença entre a instrução `mov` e `lea`, o `mov` precisa acessar a memória para obter tanto o endereço quanto o valor da variavel, o `lea` calcula o endereço, sem acessar a memória e transfere este endereço a um registrador.

```asm
lea ESI, [BUFFER] ; calcula o endereço de memória da variavel BUFFER
mov ESI, [BUFFER] ; acessa a memória e busca pelo valor dentro do endereço de BUFFER, movendo o valor para o registrador ESI
mov ESI, BUFFER ; acessa a memória e pega o endereço de memória da variavel BUFFER, movemendo este endereço para o registrador ESI

```

### add

Essa instrução serve para **somar** dois valores, mas no caso, está sendo utilizado para somar o endereço do registrador ESI com 9, fazendo o ponteiro andar 9 casas para a frente.
