    .data
hello:
    .asciz "Hello, World!"  # string to print

    .text
    .global _start

_start:
    mov $4, %eax          # syswrite
    mov $1, %ebx          # stdout
    mov $hello, %ecx      # string addr
    mov $13, %edx         # string len
    int $0x80             # syscall

    mov $1, %eax          # sysexit
    xor %ebx, %ebx        # exit with 0
    int $0x80             # syscall
