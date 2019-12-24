package com.algorithm.sort;

import com.algorithm.util.SortUtil;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.ThreadLocalRandom;

/**
 * @author czx
 * @date 2019/12/14 0:06
 */
public class QuickSort {
    public static void main(String[] args) {
        int[] arr = SortUtil.createRandom(10000, 1, 1000);
        SortUtil.duration("quickSort", QuickSort::quickSort, arr);
        arr = SortUtil.createRandom(10000, 1, 1000);
        SortUtil.duration("partitionThreeWay", QuickSort::partitionThreeWay, arr, 0, arr.length - 1);

        arr = SortUtil.createRandom(10000, 1, 1000);
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < arr.length; i++) {
            list.add(arr[i]);
        }
        long start = System.currentTimeMillis();
        Collections.sort(list);
        System.out.println("sort time: " + (System.currentTimeMillis() - start) + "ms");
    }

    private static void quickSort(int[] arr) {
        quickSort(arr, 0, arr.length - 1);
    }

    private static void quickSort(int[] arr, int l, int r) {
        // 如果小于15个元素用插入排序
        if (r - l < 15) {
            InsertSort.insertSort(arr);
            return;
        }
        int pivot = partition(arr, l, r);
        quickSort(arr, l, pivot - 1);
        quickSort(arr, pivot + 1, r);
    }

    // 分区，将比pivot小的放左边，比pivot大的放右边
    // 默认pivot设置为r
    // 这里使用的二路快排
    private static int partition(int[] arr, int l, int r) {
        // 设置一个中轴
        // 随机快排
        SortUtil.swap(arr, ThreadLocalRandom.current().nextInt(l, r), r);
        int pivot = arr[r];
        int left = l - 1;
        int right = r;

        while (left < right) {
            // 从左边找到第一个比pivot大的，准备放到右边
            while (left < right && arr[++left] < pivot) {
            }
            // 从右边找到第一个比pivot小的，准备放入左边
            while (left < right && arr[--right] > pivot) {
            }

            if (left < right) {
                SortUtil.swap(arr, left, right);
            } else {
                break;
            }
        }
        // 将中轴值和左边的第一个大于pivot的替换，返回中轴值对应的索引
        SortUtil.swap(arr, left, r);
        return left;
    }

    /**
     * -------------------------------------------
     * | |       |            | |       |        |
     * |v|  <v   |  ==v       | |       |   >v   |
     * --------lt--------------i---------gt-------
     * i为正在遍历的元素，lt代表指向比v小的索引，gt是指向比v大的索引
     * lt一开始为l，gt一开始为r+1,这是为了方便遍历
     *
     * @param arr
     * @param l
     * @param r
     */
    private static void partitionThreeWay(int[] arr, int l, int r) {
        // 如果小于15个元素用插入排序
        if (r - l < 15) {
            InsertSort.insertSort(arr);
            return;
        }
        // 随机快排
        SortUtil.swap(arr, ThreadLocalRandom.current().nextInt(l, r), r);
        int pivot = arr[l];
        int lt = l - 1;
        int gt = r + 1;
        int i = l;

        while (i < gt) {
            if (arr[i] == pivot)
                i++;
            else if (arr[i] < pivot) {
                SortUtil.swap(arr, i, lt + 1);
                i++;
                lt++;
            } else {
                SortUtil.swap(arr, i, gt - 1);
                gt--;
            }
        }
        partitionThreeWay(arr, l, lt);
        partitionThreeWay(arr, gt, r);

        /*if (r - l < 15) {
            InsertSort.insertSort(arr);
            return;
        }

        int pivot = arr[l];
        int lt = l;
        int gt = r + 1;
        int i = l + 1;

        while (i < gt) {
            // ==v
            if (arr[i] == pivot) {
                i++;
            } else if (arr[i] < pivot) { // < v
                SortUtil.swap(arr, i, lt + 1);
                i++;
                lt++;
            } else { // > v
                SortUtil.swap(arr, i, gt - 1);
                gt--;
            }
        }
        SortUtil.swap(arr, l, lt);
        lt--;
        partitionThreeWay(arr, l, lt);
        partitionThreeWay(arr, gt, r);*/
    }
}
