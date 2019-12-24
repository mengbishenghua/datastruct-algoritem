package algoritem.sort;


import algoritem.util.SortUtil;

/**
 * @author czx
 * @date 2019/12/13 16:15
 */
public class SelectSort {
    public static void main(String[] args) {
//        int[] arr = {43, 1, 5, 67, 234, 63, 57, 15, 96, 38};
        int[] arr = SortUtil.createRandom(10000, 1, 1000);
        SortUtil.duration("selectSort", SelectSort::selectSort, arr);
    }

    // {43, 1, 5, 67, 234, 63, 57, 15, 96, 38};
    private static void selectSort(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            int min = i;
            // 找到最小值的索引，然后和当前元素交换
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[min] > arr[j])
                    min = j;
            }
            if (min != i)
                SortUtil.swap(arr, min, i);
        }
    }
}
