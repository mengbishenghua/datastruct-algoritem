package function;

/**
 * @author czx
 * @date 2019/12/24
 */
@FunctionalInterface
public interface Consumer<T> {
    void forEach(T t);
}
