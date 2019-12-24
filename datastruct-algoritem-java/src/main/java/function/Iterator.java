package com.function;

import java.util.Objects;

/**
 * @author czx
 * @date 2019/12/24
 */
public interface Iterator<E> {
    boolean hasNext();

    E next();

    default void remove() {
        throw new UnsupportedOperationException("remove");
    }

    default void forEachRemaining(Consumer<? super E> action) {
        Objects.requireNonNull(action);
        while (hasNext())
            action.forEach(next());
    }
}
