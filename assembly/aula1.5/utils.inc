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

    converter_para_inteiro:
        xor ebx, ebx ; ebx é o acumulador, então temos que zera-lo
        .proximo_digito:
            movzx eax, byte[esi]  ; move o caracter atual para eax com extensão de zeros
            cmp   al,  "0"
            jb    .fim_conversao  ; se AL for menor que o caractere ASCII 0x30 (0), não é um numero, vai para .fim_conversao
            cmp   al,  "9"
            ja    .fim_conversao  ; se AL for maior que o caractere ASCII 0x39 (9), não é um numero, vai para .fim_conversao
            inc   esi             ; vai para o prox caracter da string
            sub   al,  "0"        ; convert o digito ASCII para o valor numerico
            imul  ebx, 0xA        ; multiplica o valor atual acumulado por 10
            add   ebx, eax        ; add o valor do digito atual no total acumualado
            loop  .proximo_digito ; decrementa ecx and faz um loop enquanto ecx é diferente de zero
        .fim_conversao:   
            mov eax, ebx ; move o valor inteiro final para eax para retornar 
            ret

    converter_para_string:
        add esi,       0x9
        mov byte[esi], 0xA
        mov ebx,       0xA
        .proximo_digito:
            xor  edx,   edx
            div  ebx
            add  dl,    "0"
            dec  esi
            mov  [esi], dl
            test eax,   eax
            jnz  .proximo_digito
            mov  eax,   esi
            ret

