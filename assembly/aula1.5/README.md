# Aula 1.5 - Calculadora

## Explicação em detalhes

### Conversão de string para integer *(by GPT)*

Código base:

```.asm
lea esi, [value]        ; param 1: valor que quero converter, ex: "123"
mov ecx, 0x3            ; param 2: tamanho do valor a ser convertido
call convert_to_integer 

convert_to_integer:
    xor ebx, ebx
    .next_digit:
        movzx eax, byte [esi]
        inc esi
        sub al, "0"
        imul ebx, 0xA
        add ebx, eax
        loop .next_digit
        mov eax, ebx
        ret
```

- **ESI** *(param 1)* -> aponta para o início da *string* de dígitos *ASCII* ("123")
- **ECX** *(param 2)* -> contém o número de caracteres (no caso 3)
- **EAX** -> retorno convertido (ex: 123)

- `xor EBX, EBX` -> zerando o acumulador
- `movzx` -> copia 1 byte para **EAX** e preenche o resto com zeros
- `byte [ESI]` -> lê o próximo caractere **ASCII**
- `movzx EAX, byte [ESI]` -> lê o primeiro caractere em esi, no caso de "123" ele lê "1" e preenche o restante com zeros, como o caractere 1 na tabela **ASCII**, em hexadecimal, é 0x31, então **EAX** = (**0x00000031** ou **00000031h**)
- `inc ESI` -> incrementa, ESI = ESI + 1, então, ele avança o ponteiro para o próximo caractere que é "2"
- `sub AL, "0"` -> o caractere "0" na **ASCII** é 0x30 em hexa, "1" é 0x31 e assim por diante, logo, 0x31 - 0x30 = 1, resultando em **AL** = (("1"=0x31) - ("0"=0x30) = 1 = 0x01)
- `imul EBX, 0xA` -> multiplica acumulado por 10, ou seja, cada novo dígito desloca o número um "casa decimal" para a esquerda, exemplo: se temos já 12 e agora vem o "3", EBX = 12 então é EBX = EBX * 10 = 120 (decimal).
- `add EBX, EAX` -> soma o valor do digito atual, a tabela explica mais que tudo:

| Iteração | Caractere | Valor **(EAX)** | **EBX** *antes* | **EBX** *após (*10 + **EAX**)* |
| -------- | --------- | --------------- | --------------- | ------------------------------ |
|    1     |    "1"    |        1        |        0        |         0 * 10 + 1 = 1         |
|    2     |    "2"    |        2        |        1        |         1 * 10 + 2 = 2         |
|    3     |    "3"    |        3        |        12       |         12 * 10 + 3 = 123      |
