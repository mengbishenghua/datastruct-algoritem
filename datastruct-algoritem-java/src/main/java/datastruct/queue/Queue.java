package datastruct.queue;


import datastruct.linkedlist.LinkedList;
import function.Iterator;

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
        return remove(0);
    }
}
