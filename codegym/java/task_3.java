import java.util.ArrayList;
import java.util.Arrays;
import java.io.*;

public class task_3 {
    public static void main(String[] args) {
        printTimes(10);
        // try {
        // io();
        // } catch (Exception e) {
        // // TODO: handle exception
        // }
        // ;
        testreverseArray();
        testbubbleSort();

    }

    // multiplication table
    public static void printTimes(int num) {
        ArrayList<Integer> line = new ArrayList<Integer>();
        for (int i = 0; i < num; i++) {
            line.add(i + 1);
        }
        for (int n : line) {
            // ArrayList<Integer> print = new ArrayList<Integer>();
            for (int i = 0; i < num; i++) {
                System.out.print(n * (i + 1) + " ");
                // print.add(n * (i + 1));
            }
            System.out.println();
            // System.out.println(print);
        }
    }

    // io
    public static void io() throws Exception {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in)); // 输入流
        String name = reader.readLine();
        String sTime = reader.readLine();
        System.out.print(name + "将在 " + sTime + " 年后征服全世界。哈哈哈哈！");
    }

    // type casting
    public static int sumDigitsInNumber(int number) {
        String sNum = Integer.toString(number);
        Integer result = 0;
        for (int i = 0; i < sNum.length(); i++) {
            // System.out.println(Integer.parseInt(String.valueOf(sNum.charAt(i))));
            result += Integer.parseInt(String.valueOf(sNum.charAt(i)));
        }
        return result;
    }

    // bubble sort
    public static void bubbleSort(int[] array, Boolean ascend) {
        // loop time n - 1
        for (int i = 0; i < array.length - 1; i++) {
            // loop time = n - i - 1
            for (int j = array.length - 1; j > i; j--) {
                if (array[i] >= array[j]) {
                    swap(i, j, array);
                }
            }
        }
        if (!ascend) {
            reverseArray(array);
        }
    }

    // swap
    public static void swap(int i, int j, int[] array) {
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }

    // reverse array
    public static void reverseArray(int[] array) {
        if (array.length < 2) {
            return;
        }

        int loopTimes;
        if (array.length % 2 == 0) {
            loopTimes = array.length / 2;
        } else {
            loopTimes = (array.length - 1) / 2;
        }
        for (int i = 0; i < loopTimes; i++) {
            swap(i, array.length - i - 1, array);
        }
    }

    // test reverse array
    public static void testreverseArray() {
        int[] testArray1 = { 7, 8, 9, 1, 2 };
        int[] testArray2 = { 4, 5, 9, 0 };
        int[] testResult1 = { 2, 1, 9, 8, 7 };
        int[] testResult2 = { 0, 9, 5, 4 };

        reverseArray(testArray1);
        reverseArray(testArray2);

        if (Arrays.equals(testArray1, testResult1) && Arrays.equals(testArray2, testResult2)) {
            System.out.println("reverseArray passed");
        } else {
            System.out.println("reverseArray failed");
        }
    }

    // test bubble sort
    public static void testbubbleSort() {
        int[] testArray1 = { 7, 8, 9, 1, 2, 7 };
        int[] testArray2 = { 4, 5, 9, 0 };
        int[] testResult1 = { 1, 2, 7, 7, 8, 9 };
        int[] testResult2 = { 9, 5, 4, 0 };

        bubbleSort(testArray1, true);
        bubbleSort(testArray2, false);

        if (Arrays.equals(testArray1, testResult1) && Arrays.equals(testArray2, testResult2)) {
            System.out.println("bubbleSort passed");
        } else {
            System.out.println("bubbleSort failed");
        }
    }

}
