package algoritem.util;

/**
 * @author czx
 * @date 2019/12/14 14:25
 */
@FunctionalInterface
public interface Consumer<T, U, R> {
    void accept(T t, U u, R r);
}
