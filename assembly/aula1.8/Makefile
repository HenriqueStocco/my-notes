NOME = main 

all: $(NOME).cpp $(NOME).o
	g++ $(NOME).o $(NOME).cpp -o $(NOME)
	rm -rf $(NOME).o
	clear
	./$(NOME)

%.o: %.asm
	nasm -f elf64 $<
