section .data
    message:          db  "Entre com um valor de 3 digitos: ", 0xA, 0x0
    message_len:      equ $ - message
    even_message:     db  "Número par!", 0xA, 0x0
    even_message_len: equ $ - even_message
    odd_message:      db  "Número ímpar!", 0xA, 0x0
    odd_message_len:  equ $ - odd_message


section .bss
    number: resb 0x1


section .text
    global _start
    
    _start:
        mov eax, 0x4
        mov ebx, 0x1
        mov ecx, message
        mov edx, message_len
        int 0x80

        call enter_value

        lea  esi, [number]
        mov  ecx, 0x3
        call string_to_int
        
        call check

        even:
            mov eax, 0x4
            mov ebx, 0x1
            mov ecx, even_message
            mov edx, even_message_len
            int 0x80
            jmp end
        
        odd:
            mov eax, 0x4
            mov ebx, 0x1
            mov ecx, odd_message
            mov edx, odd_message_len
            int 0x80
            jmp end

        check:
            xor edx, edx
            mov ebx, 0x2
            div ebx
            cmp edx, 0x0
            jz  even
            jmp odd
        
        enter_value:
            mov eax, 0x3
            mov ebx, 0x0
            mov ecx, number
            mov edx, 0x4
            int 0x80
            ret

        string_to_int:
            xor ebx, ebx
            .next_digit:
                movzx eax, byte[esi]
                inc   esi
                sub   al,  '0'
                imul  ebx, 0xA
                add   ebx, eax
                loop  .next_digit
                mov   eax, ebx
                ret

        int_to_string:
            mov esi,        [number]
            add esi,        0x9
            mov byte [esi], 0xA
            mov ebx,        0xA
            .next_digit:
                xor  edx,   edx
                div  ebx
                add  dl,    '0'
                dec  esi
                mov  [esi], dl
                test eax,   eax
                jnz  .next_digit
                ret
        
        end:
            mov eax, 0x1
            mov ebx, 0x0
            int 0x80