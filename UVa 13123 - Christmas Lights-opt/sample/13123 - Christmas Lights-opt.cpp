#include <bits/stdc++.h>
using namespace std;

#define MAXN 1048576
#define BLKSZ 32
#define GET(x) ((A[x>>6]>>(x&63))&1)
#define FLIP(x) (A[x>>6]^=(1ULL<<(x&63)))

struct SQRT {
	struct Node {
		int inv, has0, has1;
		inline int has() {
			return inv ? has0 : has1;
		}
	} block[MAXN>>6];
	int block_n, block_m;
	uint64_t A[MAXN>>6];
	void to_vector(char s[], int n) {
		int sn = strlen(s);
		{
			int m = n/64 + (n%64 != 0);
			memset(A, 0, sizeof(A[0])*(m/BLKSZ+1)*BLKSZ);
		}
		int m = 0;
		for (int i = sn-1; i >= 0; ) {
			uint64_t mask = 0;
			for (int j = 0; j < 16 && i >= 0; j++, i--) {
				uint64_t v = s[i] <= '9' ? s[i]-'0' : s[i]-'A'+10;
				mask = mask | (v<<(j*4));
			}
//			puts("??????");
			A[m++] = mask;
		}
		{
			m = n/64 + (n%64 != 0);
			block_m = m;
			block_n = m/BLKSZ + (m%BLKSZ != 0);
			for (int i = 0; i < m; i += BLKSZ) {
				int bl = i;
				int br = min(m, bl+BLKSZ);
				int bx = bl/BLKSZ;
				int has0 = 0, has1 = 0;
				for (int j = bl; j < br; j++)
					has1 |= A[i], has0 |= ~A[i];
				block[bx].inv = 0;
				block[bx].has0 = has0;
				block[bx].has1 = has1;
			}
		}
	}
	void update_block(int x, int n) {
		int cx = x>>6;
		int bx = cx/BLKSZ;
		int bl = bx*BLKSZ;
		int br = min(n, bl+BLKSZ);
		if (block[bx].inv) {
			for (int i = bl; i < br; i++)
				A[i] = ~A[i];
			block[bx].inv = 0;
		}
		int has0 = 0, has1 = 0;
		for (int i = bl; i < br; i++)
			has1 |= A[i], has0 |= ~A[i];
		block[bx].has0 = has0;
		block[bx].has1 = has1;
	}
	inline int get(int x, int n) {
		if (x >= n || x < 0)
			return 0;
		int cx = x>>6;
		int bx = cx/BLKSZ;
		int t = ((A[cx]>>(x&63))&1);
		return block[bx].inv ? !t : t;
	}
	inline void flip(int x) {
		A[x>>6] ^= (1ULL<<(x&63));
	}
	inline uint64_t get_inblock(int cx) {
		int bx = cx/BLKSZ;
		return block[bx].inv ? ~A[cx] : A[cx];
	}
	inline int ctzll(int cx) {
		int bx = cx/BLKSZ;
		uint64_t a = A[cx];
		if (block[bx].inv)
			a = ~a;
		return __builtin_ctzll(a);
	}
	inline int clzll(int cx) {
		int bx = cx/BLKSZ;
		uint64_t a = A[cx];
		if (block[bx].inv)
			a = ~a;
		return __builtin_clzll(a);
	}
	int prev_one(int x, int n) { // [0, x]
		if (x < 0 || x >= n)	return -1;
		update_block(x, n);
		/* find in a 64-bit integer */
		while (x >= 0 && (x&63) != 63) {
			if (get(x, n) == 1)	return x;
			x--;
		}
		if (x < 0)	return -1;
		/* find in a block */
		int cx = x>>6;
		while (cx >= 0 && (cx%BLKSZ) != BLKSZ-1) {
			if (get_inblock(cx)) {
				return (cx<<6) + 63-clzll(cx);
			}
			cx--;
		}
		if (cx < 0)	return -1;
		/* find in a block chain */
		int bx = cx/BLKSZ;
		while (bx >= 0) {
			if (block[bx].has()) {
				int bl = bx*BLKSZ;
				int br = min(block_m, bl+BLKSZ);
				for (int i = br-1; i >= bl; i--) {
					if (get_inblock(i))
						return (i<<6) + 63-clzll(i);
				}
			}
			bx--;
		}
		return -1;
	}
	int next_one(int x, int n) { // [x, n-1]
		if (x < 0 || x >= n)	return -1;
		/* find in a 64-bit integer */
		while (x < n && (x&63) != 0) {
			if (get(x, n) == 1)	return x;
			x++;
		};
		if (x >= n)	return -1;
		/* find in a block */
		int cx = x>>6;
		while (cx < block_m && (cx%BLKSZ) != 0) {
			if (get_inblock(cx))
				return (cx<<6) + ctzll(cx);
			cx++;
		} 
		if (cx >= block_m)	return -1;
		/* find in a block chain */
		int bx = cx/BLKSZ;
		while (bx < block_n) {
			if (block[bx].has()) {
				int bl = bx*BLKSZ;
				int br = min(block_m, bl+BLKSZ);
				for (int i = bl; i < br; i++) {
					if (get_inblock(i))
						return (i<<6) + ctzll(i);
				}
			}
			bx++;
		}
		return -1;
	}
	void flip(int l, int r, int n) {
		if (get(l, n) == 0) {
			int t = prev_one(l-1, n);
			if (t >= 0)
				l = t;
		}
		if (get(r, n) == 0) {
			int t = next_one(r+1, n);
			if (t >= 0)
				r = t;
		}
		{
			{
				int tmpl = l, tmpr = r;
				while (l <= r && (l&63) != 0)
					flip(l), l++;
				while (r >= l && (r&63) != 63)
					flip(r), r--;
				update_block(tmpl, n);
				update_block(tmpr, n);
				if (l > r)
					return ;
			}
			int cl = l>>6, cr = r>>6;
			{
				int tmpl = cl, tmpr = cr;
				while (cl <= cr && (cl%BLKSZ) != 0)
					A[cl] = ~A[cl], cl++;
				while (cr >= cl && (cr%BLKSZ) != BLKSZ-1)
					A[cr] = ~A[cr], cr--;
				update_block(tmpl<<6, n);
				update_block(tmpr<<6, n);
				if (cl > cr)
					return ;
			}
			int bl = cl/BLKSZ, br = cr/BLKSZ;
			for (int i = bl; i <= br; i++)
				block[i].inv ^= 1;
		}
	}
	void print(int n) {
		for (int i = 0; i < n; i++) {
			printf("%d", get(i, n));
			if (i%64 == 63)
				puts("");
		}
		puts("");
	}
	void encode(int n) {
		static char s[MAXN];
		int idx = -1;
		for (int i = 0; i < n; ) {
			int t = 0;
			for (int j = 0; j < 4 && i < n; j++, i++) {
				if (get(i, n))
					t |= 1<<j;
			}
			s[++idx] = "0123456789ABCDEF"[t];
		}
		while (idx > 0 && s[idx] == '0')
			idx--;
		while (idx >= 0)
			putchar(s[idx]), idx--;
		puts("");
	}
} g;
int main() {
	int cases = 0;
	int testcase, n, m;
	static char s[1048576];
	scanf("%d", &testcase);
	while (testcase--) {
		scanf("%d %d", &n, &m);
		scanf("%s", s);
		g.to_vector(s, n);
		for (int i = 0; i < m; i++) {
			int l, r;
			scanf("%d %d", &l, &r);
			r = n-r, l = n-l;
			swap(l, r);
			g.flip(l, r, n);
		}
		g.encode(n);
	}
	return 0;
}
