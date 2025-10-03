%include "lib.inc"


section .text

global  _start

_start:
    
    lea ESI,       [BUFFER]
    add ESI,       0x9      ; Anda 9 casas para frente com o cursor
    mov byte[ESI], LF       ; Adiciona uma quebra de linha no último byte de BUFFER

    dec  ESI         ; Se antes ele apontava para o último byte do BUFFER, agora aponta para o penúltimo
    mov, DL,    0x11 ; Move para o registrador DL o valor 17 em decimal
    add  DL,    "0"  ; Soma o caracter 0(0x30) ao 0x11, transformando em string e pegando o ASCII dele, que é o 0x41 - A
    mov  [ESI], DL   ; Move o A (0x41) para o BUFFER

    call returnResult ; Chama o marcador que irá mosrtar o resultado na tela

    mov EAX, SYS_EXIT
    xor EBX, EBX      ; mov EBX, RET_EXIT
    int SYS_CALL
