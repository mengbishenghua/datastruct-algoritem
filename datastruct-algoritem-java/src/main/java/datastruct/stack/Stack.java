package datastruct.stack;


import datastruct.arraylist.ArrayList;
import function.Consumer;

import java.util.Objects;

/**
 * @author czx
 * @date 2019/12/24
 */
public class Stack<E> extends ArrayList<E> {
    @Override
    public boolean add(int index, E e) {
        throw new UnsupportedOperationException("un supported operation");
    }

    public boolean push(E e) {
        return this.push(e);
    }

    public E peek() {
        return get(size() - 1);
    }

    public E pop() {
        return remove(size() - 1);
    }

    @SuppressWarnings("unchecked")
    @Override
    public void forEach(Consumer<? super E> action) {
        Objects.requireNonNull(action);
        int index = size() - 1;
        while (index >= 0) {
            E e = (E) element[index--];
            if (e != null) {
                action.forEach(e);
            }
        }
    }
}
