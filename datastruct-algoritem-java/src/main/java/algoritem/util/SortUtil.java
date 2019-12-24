package com.algorithm.util;

import java.util.concurrent.ThreadLocalRandom;
import java.util.function.Consumer;

/**
 * @author czx
 * @date 2019/12/13 17:19
 */
public class SortUtil {
    public static <T> void swap(int[] arr, int i, int j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }

    public static void duration(String fnName, Consumer<int[]> consumer, int[] arr) {
        long start = System.currentTimeMillis();
        consumer.accept(arr);
        System.out.println(fnName + " run time: " + (System.currentTimeMillis() - start) + "ms");
    }

    public static void duration(String fnName,
                                com.algorithm.util.Consumer<int[], Integer, Integer> consumer,
                                int[] arr, int u, int r) {
        long start = System.currentTimeMillis();
        consumer.accept(arr, u, r);
        System.out.println(fnName + " run time: " + (System.currentTimeMillis() - start) + "ms");
    }

    public static int[] createRandom(int n, int start, int end) {
        ThreadLocalRandom current = ThreadLocalRandom.current();
        int[] result = new int[n];
        for (int i = 0; i < n; i++) {
            int num = current.nextInt(start, end + 1);
            result[i] = num;
        }
        return result;
    }
}
