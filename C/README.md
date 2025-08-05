# Meu aprendizado da linguagem C

Aqui é todo o conteúdo e anotações que estudei ou estou estudando sobre a linguagem C.

## Compilação & Execução

### Linux (bash | fish | zsh)

```bash
#Compilar mostrando todos os erros possiveis para melhorar a escrita do código
gcc -o <obj_name> -Wall -Werror --pedantic-errors <filename.c>

# gcc - chamada do compilador (GNU C Compiler)

# -o - nome do binário (objeto/executável) compilado e linkado para execução do programa C

# -Wall - Ativa/mostra todas as mensagens de aviso(warnings all) padrão, ajuda a evitar más práticas e erros comuns

# -Werror - Transforma todos os avisos (warnings) em erros, impedindo a compilação do código até que todos os erros sejam solucionados.

# --pedantic-errors - Impõe estritamente o padrão ISO C: qualquer violação de conformidade com o padrão será um erro, não apenas um aviso.

# Compilando e executando
clear && gcc -o <obj_name> -Wall -Werror --pedantic-errors <filename.c> && ./<obj_name>
```

### Windows (prompt | Powershell)


```ps1
#Compilar mostrando todos os erros possiveis para melhorar a escrita do código
gcc -o <obj_name> -Wall -Werror --pedantic-errors <filename.c>

# gcc - chamada do compilador (GNU C Compiler)

# -o - nome do binário (objeto/executável) compilado e linkado para execução do programa C

# -Wall - Ativa/mostra todas as mensagens de aviso(warnings all) padrão, ajuda a evitar más práticas e erros comuns

# -Werror - Transforma todos os avisos (warnings) em erros, impedindo a compilação do código até que todos os erros sejam solucionados.

# --pedantic-errors - Impõe estritamente o padrão ISO C: qualquer violação de conformidade com o padrão será um erro, não apenas um aviso.

# Compilando e executando
clear ; gcc -o <obj_name> -Wall -Werror --pedantic-errors <filename.c> ; ./<obj_name>
```