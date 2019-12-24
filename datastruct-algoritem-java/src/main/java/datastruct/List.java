package com.datastruct;

import com.function.Iterable;

/**
 * @author czx
 * @date 2019/12/24
 */
public interface List<E> extends Iterable<E> {
    boolean add(E e);

    boolean add(int index, E e);

    E get(int index);

    boolean isEmpty();

    E remove(int index);

    int size();

    boolean contains(Object obj);

    int indexOf(Object obj);

    void clear();
}
