NAME = write_my_name


run: ${NAME}.o
	clear
	ld -s -o ${NAME} ${NAME}.o
	rm -rf *.o
	clear
	./${NAME}

%.o: %.asm
	nasm -f elf64 $<