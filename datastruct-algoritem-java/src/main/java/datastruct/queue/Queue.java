package com.datastruct.queue;

import com.datastruct.linkedlist.LinkedList;
import com.function.Iterator;

/**
 * 双向链表实现的队列
 *
 * @author czx
 * @date 2019/12/24
 */
public class Queue<E> extends LinkedList<E> {

    public boolean offer(E e) {
        return add(e);
    }

    @Override
    public Iterator<E> iterator() {
        throw new UnsupportedOperationException("un supported operation");
    }

    public E pop() {
        return remove(size() - 1);
    }
}
