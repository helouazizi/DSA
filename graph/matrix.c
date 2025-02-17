#include <stdio.h>
#define v 4

void AddEdge(int matrix[v][v],int i,int j){
    matrix[i][j] = 1;
    matrix[j][i] = 1;
    // we work with underected graph so we have no specific deredtion <-->
}

void RmoveVertex(int mat[v][v],int x){
    if (x > v || x < 0 ) {
        printf("vetrex %d not exist \n",x);
        return;
    }
    for (int i = 0 ; i < v ; i++){
        mat[i][x]= 0;
        mat[x][i]=0;
    }

}

// this function for displaying our matrix
void PrintPatrix(int mat[v][v]){
    for (int i = 0;i < v; i++ ){
        for (int j = 0;j < v ; j++){
            printf("%i ",mat[i][j]);
        }
        puts("\n");
    }
}

int main(){
    printf("hi we gonna use graph data structures\n");
    int matrix[v][v] = {0}; // in c we need to intialise matrix before using it because it conteain garbage values
    AddEdge(matrix,0,2);
    AddEdge(matrix,3,2);
    AddEdge(matrix,2,1);
    AddEdge(matrix,0,3);
    PrintPatrix(matrix);
    RmoveVertex(matrix,3);
    RmoveVertex(matrix,0);
    PrintPatrix(matrix);
    puts("this our graph as adjacentcy matrix");
}