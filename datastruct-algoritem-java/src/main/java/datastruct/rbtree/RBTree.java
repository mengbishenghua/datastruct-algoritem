package datastruct.rbtree;

/**
 * @author czx
 * @date 2020/1/2
 */
public class RBTree<T extends Comparable<? super T>> {
    private RBTreeNode header;
    private RBTreeNode nullNode;
    private static final int BLACK = 1;
    private static final int RED = 0;

    private RBTreeNode current;
    private RBTreeNode parent;
    private RBTreeNode grand;
    private RBTreeNode great;

    public RBTree() {
        nullNode = new RBTreeNode(null);
        nullNode.left = nullNode.right = nullNode;
        header = new RBTreeNode(null);
        header.left = header.right = nullNode;
    }

    public void insert(T item) {
        current = parent = grand = header;
        nullNode.element = item;

        while (compare(item, current) != 0) {
            great = grand;
            grand = parent;
            parent = current;
            current = compare(item, current) < 0 ? current.left : current.right;

            if (current.left.color == RED && current.right.color == RED) {
                handleReorient(item);
            }
        }

        if (current != nullNode) {
            return;
        }
        current = new RBTreeNode(item, nullNode, nullNode);
        if (compare(item, parent) < 0) {
            parent.left = current;
        } else {
            parent.right = current;
        }
        handleReorient(item);
    }

    private void handleReorient(T item) {
        current.color = RED;
        current.left.color = BLACK;
        current.right.color = BLACK;

        if (parent.color == RED) { // 有rotate
            grand.color = RED;
            if ((compare(item, grand) < 0) != (compare(item, parent) < 0)) {
                parent = rotate(item, grand);
            }
            current = rotate(item, great);
            current.color = BLACK;
        }
        header.right.color = BLACK;
    }

    private RBTreeNode rotate(T item, RBTreeNode parent) {
        if (compare(item, parent) < 0) {
            return parent.left = compare(item, parent) < 0 ? rotateWithLeftChild(parent.left) : // 左子树左旋: LL
                    rotateWithRightChild(parent.left); // 左子树右旋: LR
        } else {
            return parent.right = compare(item, parent) < 0 ? rotateWithLeftChild(parent.right) : // 右子树左旋:RL
                    rotateWithRightChild(parent.right); // 右子树右旋: RR
        }
    }

    private int compare(T item, RBTreeNode p) {
        if (p == header) { // 特殊处理根节点
            return 1;
        }
        return item.compareTo(p.element);
    }

    // 子树左旋转:LL
    private RBTreeNode rotateWithLeftChild(RBTreeNode node) {
        return null;
    }

    // 子树右旋: LR
    private RBTreeNode rotateWithRightChild(RBTreeNode node) {
        return null;
    }

    private class RBTreeNode {
        T element;
        RBTreeNode left;
        RBTreeNode right;
        int color;

        public RBTreeNode(T element) {
            this(element, null, null);
        }

        public RBTreeNode(T element, RBTreeNode left, RBTreeNode right) {
            this.element = element;
            this.left = left;
            this.right = right;
        }
    }
}
