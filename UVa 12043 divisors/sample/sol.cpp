#include <cstdio>
#include <cstring>
using namespace std;
int factor[100001], divisors[100001];
long long sum[100001];
int main()
{
    memset(factor, -1, sizeof factor);
    divisors[1] = sum[1] = 1;
    for (int i = 2; i <= 100000; ++i)
    {
        if (factor[i] == -1)
        {
            factor[i] = i;
            if (i <= 100000 / i)
                for (int j = i * i; j <= 100000; j += i)
                    factor[j] = i;
        }
        int p = factor[i], aux = i;
        divisors[i] = sum[i] = 1;
        while (aux % p == 0)
        {
            ++divisors[i];
            sum[i] *= p;
            aux /= p;
        }
        divisors[i] *= divisors[aux];
        sum[i] = (sum[i] * p - 1) / (p - 1) * sum[aux];
    }
    int T, a, b, k;
    scanf("%d", &T);
    while (T--)
    {
        scanf("%d %d %d", &a, &b, &k);
        long long g = 0, h = 0;
        for (int i = (a + k - 1) / k * k; i <= b; i += k)
        {
            g += divisors[i];
            h += sum[i];
        }
        printf("%lld %lld\n", g, h);
    }
    return 0;
}