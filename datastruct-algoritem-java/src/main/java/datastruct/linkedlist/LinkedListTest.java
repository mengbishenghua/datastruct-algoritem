package com.datastruct.linkedlist;

import com.datastruct.List;
import com.function.Iterator;

/**
 * @author czx
 * @date 2019/12/24
 */
public class LinkedListTest {
    public static void main(String[] args) {
        List<Integer> list = new LinkedList<>();
        for (int i = 0; i < 10; i++) {
            list.add(i);
        }

        list.forEach(e -> System.out.print(e + " "));
        System.out.println("list.size() = " + list.size());
        list.remove(0);
        list.remove(8);
        System.out.println();
        list.forEach(e -> System.out.print(e + " "));

        System.out.println();

        Iterator<Integer> iterator = list.iterator();
        while (iterator.hasNext()) {
            Integer next = iterator.next();
            if (next == 5) {
                iterator.remove();
                continue;
            }
            System.out.print(next + " ");
        }
        System.out.println("list.size() = " + list.size());
    }
}
