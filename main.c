#include<stdio.h>

int main(){
    char txt[] = "ha";
    printf("%s it size: %li\n",txt,sizeof(txt));
    int time = 19;
    (time < 20) ? printf("too young\n"): printf("yes for age\n");
    return 0 ;
};