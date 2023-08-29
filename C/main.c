#include <stdio.h>
#include <stdlib.h>
#include <math.h>

struct Analytical_Set{
double Dx;
double Dy;
double COVxy;
double PCC;
double A;
double B;
};

const char LRL[] = "Linear Regression Line: ";
const char PCC[] = "Pearson Correlation Coefficient: ";

int main(int argc, char *argv[]){
if (argc != 2){
    perror("Incorrect input arg");
    return 1;
}
char* fileName = argv[1];
    //reading file
    struct Analytical_Set Set;
    FILE *input = fopen(fileName,"r");
    if (input == NULL){
        perror("File not finde");
        return 1;
    }
    double sumX, sumY, sumXY, sumXPow2, sumYPow2, n; 
    sumX = sumY = sumXY = sumXPow2 = sumYPow2 = n = 0;
    char buffio[100];
    while(fgets(buffio,sizeof(buffio),input)!= NULL){
        double y = atof(buffio);
        double x = n; 
        n++;
        sumX += x;
        sumY += y;
        sumXY += (x * y);
        sumXPow2 += pow(x, 2);
        sumYPow2 += pow(y, 2);
    }
    fclose(input);

    Set.Dx = sumXPow2 - pow(sumX, 2)/n;
    Set.Dy = sumYPow2 - pow(sumY, 2)/n;
    Set.COVxy = sumXY - sumX*sumY/n;
    Set.A = Set.COVxy / Set.Dx;
    Set.B = (sumY - Set.A*sumX) / n;
    Set.PCC = Set.COVxy / sqrt(Set.Dx*Set.Dy);
    printf("%sy = %.6fx + %.6f\n",LRL, Set.A, Set.B);
    printf("%s%.10f\n",PCC,Set.PCC);
    return 0;
}
