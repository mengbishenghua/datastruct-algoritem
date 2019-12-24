package com.algorithm.sort;

import com.algorithm.util.SortUtil;

import java.util.Arrays;

/**
 * @author czx
 * @date 2019/12/13 19:14
 */
public class MergeSort {
    public static void main(String[] args) {
        int[] arr = SortUtil.createRandom(10000, 1, 1000);
        SortUtil.duration("mergeSOrt", MergeSort::mergeSort, arr);
    }

    private static void mergeSort(int[] arr) {
        sort(arr, 0, arr.length - 1);
    }

    // 分
    private static void sort(int[] arr, int l, int r) {
        if (l < r) {
            int m = l + (r - l) / 2;
            sort(arr, l, m);
            sort(arr, m + 1, r);
            merger(arr, l, m, r);
        }
    }

    // 合并
    private static void merger(int[] arr, int l, int m, int r) {
        if (arr.length <= 0)
            return;
        int[] temp = new int[r - l + 1];
        int p1 = l;
        int p2 = m + 1;

        // 按从小到大放入临时数组
        int i = 0;
        while (p1 <= m && p2 <= r) {
            if (arr[p1] < arr[p2]) {
                temp[i++] = arr[p1++];
            } else {
                temp[i++] = arr[p2++];
            }
        }

        // 将两边中未放完的元素放入临时数组
        while (p1 <= m) {
            temp[i++] = arr[p1++];
        }
        while (p2 <= r) {
            temp[i++] = arr[p2++];
        }

        // 将临时数组的元素赋值给原数组
        System.arraycopy(temp, 0, arr, l, temp.length);
    }
}
