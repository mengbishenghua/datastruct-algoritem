package com.datastruct.queue;

/**
 * @author czx
 * @date 2019/12/24
 */
public class QueueTest {
    public static void main(String[] args) {
        Queue<Integer> q = new Queue<>();
        for (int i = 0; i < 10; i++){
            q.offer(i);
        }

        q.forEach(e -> System.out.print(e + " "));
        System.out.println();
        System.out.println("q.size() = " + q.size());
        Integer pop = q.pop();
        System.out.println("pop = " + pop);
        q.pop();
        q.forEach(e -> System.out.print(e + " "));
        System.out.println("q.size() = " + q.size());
    }
}
