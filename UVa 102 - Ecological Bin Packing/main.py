
from itertools import permutations

def main():

    bins = 'BGC'

    with open('input.txt') as f:
        while True:
            line = f.readline()
            if not line:
                break

            row = [int(i) for i in line.split()]
            mmin = 1 << 30

            for p in permutations([0, 1, 2]):
                total = 0
                for i in range(3):
                    for j in range(3):
                        if j != p[i]:
                            total += row[i * 3 + j]

                name = ''.join(bins[i] for i in p)
                if total < mmin:
                    curr, mmin = name, total
                elif total == mmin:
                    curr = min(curr, name)

            print(f'{curr} {mmin}')


if __name__ == '__main__':
    main()