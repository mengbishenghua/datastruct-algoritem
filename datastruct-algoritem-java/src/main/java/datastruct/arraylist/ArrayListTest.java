package datastruct.arraylist;

import datastruct.List;

/**
 * @author czx
 * @date 2019/12/24
 */
public class ArrayListTest {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(16);
        for (int i = 0; i < 100; i++){
            list.add(i);
        }

        list.forEach(System.out::println);
        System.out.println("list.size() = " + list.size());
    }
}
