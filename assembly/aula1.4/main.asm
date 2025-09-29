%include "lib.inc"

section .text

global  _start

_start:
    lea esi,       [BUFFER]
    add esi,       0x9
    mov byte[esi], LF

    dec esi
    mov dl,    0x11
    add dl,    '0'
    mov [esi], dl

    call returnResult

    mov eax, SYS_EXIT
    mov ebx, RET_EXIT
    int SYS_CALL
