package datastruct.tree;

/**
 * @author czx
 * @date 2019/12/26
 */
public class BinarySearchTreeTest {
    public static void main(String[] args) {
        BinarySearchTree<Integer>br = new BinarySearchTree<>();
        br.insert(6);
        br.insert(2);
        br.insert(8);
        br.insert(1);
        br.insert(5);
        br.insert(3);
        br.insert(4);
        br.inOrder();

        System.out.println();
        br.remove(2);
        br.inOrder();
    }
}
