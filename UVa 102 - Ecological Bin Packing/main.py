
from itertools import permutations

def main():
    bins = "BGC"

    with open("input.txt") as f:
        for line in f:
            row = [int(i) for i in line.split()]
            minim = 1 << 30

            for p in permutations([0, 1, 2]):
                total = 0
                for i in range(3):
                    for j in range(3):
                        if j != p[i]:
                            total += row[i * 3 + j]

                name = "".join(bins[i] for i in p)
                if total < minim:
                    curr, minim = name, total
                elif total == minim:
                    curr = min(curr, name)

            print(f"{curr} {minim}")


if __name__ == "__main__":
    main()
