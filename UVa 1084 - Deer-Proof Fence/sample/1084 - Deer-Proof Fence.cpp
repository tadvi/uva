#include <stdio.h>
#include <stdio.h>
#include <math.h>
#include <vector>
#include <assert.h>
#include <algorithm>
using namespace std;
#define eps 1e-8
struct Pt {
    double x, y;
    Pt(double a = 0, double b = 0):
    	x(a), y(b) {}	
	Pt operator-(const Pt &a) const {
        return Pt(x - a.x, y - a.y);
    }
    Pt operator+(const Pt &a) const {
        return Pt(x + a.x, y + a.y);
    }
    Pt operator*(const double a) const {
        return Pt(x * a, y * a);
    }
    bool operator==(const Pt &a) const {
    	return fabs(x - a.x) < eps && fabs(y - a.y) < eps;
	}
    bool operator<(const Pt &a) const {
		if (fabs(x - a.x) > eps)
			return x < a.x;
		if (fabs(y - a.y) > eps)
			return y < a.y;
		return false;
	}
	double length() {
		return hypot(x, y);
	}
	void read() {
		scanf("%lf %lf", &x, &y);
	}
};
double dot(Pt a, Pt b) {
	return a.x * b.x + a.y * b.y;
}
double cross(Pt o, Pt a, Pt b) {
    return (a.x-o.x)*(b.y-o.y)-(a.y-o.y)*(b.x-o.x);
}
double cross2(Pt a, Pt b) {
    return a.x * b.y - a.y * b.x;
}
int between(Pt a, Pt b, Pt c) {
	return dot(c - a, b - a) >= -eps && dot(c - b, a - b) >= -eps;
}
int onSeg(Pt a, Pt b, Pt c) {
	return between(a, b, c) && fabs(cross(a, b, c)) < eps;
}
struct Seg {
	Pt s, e;
	double angle;
	int label;
	Seg(Pt a = Pt(), Pt b = Pt(), int l=0):s(a), e(b), label(l) {
//		angle = atan2(e.y - s.y, e.x - s.x);
	}
	bool operator<(const Seg &other) const {
		if (fabs(angle - other.angle) > eps)
			return angle > other.angle;
		if (cross(other.s, other.e, s) > -eps)
			return true;
		return false;
	}
	bool operator!=(const Seg &other) const {
		return !((s == other.s && e == other.e) || (e == other.s && s == other.e));
	}
};
Pt getIntersect(Seg a, Seg b) {
	Pt u = a.s - b.s;
    double t = cross2(b.e - b.s, u)/cross2(a.e - a.s, b.e - b.s);
    return a.s + (a.e - a.s) * t;
}
double getAngle(Pt va, Pt vb) { // segment, not vector
	return acos(dot(va, vb) / va.length() / vb.length());
}
Pt rotateRadian(Pt a, double radian) {
	double x, y;
	x = a.x * cos(radian) - a.y * sin(radian);
	y = a.x * sin(radian) + a.y * cos(radian);
	return Pt(x, y);
}
int monotone(int n, Pt p[], Pt ch[]) {
    sort(p, p+n);
    int i, m = 0, t;
    for(i = 0; i < n; i++) {
        while(m >= 2 && cross(ch[m-2], ch[m-1], p[i]) <= 0)
            m--;
        ch[m++] = p[i];
    }
    for(i = n-1, t = m+1; i >= 0; i--) {
        while(m >= t && cross(ch[m-2], ch[m-1], p[i]) <= 0)
            m--;
        ch[m++] = p[i];
    }
    return m-1;
}
const double pi = acos(-1);
double computeFenceLen(Pt ch[], int n, double r) {
	if (n == 1)
		return r * pi * 2;
	if (n == 2)
		return r * pi * 2 + (ch[0] - ch[1]).length() * 2;
	double ret = 0;
	for (int i = 0, j = n-1; i < n; j = i++)
		ret += (ch[i] - ch[j]).length();
	ret += r * pi * 2;
	return ret;
}
int main() {
    int n, R, cases = 0;
    while (scanf("%d %d", &n, &R) == 2 && n) {
    	Pt D[32], ch[32];
    	for (int i = 0; i < n; i++)
    		scanf("%lf %lf", &D[i].x, &D[i].y);
    	double dp[1024] = {};
    	for (int i = 1; i < (1<<n); i++) {
    		double &val = dp[i];
    		val = 1e+30;
    		for (int j = (i-1)&i; j; j = (j-1)&i)
    			val = min(val, dp[j] + dp[i-j]);
    		
    		int m = 0;
    		Pt A[32];
    		for (int j = 0; j < n; j++) {
    			if ((i>>j)&1)
    				A[m++] = D[j];
    		}
    		
    		int cn = monotone(m, A, ch);
    		double len = computeFenceLen(ch, cn, R);
//    		for (int j = 0; j < cn; j++)
//    			printf("%lf %lf\n", ch[j].x, ch[j].y);
//    		printf("length - %lf\n", len);
    		val = min(val, len);
    	}
    	printf("Case %d: length = %.2lf\n", ++cases, dp[(1<<n)-1]);
    }
    return 0;
}
/*
3 2 
0 0 
2 0 
10 0 

5 3
7 8
9 9
9 9
9 9
2 2

3 3
7 8
9 9
2 2

5 0
4 2
1 5
8 5
2 2
4 9
 */
