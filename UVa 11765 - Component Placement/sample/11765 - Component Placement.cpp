#include <stdio.h>
#include <string.h>
#include <math.h>
#include <algorithm>
#include <queue>
#include <stack>
#include <set>
using namespace std;
struct Node {
    int x, y, v;// x->y, v
    int next;
} edge[100005];
int e, head[505], prev[505], record[505];
int level[505], visited[505];
void addEdge(int x, int y, int v) {
    edge[e].x = x, edge[e].y = y, edge[e].v = v;
    edge[e].next = head[x], head[x] = e++;
    edge[e].x = y, edge[e].y = x, edge[e].v = 0;
    edge[e].next = head[y], head[y] = e++;
}
bool buildLevelGraph(int s, int t) {
    memset(level, 0, sizeof(level));
    queue<int> Q;
    Q.push(s), level[s] = 1;
    while(!Q.empty()) {
        int tn = Q.front();
        Q.pop();
        for(int i = head[tn]; i != -1; i = edge[i].next) {
            int y = edge[i].y;
            if(edge[i].v > 0 && level[y] == 0) {
                level[y] = level[tn] + 1;
                Q.push(y);
            }
        }
    }
    return level[t] > 0;
}
int constructBlockingFlow(int s, int t) {
    int ret = 0;
    stack<int> stk;
    memset(visited, 0, sizeof(visited));
    stk.push(s);
    while(!stk.empty()) {
        int now = stk.top();
        if(now != t) {
            for(int i = head[now]; i != -1; i = edge[i].next) {
                int y = edge[i].y;
                if(visited[y] || level[y] != level[now] + 1)
                    continue;
                if(edge[i].v > 0) {
                    stk.push(y), prev[y] = now, record[y] = i;
                    break;
                }
            }
            if(stk.top() == now)
                stk.pop(), visited[now] = 1;
        } else {
            int flow = 1e+9, bottleneck;
            for(int i = t; i != s; i = prev[i]) {
                int ri = record[i];
                flow = min(flow, edge[ri].v);
            }
            for(int i = t; i != s; i = prev[i]) {
                int ri = record[i];
                edge[ri].v -= flow;
                edge[ri^1].v += flow;
                if(edge[ri].v == 0)
                    bottleneck = prev[i];
            }
            while(!stk.empty() && stk.top() != bottleneck)
                stk.pop();
            ret += flow;
        }
    }
    return ret;
}
int maxflowDinic(int s, int t) {
    int flow = 0;
    while(buildLevelGraph(s, t))
        flow += constructBlockingFlow(s, t);
    return flow;
}
int main() {
    int n, m, testcase;
    int U[1024], D[1024], kind, x, y, v;
    scanf("%d", &testcase);
    while(testcase--) {
        e = 0;
        memset(head, -1, sizeof(head));
        
        scanf("%d %d", &n, &m);
        for(int i = 0; i < n; i++)
        	scanf("%d", &U[i]);
        for(int i = 0; i < n; i++)
        	scanf("%d", &D[i]);
    	
    	int source = n + 1, sink = source + 1;
    	const int INF = 0x3f3f3f3f;
        for(int i = 0; i < n; i++) {
        	scanf("%d", &kind);
        	if(kind == 0) {
        		addEdge(source, i, U[i]);
        		addEdge(i, sink, D[i]);
        	} else if(kind == 1) {
        		addEdge(source, i, U[i]);
        		addEdge(i, sink, INF);
        	} else {
        		addEdge(source, i, INF);
        		addEdge(i, sink, D[i]);
        	}
        }
        for(int i = 0; i < m; i++) {
        	scanf("%d %d %d", &x, &y, &v);
        	x--, y--;
        	addEdge(x, y, v);
        	addEdge(y, x, v);
        }
        printf("%d\n", maxflowDinic(source, sink));
    }
    return 0;
}
