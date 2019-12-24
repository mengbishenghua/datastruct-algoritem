package com.algorithm.sort;

import com.algorithm.util.SortUtil;

import java.util.Arrays;

/**
 * 堆排序，堆很重要，很多的贪心算法都有堆的思想
 * @author czx
 * @date 2019/12/14 17:40
 */
public class HeapSort {
    public static void main(String[] args) {
        int[] arr = SortUtil.createRandom(1000000, 1, 1000);
        SortUtil.duration("heap sort: ", HeapSort::heapSort, arr);
//        System.out.println(Arrays.toString(arr));
    }

    public static void heapSort(int[] arr) {
        if (arr == null || arr.length < 2) {
            return;
        }

        for (int i = 0; i < arr.length; i++) {
            // 先变为大根堆，未必有序,
            // 什么是大根堆：就是每棵二叉树的根节点的值最大(不一定是整棵树，小树也同理)
            heapInsert(arr, i);
        }

        int size = arr.length;
        SortUtil.swap(arr, 0, --size);
        while (size > 0) {
            // 排序，把最大的数放数组最后面，遍历从0到不包括最后面那个元素，转为大根堆，依次类推
            heapIfy(arr, 0, size);
            SortUtil.swap(arr, 0, --size);
        }
    }

    // 将数组转为大根堆
    private static void heapInsert(int[] arr, int i) {
        // 如果子节点大于父节点，交换，
        while (arr[i] > arr[(i - 1) / 2]) {
            SortUtil.swap(arr, i, (i - 1) / 2);
            i = (i - 1) / 2;
        }
    }

    // 将0号位置往下沉，直到符合大根堆
    private static void heapIfy(int[] arr, int i, int size) {
        int left = 2 * i + 1;
        while (left < size) {
            // 如果右孩子不越界，获取左右孩子最大值索引
            int max = left + 1 < size && arr[left + 1] > arr[left] ? left + 1 : left;
            // 获取父亲和左或右孩子的最大值索引
            max = arr[max] > arr[i] ? max : i;
            // 如果父节点最大，不用交换,因为我们只动了0和--size的交换，因为我们一开始就已经转为大根堆了(heapInsert()方法)，
            // 也就是说其他的数符合大根堆，就0号位置不符合，所以让0号位置不断下沉，调整它的位置
            // 直到符合大根堆
            if (max == i) {
                break;
            }
            SortUtil.swap(arr, max, i);
            // i往下走
            i = max;
            left = 2 * i + 1;
        }
    }
}
