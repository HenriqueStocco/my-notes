%include "lib.inc"

section .text

global  _start

_start:
    
    lea ESI,       [BUFFER]
    add ESI,       0x9
    mov BYTE[ESI], LF

    dec ESI
    mov DL,    0x11
    add DL,    "0"
    mov [ESI], DL

    call returnResult

    mov EAX, SYS_EXIT
    mov EBX, RET_EXIT
    int SYS_CALL
