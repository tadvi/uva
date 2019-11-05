
class Cache():

    def __init__(self):
        self.cache = {1: 1}

    def calc(self, n):
        if n in self.cache:
            return self.cache[n]

        if n%2 == 0:
            v = self.calc(n/2)
        else:
            v = self.calc(3*n+1)
        self.cache[n] = v+1
        return self.cache[n]


def main():
    cache = Cache()

    with open("input.txt") as f:
        for line in f:
            x, y = (int(a) for a in line.split())

            maxx = 0
            for n in range(x, y+1):
                v = cache.calc(n)
                if maxx < v:
                    maxx = v

            print(x, y, maxx)

if __name__ == '__main__':
    main()
