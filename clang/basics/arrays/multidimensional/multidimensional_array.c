/*
    Arrays de 2 dimens�es
    obs: A logica das dimens�es serve para mais de 2 dimensões, ex: int a[2][2][2]...[n]

    São arrays que simulam linhas e colunas.

    Sintaxe de um array com 2 dimens�es:
    int array[0] - A primeira dimensão é a linha
    int array[0][1] - A segunda dimensão é a coluna

    Visualizando:
    int a[0][0];
                    | column 0  | column 1 | column 2 |
            row 0   | a[0][0]   | a[0][1]  | a[0][2]  |
            row 1   | a[1][0]   | a[1][1]  | a[1][2]  |
            row 2   | a[2][0]   | a[2][1]  | a[2][2]  |
*/
#include <stdio.h>

int main(void)
{
    // Declarando um array multidimensional.
    // int array[4][4];
    // Inicializando um array multidimensional
    int array[4][4] = {
        {
            1, 2, 3, 4 // row: 0, col: 0-3
        },
        {
            1, 2, 3, 4 // row: 1, col: 0-3
        },
        {
            1, 2, 3, 4 // row: 2, col: 0-3
        },
        {
            1, 2, 3, 4 // row: 3, col: 0-3
        }};
    // Declarar assim é a mesma coisa que acima: int array[4][4] = {1,2,3,4,1,2,3,4,1,2,3,4,1,2,3,4};

    // Acessando um array multidimensional
    printf("ROW 2 & COL 3: %d\n", array[2][2]);

    return 0;
}
