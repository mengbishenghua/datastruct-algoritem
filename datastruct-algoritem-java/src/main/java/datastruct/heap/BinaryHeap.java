package datastruct.heap;

/**
 * 优先队列：当前节点的左子树：2i, 右子树：2i+1
 * 优先队列父节点：i/2
 * 这里实现了最小堆
 * 这里的二叉堆根节点从1开始的，0号位置保留
 *
 * @author czx
 * @date 2019/12/30
 */
public class BinaryHeap<T extends Comparable<? super T>> {
    private static final int DEFAULT_CAPACITY = 10;
    private int currentSize;
    private Object[] array;

    public BinaryHeap() {
    }

        public BinaryHeap(int capacity) {
        if (capacity <= 0) {
            throw new IndexOutOfBoundsException("capacity is must not <= 0");
        }
        array = new Object[capacity];
    }

    public BinaryHeap(T[] items) {
        currentSize = items.length;
        array = new Comparable[(currentSize + 2) * 11 / 10];
        int i = 1;
        for (T item : items) {
            array[i++] = item;
        }
        buildHeap();
    }

    @SuppressWarnings("unchecked")
    public void insert(T t) {
        if (currentSize == array.length - 1) {
            enlargeArray(2 * array.length - 1);
        }
        int hole = ++currentSize;
        // 0号位置我们是预留的空位，1号才是堆的根节点
        for (array[0] = t; t.compareTo((T) array[hole / 2]) < 0; hole /= 2) {
            array[hole] = array[hole / 2];
        }
        array[hole] = t;
    }

    @SuppressWarnings("unchecked")
    public T findMin() {
        return (T) array[1];
    }

    public T deleteMin() {
        if (isEmpty()) {
            throw new ArrayIndexOutOfBoundsException("index out of bound");
        }
        T min = findMin();
        // 将根节点设置为空穴
        array[1] = array[currentSize--];
        // 从根节点开始下游调整
        shiftDown(1);
        return min;
    }

    public boolean isEmpty() {
        return currentSize == 0;
    }

    public void makeEmpty() {
        array = new Object[DEFAULT_CAPACITY];
    }

    @SuppressWarnings("unchecked")
    private void shiftDown(int hole) {
        int child = 0;
        T tmp = (T) array[hole]; // 这是currentSize的值
        for (; hole * 2 <= currentSize; hole = child) {
            child = 2 * hole; // 获取左孩子
            // 如果有右孩子，判断右孩子是否小于左孩子
            if (child != currentSize && ((T) array[child + 1]).compareTo((T) array[child]) < 0) {
                child++;
            }
            // 如果左右孩子最小值小于父节点，设置值,hole往下移
            if (((T) array[child]).compareTo(tmp) < 0) {
                array[hole] = array[child];
            } else {
                // 调整完成
                break;
            }
        }
        array[hole] = tmp;
    }

    private void buildHeap() {
        for (int i = currentSize / 2; i > 0; i--) {
            shiftDown(i);
        }
    }

    private void enlargeArray(int newSize) {
        Object[] newArray = new Object[newSize];
        System.arraycopy(array, 1, newArray, 1, Math.min(array.length, newSize));
        currentSize = array.length;
        array = newArray;
    }
}
