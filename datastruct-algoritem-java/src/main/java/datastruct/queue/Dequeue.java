package datastruct.queue;

/**
 * 基于双向链表实现的双端队列
 *
 * @author czx
 * @date 2019/12/24
 */
public class Dequeue<E> extends Queue<E> {
    public boolean offer(E e) {
        return add(e);
    }

    public boolean offerFront(E e) {
        return add(0, e);
    }

    public E popFront() {
        return remove(0);
    }
}
