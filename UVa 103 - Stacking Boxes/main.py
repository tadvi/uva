
def main():
    with open("input.txt") as f:
        while line := f.readline():
            n, d = [int(i) for i in line.split()]


if __name__ == "__main__":
    main()