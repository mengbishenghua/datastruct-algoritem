package algoritem.search;


import static algoritem.sort.InsertSort.insertSort;

/**
 * 二分查找法要先排序
 *
 * @author czx
 * @date 2019/12/13 18:10
 */
public class BinarySearch {
    public static void main(String[] args) {
        int[] arr = {43, 1, 5, 67, 234, 63, 57, 15, 96, 38};
        insertSort(arr);
        int i = binarySearch(arr, 5);
        System.out.println(i);
        System.out.println(arr[i]);

        boolean b = binarySearchRecursion(arr, 100, 0, arr.length - 1);
        System.out.println("b = " + b);
    }

    public static int binarySearch(int[] arr, int x) {
        int start = 0;
        int end = arr.length;
        while (start < end) {
            int mid = start + (end - start) / 2;
            if (arr[mid] > x) {
                end = mid;
            } else if (arr[mid] < x) {
                start = mid + 1;
            } else {
                return mid;
            }
        }
        return -1;
    }

    public static boolean binarySearchRecursion(int[] arr, int x, int start, int end) {
        int mid = start + (end - start) >> 2;
        if (mid > end || mid < start){
            return false;
        }
        if (arr[mid] == x) {
            return true;
        } else if (arr[mid] > x) {
            return binarySearchRecursion(arr, x, start, mid - 1);
        } else {
            return binarySearchRecursion(arr, x, mid + 1, end);
        }
    }
}
