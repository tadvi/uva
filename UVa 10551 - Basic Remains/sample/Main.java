import java.io.*;
import java.util.*;
import java.math.*;
class Main { /* UVa 10551 - Basic Remains */
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        while (true) {
            int b = sc.nextInt();
            if (b == 0) break;
            String p_str = sc.next();
            BigInteger p = new BigInteger(p_str, b); // special constructor!
            String m_str = sc.next();
            BigInteger m = new BigInteger(m_str, b); // 2nd parameter is the radix/base
            System.out.println((p.mod(m)).toString(b)); // can output in any radix/base
        }
    }
}