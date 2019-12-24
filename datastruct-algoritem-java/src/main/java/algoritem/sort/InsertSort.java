package com.algorithm.sort;

import com.algorithm.util.SortUtil;

import java.util.Arrays;

/**
 * @author czx
 * @date 2019/12/13 16:20
 */
public class InsertSort {
    public static void main(String[] args) {
        int[] arr = SortUtil.createRandom(10000, 1, 1000);
        SortUtil.duration("insertSort", InsertSort::insertSort, arr);
    }

    // {43, 1, 5, 67, 234, 63, 57, 15, 96, 38};
    public static void insertSort(int[] arr) {
        for (int i = 1; i < arr.length; i++) {
            for (int j = i - 1; j >= 0 && arr[j] > arr[j + 1]; j--) {
                SortUtil.swap(arr, j, j + 1);
            }
        }
    }
}
