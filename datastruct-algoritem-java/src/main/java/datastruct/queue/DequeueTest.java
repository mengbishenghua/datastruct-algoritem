package datastruct.queue;

/**
 * @author czx
 * @date 2019/12/24
 */
public class DequeueTest {
    public static void main(String[] args) {
        Dequeue<Integer> dq = new Dequeue<>();
        for (int i = 0; i < 10; i++) {
            dq.push(i);
        }
        dq.forEach(e -> System.out.print(e + " "));
        System.out.println("dq.size() = " + dq.size());
        System.out.println();
        dq.pushFront(100);
        dq.pushFront(200);
        dq.forEach(e -> System.out.print(e + " "));
        System.out.println("dq.size() = " + dq.size());
        dq.pop();
        dq.popFront();
        dq.forEach(e -> System.out.print(e + " "));
        System.out.println("dq.size() = " + dq.size());
    }
}
