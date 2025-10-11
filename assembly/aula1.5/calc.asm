%include "utils.inc"


section .data
    titulo             db LF, "+-------------+", LF, "| Calculadora |", LF, "+-------------+", LF, NULL
    obter_valor_1      db LF, "Valor 1: ", NULL
    obter_valor_2      db LF, "Valor 2: ", NULL
    opcao_1            db LF, "1. Adicionar", NULL
    opcao_2            db LF, "2. Subtrair", NULL
    opcao_3            db LF, "3. Multiplicar", NULL
    opcao_4            db LF, "4. Dividir", NULL
    mensagem_opcao     db LF, "Deseja realizar? ", NULL
    mensagem_erro      db LF, "Valor de Opção Inválido", NULL
    processo_1         db LF, "Processo de Adicionar", LF, NULL
    processo_2         db LF, "Processo de Subtrair", NULL
    processo_3         db LF, "Processo de Multiplicar", NULL
    processo_4         db LF, "Processo de Dividir", NULL
    mensagem_fim       db LF, "Terminei", LF, NULL
    mensagem_resultado db LF, "Resultado: ", NULL
    

section .bss
    opcao     resb 0x1
    numero_1  resb 0x4
    numero_2  resb 0x4
    resultado resb 0x4


section .text

global  _start

_start:
    ; TITULO
    mov  ecx, titulo
    call mostrar_saida

    ; OPÇÕES
    mov  ecx, opcao_1
    call mostrar_saida
    mov  ecx, opcao_2
    call mostrar_saida
    mov  ecx, opcao_3
    call mostrar_saida
    mov  ecx, opcao_4
    call mostrar_saida

    ; OBTER OPÇÃO
    mov  ecx, mensagem_opcao
    call mostrar_saida
    mov  eax, SYS_READ
    mov  ebx, STD_IN
    mov  ecx, opcao
    mov  edx, 0x2
    int  SYS_CALL

    ; CONVERTE A OPÇÃO EM NUMERO
    mov ah, [opcao]
    sub ah, "0"

    ; VERIFICA SE EH CORRETO
    cmp ah, 4
    jg  mostrar_erro
    cmp ah, 1
    jl  mostrar_erro

    ; OBTER VALORES
    mov  ecx, obter_valor_1
    call mostrar_saida
    mov  eax, SYS_READ
    mov  ebx, STD_IN
    mov  ecx, numero_1
    mov  edx, 0x4
    int  SYS_CALL

    mov  ecx, obter_valor_2
    call mostrar_saida
    mov  eax, SYS_READ
    mov  ebx, STD_IN
    mov  ecx, numero_2
    mov  edx, 0x4
    int  SYS_CALL

    ; CONVERTE A OPÇÃO EM NUMERO
    mov ah, [opcao]
    sub ah, "0"

    ; CONDICAO DE ACORDO COM A OPÇAO
    cmp ah, 1
    je  adicionar
    cmp ah, 2
    je  subtrair
    cmp ah, 3
    je  multiplicar
    cmp ah, 4
    je  dividir
    

; MARCADORES DE OPERACAO
    adicionar:
        mov  ecx, processo_1
        call mostrar_saida

        lea  esi, [numero_1]
        mov  ecx, 0x4
        call converter_para_inteiro
        mov  edi, eax

        lea  esi, [numero_2]
        mov  ecx, 0x4
        call converter_para_inteiro

        add eax, edi

        lea  esi, [resultado]
        call converter_para_string
        mov  ebp, eax

        mov  ecx, mensagem_resultado
        call mostrar_saida

        mov  ecx, ebp
        call mostrar_saida

        jmp saida

    subtrair:
        mov  ecx, processo_2
        call mostrar_saida

        lea  esi, [numero_1]
        mov  ecx, 0x1
        call converter_para_inteiro
        mov  edi, eax

        lea  esi, [numero_2]
        mov  ecx, 0x1
        call converter_para_inteiro
        
        sub edi, eax
        xor eax, eax
        mov eax, edi

        lea  esi, [resultado]
        call converter_para_string
        mov  ebp, eax

        mov  ecx, mensagem_resultado
        call mostrar_saida

        mov  ecx, ebp
        call mostrar_saida

        jmp saida

    multiplicar:
        mov  ecx, processo_3
        call mostrar_saida

        lea  esi, [numero_1]
        mov  ecx, 0x4
        call converter_para_inteiro
        mov  edi, eax

        lea  esi, [numero_2]
        mov  ecx, 0x4
        call converter_para_inteiro
        mov  ebp, eax
        mov  eax, edi
        
        mul ebp
        xor edi, edi

        lea  esi, [resultado]
        call converter_para_string
        mov  edi, eax

        mov  ecx, mensagem_resultado
        call mostrar_saida

        mov  ecx, edi
        call mostrar_saida

        jmp saida

    dividir: 
        mov  ecx, processo_4
        call mostrar_saida

        lea  esi, [numero_2]
        mov  ecx, 0x4
        call converter_para_inteiro
        mov  edi, eax

        lea  esi, [numero_1]
        mov  ecx, 0x4
        call converter_para_inteiro

        div edi
        xor edi, edi

        lea  esi, [resultado]
        call converter_para_string
        mov  edi, eax

        mov  ecx, mensagem_resultado
        call mostrar_saida

        mov  ecx, edi
        call mostrar_saida

        jmp saida

    mostrar_erro:
        mov  ecx, mensagem_erro
        call mostrar_saida
        jmp  saida

; TERMINO DO PROGRAMA
    saida:
        mov  ecx, mensagem_fim
        call mostrar_saida
        mov  eax, SYS_EXIT
        xor  ebx, ebx
        int  SYS_CALL