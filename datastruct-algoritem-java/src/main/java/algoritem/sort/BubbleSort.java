package algoritem.sort;


import algoritem.util.SortUtil;

/**
 * @author czx
 * @date 2019/12/13 16:18
 */
public class BubbleSort {
    public static void main(String[] args) {
        int[] arr = SortUtil.createRandom(10000, 1, 1000);
        SortUtil.duration("bubbleSort", BubbleSort::bubbleSort, arr);
    }

    // {43, 1, 5, 67, 234, 63, 57, 15, 96, 38};
    private static void bubbleSort(int[] arr) {
        for (int i = 0; i < arr.length-1; i++) {
            for (int j = 0; j < arr.length - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    SortUtil.swap(arr, j, j + 1);
                }
            }
        }
    }
}
