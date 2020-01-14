package algoritem.dp;

/**
 * @author czx
 * @date 2020/1/1
 */
public class EvalSumN {
    public static void main(String[] args) {
        System.out.println("eval(5) = " + eval(5));
        System.out.println("eval(10) = " + eval(10));
    }

    public static double eval(int n) {
        double[] c = new double[n + 1];
        c[0] = 1;
        for (int i = 1; i <= n; i++) {
            double sum = 0;
            for (int j = 0; j < i; j++) {
                sum += c[j];
            }
            c[i] = 2.0 * sum / i + i;
        }
        return c[n];
    }
}
