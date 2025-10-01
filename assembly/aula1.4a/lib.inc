<<<<<<< HEAD:assembly/aula1.4a/lib.inc
section .data

=======
segment .data
>>>>>>> 06bf423947b09e8b386c8afae13c279178160198:assembly/aula1.4/lib.inc
    LF       equ 0xA
    NULL     equ 0x0
    SYS_CALL equ 0x80

    SYS_EXIT  equ 0x1
    SYS_READ  equ 0x3
    SYS_WRITE equ 0x4

    STD_IN   equ 0x0
    STD_OUT  equ 0x1
    RET_EXIT equ 0x0

    BUFFER_LEN equ 0xA


<<<<<<< HEAD:assembly/aula1.4a/lib.inc
section .bss
    
    BUFFER resb BUFFER_LEN
=======
segment .bss
    BUFFER resb 0xA
>>>>>>> 06bf423947b09e8b386c8afae13c279178160198:assembly/aula1.4/lib.inc


segment .text

returnResult:
    mov EAX, SYS_WRITE
    mov EBX, STD_OUT
    mov ECX, BUFFER
    mov EDX, BUFFER_LEN
    int SYS_CALL
    ret
