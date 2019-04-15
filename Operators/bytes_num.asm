section .data

section .text
    global _bytes_num

_bytes_num:
        mov rax, [rsp+4]    ;   argument 1
        mov rcx, 0          ;   счётчик, обнуляем
    l1:
        shr rax, 1          ;   right shift
        jnz l1              ;   if not zero jump to l1
        inc rcx
        ret