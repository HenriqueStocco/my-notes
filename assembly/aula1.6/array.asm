section .data
    ; msg1:     db  "Parte 1", 0xA, 0x0
    ; msg1_len: equ $- msg1
    ; msg2:     db  "Parte 2", 0xA, 0x0
    ; msg3:     db  "Parte 3", 0xA, 0x0
    ; msg4:     db  "Parte 4", 0xA, 0x0
    array: dd 10, 22, 13, 15, 55


section .bss
    BUFFER: resb 0x4


section .text

global  _start

_start:
    ; mov eax, 0x4
    ; mov ebx, 0x1
    ; mov ecx, msg1
    ; mov edx, 0x4
    ; int 0x80

    mov  eax, [array + 4 * 1]
    call to_string
    call return
    
    mov eax, 0x1
    mov ebx, 0x0
    int 0x80

    return:
        mov eax, 0x4
        mov ebx, 0x1
        mov ecx, BUFFER
        mov edx, 0xA
        int 0x80
        ret

    to_string:
        lea esi,       [BUFFER]
        add esi,       0x9
        mov byte[esi], 0xA
        mov ebx,       0xA
        .next:
            xor  edx,   edx
            div  ebx
            add  dl,    "0"
            dec  esi
            mov  [esi], dl
            test eax,   eax
            jnz  .next
            ret