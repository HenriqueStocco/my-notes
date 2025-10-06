segment .data
    LF   equ 0xA
    NULL equ 0xD

    SYS_EXIT  equ 0x1
    SYS_READ  equ 0x3
    SYS_WRITE equ 0x4
    SYS_CALL  equ 0x80
    
    STD_IN   equ 0x0
    STD_OUT  equ 0x1
    RET_EXIT equ 0x0

    BUFFER_LEN equ 0xA


segment .bss
    BUFFER resb BUFFER_LEN


segment .text
    returnResult:
        mov EAX, SYS_WRITE  ; RAX - n chamada de sistema - regis. acumulador
        mov EBX, STD_OUT    ; devia ser o RDI - file descriptor - regis. de destino
        mov ECX, BUFFER     ; devia ser o RSI - aponta para o 1fh endereço da string - regis. de origem-
        mov EDX, BUFFER_LEN ; RDX - recebe a quantidade de bytes - regis. de dados
        int SYS_CALL        ; syscall
        ret