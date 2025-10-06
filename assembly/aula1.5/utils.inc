segment .data
    LF   equ 0xA ; Line feed
    NULL equ 0x0 ; Null

    SYS_READ  equ 0x3  ; Chamada OS de leitura
    SYS_WRITE equ 0x4  ; Chamada OS de escrita
    SYS_CALL  equ 0x80 ; Chamada OS de execucao
    SYS_EXIT  equ 0x1  ; Chamada OS de finalizacao

    STD_IN   equ 0x0 ; Onde será feita a leitura
    STD_OUT  equ 0x1 ; Onde será feita a escrita
    RET_EXIT equ 0x0 ; Onde será o termino da execucao


segment .text
    mostrar_saida:
        call tamanho_string
        mov  eax, SYS_WRITE
        mov  ebx, STD_OUT
        int  SYS_CALL
        ret
    
    tamanho_string:
        mov edx, ecx
    
    proximo_char:
        cmp byte[edx], NULL
        jz  terminei
        inc edx
        jmp proximo_char

    terminei:
        sub edx, ecx
        ret