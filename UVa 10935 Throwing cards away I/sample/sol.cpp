#include<cstdio>
#include<queue>
using namespace std;

int main()
{
    int n;
    while (scanf("%d", &n) == 1 && n > 0)
    {
        queue<int> q;
        for (int i = 0; i < n; i++) q.push(i + 1);
        int idx = 1;
        printf("Discarded cards:");
        for (int i = 0; i < n - 1; i++)
        {
            if (i == 0) printf(" "); else printf(", ");
            printf("%d", q.front());
            q.pop();
            q.push(q.front());
            q.pop();
        }
        printf("\nRemaining card: %d\n", q.front());
    }
    return 0;
}
