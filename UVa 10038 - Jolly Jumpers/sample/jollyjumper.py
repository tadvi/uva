'''
Jolly Jumpers
PC/UVa IDs: 110201/10038, Popularity: A, Success rate: average Level: 1
'''


def isJollyJumper(array):
    size = len(array)
    jumps = set()
    if size == 1:
        return True
    for i in range(size - 1):
        diff = abs(array[i + 1] - array[i])
        if diff < size and diff >= 1:
            jumps.add(diff)
    if len(jumps) == size - 1:
        return 'Jolly'
    return 'Not Jolly'

if __name__ == '__main__':
    f = open('input.txt')
    for line in f:
        try:
            array = map(int, line.split())
        except:
            break
        print(isJollyJumper(array[1:]))
