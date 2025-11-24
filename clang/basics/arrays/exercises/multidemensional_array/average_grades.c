/*
    Vamos tentar descobrir as notas médias de um grupo de cinco alunos em duas disciplinas, Matemática e Física. Para isso, usamos uma matriz bidimensional chamada notas. As notas correspondentes à Matemática seriam armazenadas na primeira linha (notas[0]), enquanto as correspondentes à Física seriam armazenadas na segunda linha (notas[1]). Complete as etapas a seguir para que você possa executar este programa.

    - Declare notas como uma matriz bidimensional de inteiros.
    - Complete os loops for especificando suas condições de término.
    - Calcule as notas médias obtidas em cada disciplina.
 */

#include <stdio.h>

int main(void)
{
    float media;
    int i, j, notas[2][5];

    // Notas de Matemática
    notas[0][0] = 80;
    notas[0][1] = 70;
    notas[0][2] = 65;
    notas[0][3] = 89;
    notas[0][4] = 90;

    // Notas de Física
    notas[1][0] = 85;
    notas[1][1] = 80;
    notas[1][2] = 85;
    notas[1][3] = 82;
    notas[1][4] = 87;

    /* TODO: completar o loop for com condições de terminação adequadas */
    for (i = 0; i < 2; i++)
    {
        media = 0;
        for (j = 0; j < 5; j++)
        {
            media += notas[i][j];
        }

        media /= 5.0;
        
        /* TODO: calcular as notas médias para a disciplina i */
        printf("A nota media obtida na disciplina %d e: %.2f\n", i, media);
    }
    return 0;
}