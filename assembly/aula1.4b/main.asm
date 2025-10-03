%include "lib.inc"

section .data
    v1 dw "105", LF, NULL


section .text

global  _start

_start:
    call convert_value
    call show_value

    mov EAX, SYS_EXIT
    mov EBX, RET_EXIT
    int SYS_CALL

convert_value:
    lea  ESI, [v1]
    mov  ECX, 0x3
    call string_to_int
    add  EAX, 0x2
    ret

show_value:
    call int_to_string
    call returnResult
    ret

string_to_int:
    xor EBX, EBX
.next_digit:
    movzx EAX, byte[ESI]
    inc   ESI
    sub   AL,  "0"
    imul  EBX, 0xA       ; EBX (0) = EBX * 10
    add   EBX, EAX       ; EBX = EBX + EAX
    loop  .next_digit    ; while(--ECX)
    mov   EAX, EBX
    ret

int_to_string:
    lea ESI,       [BUFFER]
    add ESI,       0x9
    mov byte[ESI], 0xA
    mov EBX,       0xA
.next_digit:
    xor  EDX,   EDX
    div  EBX
    add  DL,    "0"
    dec  ESI
    mov  [ESI], DL
    test EAX,   EAX
    jnz  .next_digit
    ret