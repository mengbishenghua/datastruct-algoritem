package datastruct.tree;

/**
 * 二叉搜索树
 *
 * @author czx
 * @date 2019/12/26
 */
public class BinarySearchTree<T extends Comparable<? super T>> {
    private Node root;

    public BinarySearchTree() {
    }

    public BinarySearchTree(Node root) {
        this.root = root;
    }

    public void setRoot(Node root) {
        this.root = root;
    }

    public boolean isEmpty() {
        return root == null;
    }

    public boolean contains(T t) {
        return contains(root, t);
    }

    private boolean contains(Node node, T t) {
        if (node == null) {
            return false;
        }
        int res = node.data.compareTo(t);
        if (res > 0) {
            return contains(node.left, t);
        } else if (res < 0) {
            return contains(node.right, t);
        } else {
            return true;
        }
    }

    public T findMin() {
        if (isEmpty()) {
            throw new RuntimeException("tree is empty");
        }
        return findMin(root).data;
    }

    private Node findMin(Node node) {
        while (node.left != null) {
            node = node.left;
        }
        return node;
    }

    public T findMax() {
        if (isEmpty()) {
            throw new RuntimeException("tree is empty");
        }
        return findMax(root).data;
    }

    private Node findMax(Node node) {
        while (node.right != null) {
            node = node.right;
        }
        return node;
    }

    public void insert(T t) {
        root = insert(root, t);
    }

    private Node insert(Node node, T t) {
        if (node == null) { // 根节点
            return new Node(t, null, null);
        }
        int res = node.data.compareTo(t);
        if (res > 0) {
            node.left = insert(node.left, t);
        } else if (res < 0) {
            node.right = insert(node.right, t);
        }
        return node;
    }

    public void remove(T t) {
        root = remove(root, t);
    }

    private Node remove(Node node, T t) {
        if (node == null) {
            return null;
        }
        int res = t.compareTo(node.data);
        if (res < 0) {
            node.left = remove(node.left, t);
        } else if (res > 0) {
            node.right = remove(node.right, t);
        } else if (node.left != null && node.right != null) {
            // 找到了并且有左右子树
            // 删除右子树最小节点，实际是将最小节点的子节点返回，然后父节点的子节点指向返回的节点
            // 将最小节点的值和当前节点交换
            Node min = findMin(node.right);
            node.right = remove(node.right, min.data);
            node.data = min.data;
        } else {
            // 找到了只有一个树
            node = node.left != null ? node.left : node.right;
        }
        return node;
    }

    public void inOrder() {
        inOrder(root);
    }

    public void inOrder(Node node) {
        if (node == null) {
            return;
        }
        inOrder(node.left);
        System.out.print(node.data + " ");
        inOrder(node.right);
    }

    private class Node {
        T data;
        Node left;
        Node right;

        public Node(T data) {
            this(data, null, null);
        }

        public Node(T data, Node left, Node right) {
            this.data = data;
            this.left = left;
            this.right = right;
        }
    }
}
