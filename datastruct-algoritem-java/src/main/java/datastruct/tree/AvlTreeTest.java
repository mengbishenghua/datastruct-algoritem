package datastruct.tree;

/**
 * @author czx
 * @date 2019/12/26
 */
public class AvlTreeTest {
    public static void main(String[] args) {
        AvlTree<Integer> tree = new AvlTree<>();
        tree.insert(6);
        tree.insert(2);
        tree.insert(8);
        tree.insert(1);
        tree.insert(5);
        tree.insert(3);
        tree.insert(4);
        tree.inOrder();
    }
}
