package datastruct.linkedlist;

import datastruct.List;
import function.Consumer;
import function.Iterator;

import java.util.Objects;

/**
 * 双向链表实现的列表
 *
 * @author czx
 * @date 2019/12/24
 */
public class LinkedList<E> implements List<E> {
    private Node head;
    private Node tail;
    private int size;

    public LinkedList() {
        head = new Node();
        tail = new Node();
        head.next = tail;
        tail.prev = head;
    }

    @Override
    public boolean add(E e) {
        return add(size, e);
    }

    @Override
    public boolean add(int index, E e) {
        if (head.next == tail) {
            Node node = new Node(e, head, tail);
            head.next = tail.prev = node;
        } else {
            Node curr = head;
            for (int i = 0; i < index; i++) {
                curr = curr.next;
            }
            Node node = new Node(e, curr, curr.next);
            curr.next = curr.next.prev = node;
        }
        size++;
        return true;
    }

    @Override
    public E get(int index) {
        if (index >= size) {
            throw new IndexOutOfBoundsException("index out of bound");
        }
        Node curr = head.next;
        for (int i = 0; i < index; i++) {
            curr = curr.next;
        }
        return curr.e;
    }

    @Override
    public E remove(int index) {
        if (index >= size) {
            throw new IndexOutOfBoundsException("index out of bound");
        }
        Node curr = head;
        for (int i = 0; i < index; i++) {
            curr = curr.next;
        }
        Node oldValue = curr.next;
        curr.next = oldValue.next;
        oldValue.next.prev = curr;
        size--;
        return oldValue.e;
    }

    @Override
    public boolean isEmpty() {
        return size() == 0;
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean contains(Object obj) {
        return indexOf(obj) > 0;
    }

    @Override
    public int indexOf(Object obj) {
        Node curr = head.next;
        int i = 0;
        if (obj == null) {
            while (curr != tail) {
                if (curr.e == null) {
                    return i;
                }
                curr = curr.next;
                i++;
            }
        } else {
            while (curr != tail) {
                if (curr.e.equals(obj)) {
                    return i;
                }
                curr = curr.next;
                i++;
            }
        }
        return -1;
    }

    @Override
    public void clear() {
        head.next = tail;
        tail.prev = head;
        size = 0;
    }

    @Override
    public Iterator<E> iterator() {
        return new Itr();
    }

    @Override
    public void forEach(Consumer<? super E> action) {
        Objects.requireNonNull(action);
        Node curr = head.next;
        while (curr != tail) {
            if (curr.e != null) {
                action.forEach(curr.e);
            }
            curr = curr.next;
        }
    }

    private class Node {
        E e;
        Node prev;
        Node next;

        Node() {
        }

        Node(E e, Node prev, Node next) {
            this.e = e;
            this.prev = prev;
            this.next = next;
        }
    }

    private class Itr implements Iterator<E> {
        Node curr = LinkedList.this.head.next;

        @Override
        public boolean hasNext() {
            return curr != tail;
        }

        @Override
        public E next() {
            Node res = curr;
            curr = curr.next;
            return res.e;
        }

        @Override
        public void remove() {
            Node rm = curr.prev;
            rm.prev.next = curr;
            curr.prev = rm.prev;
            size--;
        }

        @Override
        public void forEachRemaining(Consumer action) {

        }
    }
}
