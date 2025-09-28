segment .data
    
    LF       equ 0xA  ; Line Feed
    NULL     equ 0xD  ; Caractere ASCII para nulo, usado para indicar o final da string
    SYS_CALL equ 0x80 ; Envio informação ao SO
    
    ; Esse vão em EAX ou RAX
    SYS_EXIT  equ 0x1 ; Código de chamada para finalizar
    SYS_READ  equ 0x3 ; Operação de leitura STDIN
    SYS_WRITE equ 0x4 ; Operação de escrita STDOUT

    ; Esse vão em EBX ou RBX
    STD_IN   equ 0x0 ; Entrada padrão
    STD_OUT  equ 0x1 ; Saída padrão
    RET_EXIT equ 0x0 ; Operação realizada com sucesso


section .data
    
    msg db  "Entre com seu nome: ", LF, NULL
    len equ $- msg


section .bss

    name resb 0xA ; é necessário passar a quantidade de espaço que será reservado


section .text

global  _start

_start:

write_to_stdout:

    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, msg
    mov edx, len
    int SYS_CALL

read_from_stdin:

    mov eax, SYS_READ ; Chamada para o kernel para STDIN
    mov ebx, STD_IN   ; File descriptor para o STDIN
    mov ecx, name     ; onde será armazenado o que vier de STDIN
    mov edx, 0xA      ; quantos caracteres ele irá ler (pegar)
    int SYS_CALL

write_name:

    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, name
    mov edx, 0xA       ; Quantos bytes a nossa variável name irá imprimir
    int SYS_CALL

end:
    mov eax, SYS_EXIT
    mov ebx, RET_EXIT
    int SYS_CALL