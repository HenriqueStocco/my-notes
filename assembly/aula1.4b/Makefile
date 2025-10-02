NAME = main

run: ${NAME}.o
	ld -s -o ${NAME} ${NAME}.o
	rm -rf *.o
	clear
	./${NAME}

%.o: %.asm
	nasm -f elf64 $<
