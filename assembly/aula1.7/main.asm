%include 'lib.inc'

section .data
    par_message:   db  "Par", 0xA, 0x0
    par_len:       equ $- par_message
    impar_message: db  "Impar", 0xA, 0x0
    impar_len:     equ $- impar_message


section .text

global  _start

_start:
    mov edx, 0x0 ; edx = 0x0 -> 0
    mov eax, 0xA ; eax = 0xA -> 10
    mov ebx, 0x2 ; ebx = 0x2 -> 2
    div ebx      ; eax = eax / ebx -> 10 / 2 = 5 resto -> edx = 0

    cmp edx, 0
    je  par
    jmp impar

    mov eax, 0x1
    mov ebx, 0x0
    int 0x80
    
    par:
        mov eax, 0x4
        mov ebx, 0x1
        mov ecx, par_message
        mov edx, par_len
        int 0x80
        ret
    
    impar:
        mov eax, 0x4
        mov ebx, 0x1
        mov ecx, impar_message
        mov edx, impar_len
        int 0x80
        ret
