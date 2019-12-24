package datastruct.stack;

/**
 * @author czx
 * @date 2019/12/24
 */
public class StackTest {
    public static void main(String[] args) {
        Stack<Integer> stack = new Stack<>();
        for (int i = 0; i < 10; i++) {
            stack.push(i);
        }

        stack.forEach(e -> System.out.print(e + " "));
    }
}
