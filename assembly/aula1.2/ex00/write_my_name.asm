section .data

welcome          db  "Bem vindo!", 0xA, 0xD
welcome_len      equ $ - welcome

enter_first_name db  "Digíte o primeiro nome:", 0xA, 0xD
enter_first_len  equ $ - enter_first_name

your_first_name  db  "Seu nome eh "
your_first_len   equ $ - your_first_name

enter_last_name  db  "Digíte seu sobrenome:", 0xA, 0xD
enter_last_len   equ $ - enter_last_name

your_last_name   db  "Seu sobrenome eh "
your_last_len    equ $ - your_last_name


section    .bss

first_name resb 0x10
last_name  resb 0x10


section    .text

global     _start

_start:

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, welcome
    mov EDX, welcome_len
    int 0x80

save_first_name:

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, enter_first_name
    mov EDX, enter_first_len
    int 0x80

    mov EAX, 0x3
    mov EBX, 0x0
    mov ECX, first_name
    mov EDX, 0x10
    int 0x80

save_last_name:

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, enter_last_name
    mov EDX, enter_last_len
    int 0x80

    mov EAX, 0x3
    mov EBX, 0x0
    mov ECX, last_name
    mov EDX, 0x10
    int 0x80

print_first_name:

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, your_first_name
    mov EDX, your_first_len
    int 0x80

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, first_name
    mov EDX, 0x10
    int 0x80

print_last_name:

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, your_last_name
    mov EDX, your_last_len
    int 0x80

    mov EAX, 0x4
    mov EBX, 0x1
    mov ECX, last_name
    mov EDX, 0x10
    int 0x80

end_program:
    mov EAX, 0x1
    mov EBX, 0x0
    int 0x80