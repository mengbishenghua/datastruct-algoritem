package datastruct.tree;

/**
 * AVL树：AVL树是平衡搜索树，空树是AVL树，AVL树定义：
 * 1.左子树和右子树是AVL树
 * 2.左子树的高度-右子树的高度<=1
 * 3.它是搜索树
 * 相比于搜索树，多了一个高度和平衡因子函数进行旋转
 *
 * @author czx
 * @date 2019/12/26
 */
public class AvlTree<T extends Comparable<? super T>> {
    AvlNode root;

    public void insert(T t) {
        root = insert(root, t);
    }

    public void remove(T t) {
        root = remove(root, t);
    }

    public void inOrder() {
        inOrder(root);
    }

    public void inOrder(AvlNode node) {
        if (node == null) {
            return;
        }
        if (node.left != null) {
            inOrder(node.left);
        }
        System.out.print(node.value + " ");
        if (node.right != null) {
            inOrder(node.right);
        }
    }

    private AvlNode insert(AvlNode node, T t) {
        if (node == null) {
            return new AvlNode(t);
        }
        int comp = t.compareTo(node.value);
        if (comp < 0) {
            node.left = insert(node.left, t);
        } else if (comp > 0) {
            node.right = insert(node.right, t);
        }
        return balance(node);
    }

    private AvlNode remove(AvlNode node, T t) {
        if (node == null) {
            return null;
        }
        int compare = t.compareTo(node.value);

        if (compare < 0) {
            node.left = remove(node.left, t);
        } else if (compare > 0) {
            node.right = remove(node.right, t);
        } else if (node.left != null && node.right != null) {
            T value = findMin(node.right).value;
            node.right = remove(node.right, value);
            node.value = value;
        } else {
            node = node.left != null ? node.left : node.right;
        }
        return balance(node);
    }

    private AvlNode findMin(AvlNode node) {
        while (node.left != null) {
            node = node.left;
        }
        return node;
    }

    // 平衡树
    private AvlNode balance(AvlNode node) {
        if (node == null) {
            return null;
        }
        // 如果左子树大于右子树,说明插入的节点是在左子树
        if (height(node.left) - height(node.right) > 1) {
            // 如果左子树的左子树高度大于左子树的右子树高度，单旋转
            if (height(node.left.left) >= height(node.left.right)) {
                // 单旋转
                node = rotateWithLeft(node);
            } else {
                // 双旋转
                node = doubleWithLeft(node);
            }
        } else if (height(node.right) - height(node.left) > 1) {
            if (height(node.right.right) >= height(node.right.left)) {
                node = rotateWithRight(node);
            } else {
                node = doubleWithRight(node);
            }
        }
        node.h = Math.max(height(node.left), height(node.right)) + 1;
        return node;
    }

    // 左单旋
    private AvlNode rotateWithLeft(AvlNode k2) {
        AvlNode k1 = k2.left;
        k2.left = k1.right; // 当前节点左子树指向他的左节点的右子树
        k1.right = k2; // 当前节点的右子树指向当前节点
        k2.h = Math.max(height(k2.left), height(k2.right)) + 1;// 调整后的右子树的最大高度
        k1.h = Math.max(height(k1.left), k2.h) + 1; // 调整后的根节点的最大高度
        return k1;
    }

    // 右单旋
    private AvlNode rotateWithRight(AvlNode k1) {
        AvlNode k2 = k1.right;
        k1.right = k2.left;
        k2.left = k1;
        k1.h = Math.max(height(k1.left), height(k1.right)) + 1;
        k2.h = Math.max(height(k2.right), k1.h) + 1;
        return k2;
    }

    // 左双旋
    private AvlNode doubleWithLeft(AvlNode k3) {
        k3.left = rotateWithRight(k3.left);
        return rotateWithLeft(k3);
    }

    // 右双旋
    private AvlNode doubleWithRight(AvlNode k2) {
        k2.right = rotateWithLeft(k2.right);
        return rotateWithRight(k2);
    }

    private int height(AvlNode node) {
        return node == null ? -1 : node.h;
    }

    private class AvlNode {
        T value;
        AvlNode left;
        AvlNode right;
        int h;

        public AvlNode(T value) {
            this(value, null, null, 0);
        }

        public AvlNode(T value, AvlNode left, AvlNode right, int h) {
            this.value = value;
            this.left = left;
            this.right = right;
            this.h = h;
        }
    }
}
