segment .data

    LF       equ 0xA  ; Line feed
    NULL     equ 0xD  ; Null
    SYS_CALL equ 0x80 ; Envia informão ao Sistem Operacional

    ; EAX ou RAX
    SYS_EXIT  equ 0x1 ; Código de chamada para finalizar
    SYS_READ  equ 0x3 ;Operação de leitura em STDIN
    SYS_WRITE equ 0x4 ; Operação de escrita em STDOUT

    ; EBX ou RBX
    FD_READ  equ 0x0 ; Descritor de arquivo para a entrada padrão
    FD_OUT   equ 0x1 ; Descritor de arquivo para a saída padrão
    RET_EXIT equ 0x0 ; Operação realizada com sucesso


section .data

    x dd 10 ; double word = 4 bytes
    y dd 20 ; double word = 4 bytes

    msg1     db  "X maior que Y", LF, NULL
    msg1_len equ $ - msg1
    
    msg2     db  "Y maior que X", LF, NULL
    msg2_len equ $ - msg2
    

section .text


global  _start

_start:

    mov EAX, DWORD[x] ; conversão relativa
    mov EBX, DWORD[y] ; conversão relativa
    ; if - utiliza o operador CMP
    cmp EAX, EBX      ; compara o que estiver em EAX (x = 50) com o que estiver em EBX (y = 10)
    jge greater       ; se x maior que y, executa o que estiver na label greater
    mov ECX, msg2     ; se não, vai imprimir a mensagem 2
    mov EDX, msg2_len
    jmp continue      ; então, vai para a continução (a label continue), e aí no caso ele vai finalizar o programa, imprimindo a mensagem de acordo com a condição.
    
greater:
    mov ECX, msg1
    mov EDX, msg1_len

continue:
    mov EAX, SYS_WRITE
    mov EBX, FD_OUT
    int SYS_CALL

; end_program:
    mov EAX, SYS_EXIT
    xor EBX, EBX      ; mov EBX, RET_EXIT
    int SYS_CALL