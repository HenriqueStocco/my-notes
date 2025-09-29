NAME = opa

start: ${NAME}.o
	clear
	ld -s -0 ${NAME} ${NAME}.o
	rm -rf *.o
	clear
	./${NAME}

%.o: %.asm:
	nasm -f elf64 $<