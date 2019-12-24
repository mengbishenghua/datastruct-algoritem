package function;

/**
 * @author czx
 * @date 2019/12/24
 */
public interface Iterable<E> {

    Iterator<E> iterator();

    default void forEach(Consumer<? super E> action) {
    }
}
