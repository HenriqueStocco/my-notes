section .data
    LF       equ 0xA
    NULL     equ 0x0
    SYS_CALL equ 0x80

    SYS_EXIT  equ 0x1
    SYS_READ  equ 0x3
    SYS_WRITE equ 0x4

    STD_IN   equ 0x0
    STD_OUT  equ 0x1
    RET_EXIT equ 0x0

    BUFFER_LEN equ 0x2


section .bss
    BUFFER resb BUFFER_LEN


section .text

returnResult:
    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, BUFFER
    mov edx, BUFFER_LEN
    int SYS_CALL
    ret