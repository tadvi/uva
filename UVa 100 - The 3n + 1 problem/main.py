
def calc(n, memo):
    if n in memo:
        return memo[n]

    if n%2 == 0:
        v = calc(n/2, memo)
    else:
        v = calc(3*n+1, memo)
    memo[n] = v+1
    return memo[n]


def main():
    memo = {1: 1}

    with open('input.txt') as f:
        while True:
            s = f.readline()
            if not s:
                break
            x, y = (int(a) for a in s.split())

            mx = 0
            for n in range(x, y+1):
                v = calc(n, memo)
                if mx < v:
                    mx = v

            print(x, y)

if __name__ == '__main__':
    main()
