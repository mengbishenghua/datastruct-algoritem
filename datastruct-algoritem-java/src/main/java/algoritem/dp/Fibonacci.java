package algoritem.dp;

/**
 * @author czx
 * @date 2020/1/1
 */
public class Fibonacci {
    public static void main(String[] args) {
        System.out.println("fibonacci(5) = " + fibonacci(5));
        System.out.println("fibonacci(10000) = " + fibonacci(1000000));
    }

    public static int fibonacci(int n) {
        if (n <= 1) {
            return 1;
        }
        int last = 1;
        int nextToLast = 1;
        int answer = 0;
        for (int i = 2; i <= n; i++) {
            answer = last + nextToLast;
            nextToLast = last;
            last = answer;
        }
        return answer;
    }
}
