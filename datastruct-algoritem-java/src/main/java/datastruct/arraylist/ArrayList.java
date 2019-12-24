package datastruct.arraylist;

import datastruct.List;
import function.Consumer;
import function.Iterator;

import java.util.Arrays;
import java.util.ConcurrentModificationException;
import java.util.NoSuchElementException;
import java.util.Objects;

/**
 * 数组实现的列表, 大部分参考JDK8源码
 *
 * @author czx
 * @date 2019/12/24
 */
public class ArrayList<E> implements List<E> {
    private static final Object[] DEFAULTCAPACITY_EMPTY_ELEMENTDATA = {};
    private static final int MAX_ARRAY_SIZE = Integer.MAX_VALUE - 8;
    protected Object[] element;
    private int size;
    private static int DEFAULT_CAPACITY = 16;
    private transient int modCount = 0;

    public ArrayList() {
        element = new Object[DEFAULT_CAPACITY];
    }

    public ArrayList(int capacity) {
        if (capacity <= 0 || capacity >= Integer.MAX_VALUE) {
            throw new IndexOutOfBoundsException("capacity out of bound");
        }
        element = new Object[capacity];
        size = 0;
    }

    @Override
    public boolean add(E e) {
        ensureCapacityInternal(size + 1);
        element[size++] = e;
        return true;
    }

    @Override
    public boolean add(int index, E e) {
        ensureCapacityInternal(size + 1);
        System.arraycopy(element, index, element, index + 1,
                size - index);
        element[index] = e;
        size++;
        return true;
    }

    @Override
    @SuppressWarnings("unchecked")
    public E get(int index) {
        return (E) element[index];
    }

    @Override
    public boolean isEmpty() {
        return size() == 0;
    }

    @Override
    public E remove(int index) {
        rangeCheck(index);
        modCount++;
        E oldValue = elementData(index);
        int numMoved = size - index - 1;
        if (numMoved > 0) {
            // 0, 1, 2, 3, 4
            // 从index+1开始往左移动一位
            System.arraycopy(element, index + 1, element, index,
                    numMoved);
        }
        element[--size] = null;
        return oldValue;
    }

    // minCapacity = size + 1
    private void ensureCapacityInternal(int minCapacity) {
        if (element == DEFAULTCAPACITY_EMPTY_ELEMENTDATA) {
            minCapacity = Math.max(DEFAULT_CAPACITY, minCapacity);
        }
        modCount++;
        if (minCapacity - element.length > 0) {
            // overflow-conscious code
            int oldCapacity = element.length;
            int newCapacity = oldCapacity + (oldCapacity >> 1);
            if (newCapacity - minCapacity < 0)
                newCapacity = minCapacity;
            if (newCapacity - MAX_ARRAY_SIZE > 0)
                newCapacity = hugeCapacity(minCapacity);
            // minCapacity is usually close to size, so this is a win:
            element = Arrays.copyOf(element, newCapacity);
        }
    }

    private int hugeCapacity(int minCapacity) {
        if (minCapacity < 0) // overflow
            throw new OutOfMemoryError();
        return (minCapacity > MAX_ARRAY_SIZE) ?
                Integer.MAX_VALUE :
                MAX_ARRAY_SIZE;
    }

    @SuppressWarnings("unchecked")
    private E elementData(int index) {
        return (E) element[index];
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean contains(Object obj) {
        return indexOf(obj) >= 0;
    }

    @Override
    public int indexOf(Object o) {
        if (o == null) {
            for (int i = 0; i < size; i++)
                if (element[i] == null)
                    return i;
        } else {
            for (int i = 0; i < size; i++)
                if (o.equals(element[i]))
                    return i;
        }
        return -1;
    }

    @Override
    public void clear() {
        modCount++;

        // clear to let GC do its work
        for (int i = 0; i < size; i++)
            element[i] = null;

        size = 0;
    }

    @Override
    public Iterator<E> iterator() {
        return new Itr();
    }

    @Override
    @SuppressWarnings("unchecked")
    public void forEach(Consumer<? super E> action) {
        Objects.requireNonNull(action);
        for (Object t : element) {
            if (t != null) {
                action.forEach((E) t);
            }
        }
    }

    private void rangeCheck(int index) {
        if (index >= size)
            throw new IndexOutOfBoundsException(outOfBoundsMsg(index));
    }

    private String outOfBoundsMsg(int index) {
        return "Index: " + index + ", Size: " + size;
    }

    @SuppressWarnings("unchecked")
    private class Itr implements Iterator<E> {
        int cursor;       // index of next element to return
        int lastRet = -1; // index of last element returned; -1 if no such
        int expectedModCount = modCount;

        @Override
        public boolean hasNext() {
            return cursor != size;
        }

        @Override
        public E next() {
            checkForComodification();
            int i = cursor;
            if (i >= size) {
                throw new NoSuchElementException();
            }
            Object[] element = ArrayList.this.element;
            if (i >= element.length) {
                throw new ConcurrentModificationException();
            }
            cursor = i + 1;
            return (E) element[lastRet = i];
        }

        @Override
        public void remove() {
            if (lastRet < 0)
                throw new IllegalStateException();
            checkForComodification();

            try {
                ArrayList.this.remove(lastRet);
                cursor = lastRet;
                lastRet = -1;
                expectedModCount = modCount;
            } catch (IndexOutOfBoundsException ex) {
                throw new ConcurrentModificationException();
            }
        }

        @Override
        public void forEachRemaining(Consumer<? super E> action) {
            Objects.requireNonNull(action);
            final int size = ArrayList.this.size;
            int i = cursor;
            if (i >= size) {
                return;
            }
            final Object[] elementData = ArrayList.this.element;
            if (i >= elementData.length) {
                throw new ConcurrentModificationException();
            }
            while (i != size && modCount == expectedModCount) {
                action.forEach((E) elementData[i++]);
            }
            // update once at end of iteration to reduce heap write traffic
            cursor = i;
            lastRet = i - 1;
            checkForComodification();
        }

        final void checkForComodification() {
            if (modCount != expectedModCount)
                throw new ConcurrentModificationException();
        }
    }
}
