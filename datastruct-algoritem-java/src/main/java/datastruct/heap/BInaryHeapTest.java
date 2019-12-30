package datastruct.heap;

/**
 * @author czx
 * @date 2019/12/30
 */
public class BInaryHeapTest {
    public static void main(String[] args) {
        Integer[] arr = {13, 14, 16, 19, 21, 19, 68, 65, 26, 32, 31};
        BinaryHeap<Integer> heap = new BinaryHeap<>(arr);
        System.out.println(heap.findMin());
        System.out.println(heap.deleteMin());
        heap.insert(10);
        System.out.println(heap.findMin());
    }
}
