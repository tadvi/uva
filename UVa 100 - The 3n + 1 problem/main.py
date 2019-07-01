
def calc(n, cache):
    if n in cache:
        return cache[n]

    if n%2 == 0:
        v = calc(n/2, cache)
    else:
        v = calc(3*n+1, cache)
    cache[n] = v+1
    return cache[n]


def main():
    cache = {1: 1}

    with open("input.txt") as f:
        for line in f:
            x, y = (int(a) for a in line.split())

            maxx = 0
            for n in range(x, y+1):
                v = calc(n, cache)
                if maxx < v:
                    maxx = v

            print(x, y, maxx)

if __name__ == '__main__':
    main()
