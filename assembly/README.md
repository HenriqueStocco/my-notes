# Assembly

Parte do repositório destinado ao meu aprendizado da linguagem de baixo nível `Assembly`.

> [!WARNING]
>
> O estudo é focado na arquitetura **Intel x86_64**, em kernel **Linux**.
>
> Sistema Operacional - VM com Debian Server 12.
>
> IDE - Visual Studio Code
>
> Conexão remote via SSH para a minha VM dentro do VSCode
> 
> O compilador que eu utilizo é o **NASM** *The Netwide Assembler* e o **as** *assembler* nativo do linux.
>
> Para o linkeditor, utilizei o **ld**, padrão do linux.
>
> E para a compilação, linkedição e execução automática, utilizei **Makefiles**.

> [!NOTE]
>
> Até o momento, todo o conteúdo aprendido, contido neste repositório, é baseado em uma playlist de vídeos do YouTube, produzido pelo Fernando Anselmo.
> E pesquisas encima do conteúdo, feito por mim, por curiosidade ou erros que ocorreram em meus testes.
>
> Link: https://youtube.com/playlist?list=PLxTkH01AauxRm0LFLlOA9RR5O6hBLqBtC&si=WEu4P8q_I6tInIsk

## Executando

Bom, a partir da aula 1.2, o Fernando passou um Makefile para automatizar a compilação e linkedição do código assembly. Eu modifiquei para limpar o terminal e executar o binário executável, mas, abaixo está o processo manual, utilizando o compilador NASM:

```sh
clear && \
    nasm elf64 <file_name>.asm && \
    ld -s -o <name> <file_name>.o && \
    ./<name>
```

Compilando com **as**, assembler (compilador) nativo do linux:

```sh
clear && \
    as <file_name>.asm -o <file_name>.o && \
    ld -s -o <name> <file_name>.o && \
    ./<name>
```

> Obs: no comando de linkedição com **ld**, a flag -s é responsável por remover qualquer flag adicional de debugging, por exemplo, que o dbg busca nos binários, que o gcc com a flag -g gera.