import java.util.ArrayList;
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
        System.out.println(sumDigitsInNumber(546));

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
}
