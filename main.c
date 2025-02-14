#include <stdio.h>

int sum(int h);

int main()
{
    int x = sum(10);
    printf("%d\n",x);
};



int sum(int h)
{
    if (h > 0)
    {
        return h+ sum(h - 1);
    }
    else
    {
        return 0;
    }
}
