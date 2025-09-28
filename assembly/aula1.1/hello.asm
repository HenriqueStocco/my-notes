section .data ; seção de "constantes"
    msg: db  'Hello World!', 0xA
    tam: equ $- msg

section .text

global  _start

_start:
    
    mov eax, 0x4 ; indica ao SO que eu irei mandar algo para a saida padrão (STDOUT), isso por que estamos usando a syscall 0x4
    mov ebx, 0x1 ; indica também que vai ser utilizado a saída padrão, os registradores EAX e EBX tem paridades, andam quase sempre juntos
    mov ecx, msg
    mov edx, tam ; aqui influencia em quantos caracteres (ou endereços de memória/indices) ele vai buscar referente ao valor que está em ECX
    int 0x80

    ; finalizando o programa 

    mov eax, 0x1 ; indica ao SO que está finalizando o programa
    mov ebx, 0x0 ; indicad ao SO o retorno é 0
    int 0x80
    ; int significa interrupção, e o 0x80 é o número de interrupção
    ; o 0x80 indica ao SO que ele deve executar todos as movimentações que eu fiz acima
