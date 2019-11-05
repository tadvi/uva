

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

    def process(self, action, src, directive, dest):
        if action == "move":
            self.reset(src)

        if directive == "onto":
            self.reset(dest)

        if action == "move":
            from_, to = self.find(src), self.find(dest)
            self.blocks[from_] = self.blocks[from_][:-1]
            self.blocks[to].append(src)

        elif action == "pile":
            from_, to = self.find(src), self.find(dest)
            for i, val in enumerate(self.blocks[from_]):
                if val == src:
                    blocksToMove = self.blocks[from_][i:]
                    self.blocks[from_] = self.blocks[from_][:i]
                    self.blocks[to].extend(blocksToMove)
                    break

    def result(self):
        for i, line in enumerate(self.blocks):
            print(i, ":", " ".join(str(b) for b in line))


def main():
    with open("input.txt") as f:
        n = int(f.readline())
        blocks = Blocks(n)

        for line in f:
            if line == "quit":
                break

            action, from_, directive, to = (a for a in line.split())
            from_, to = int(from_), int(to)
            blocks.process(action, from_, directive, to)

        blocks.result()

if __name__ == "__main__":
    main()
