

class Blocks:

    def __init__(self, n):
        self.blocks = [[i] for i in range(n)]

    def find(self, val):
        """ Return line index that has value. """
        for i, line in enumerate(self.blocks):
            if val in line:
                return i
        assert False, "not found"
        return -1

    def reset(self, val):
        pos = self.find(val)
        stack = self.blocks[pos]

        while stack[-1] != val:
            self.blocks[stack[-1]].append(self.blocks[pos][-1])
            stack = stack[:-1]

    def process(self, s1, n1, s2, n2):
        if s1 == "move":
            self.reset(n1)

        if s2 == "onto":
            self.reset(n2)

        if s1 == "move":
            pos1, pos2 = self.find(n1), self.find(n2)
            self.blocks[pos1] = self.blocks[pos1][:-1]
            self.blocks[pos2].append(n1)

        elif s1 == "pile":
            pos1, pos2 = self.find(n1), self.find(n2)
            for i, val in enumerate(self.blocks[pos1]):
                if val == n1:
                    blocksToMove = self.blocks[pos1][i:]
                    self.blocks[pos1] = self.blocks[pos1][:i]
                    self.blocks[pos2].extend(blocksToMove)
                    break

    def result(self):
        for i, line in enumerate(self.blocks):
            print(i, ":", " ".join(str(b) for b in line))


def main():
    with open("input.txt") as f:
        n = int(f.readline())
        blocks = Blocks(n)

        for s in f:
            if s == "quit":
                break

            s1, n1, s2, n2 = (a for a in s.split())
            n1, n2 = int(n1), int(n2)
            blocks.process(s1, n1, s2, n2)

        blocks.result()

if __name__ == "__main__":
    main()
