#include <stdio.h>
char buf[16777216];
static unsigned char b64tab[256] = {
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,62,64,64,64,63,
	52,53,54,55,56,57,58,59,60,61,64,64,64,64,64,64,
	64, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,
	15,16,17,18,19,20,21,22,23,24,25,64,64,64,64,64,
	64,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,
	41,42,43,44,45,46,47,48,49,50,51,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,
	64,64,64,64,64,64,64,64,64,64,64,64,64,64,64,64
};
int main() {
	int c;
	while (1) {
		int has = 0, e = 3, cnt = 0;
		unsigned int val = 0;
		while (1) {
			c = getchar();
			if (c == EOF || c == '#') 
				break;
			if (c != '=') {
				if (b64tab[(unsigned char)c] == 64)
					continue;
				val |= b64tab[(unsigned char)c];
			} else
				e--;
			cnt++, has = 1;
			if (cnt == 4) {
				while (e-- > 0) {
					putchar((val >> 16)&0xff);
					val <<= 8;
				}
				cnt = 0, e = 3, val = 0;
			} else {
				val <<= 6;
			}
		}
		if (has)
			printf("#");
		else
			break;
	}
	return 0;
}
/*
VGhpc0lzVGVzdA==
#
QSBUZXN0IElucHV0W3so
KX1d
##
*/
