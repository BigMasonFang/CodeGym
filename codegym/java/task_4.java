import java.io.*;

public class task_4 {
    public static void main(String[] args) {
        checkSeason(12);
        checkSeason(4);
        checkSeason(7);
        checkSeason(10);
        // try {
        // triangleCheck();
        // } catch (Exception e) {
        // //TODO: handle exception
        // };
        System.out.println(6 % 5);

    }

    public static void checkSeason(int s) {
        switch (s) {
            case 12:
            case 1:
            case 2:
                System.out.println("冬季");
                break;
            case 3:
            case 4:
            case 5:
                System.out.println("春季");
                break;
            case 6:
            case 7:
            case 8:
                System.out.println("夏季");
                break;
            case 9:
            case 10:
            case 11:
                System.out.println("秋季");
                break;
            default:
                break;
        }
    }

    public static void triangleCheck() throws Exception {
        // 在此编写你的代码
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in)); // 输入流
        String sA = reader.readLine();
        String sB = reader.readLine();
        String sC = reader.readLine();
        Integer a = Integer.parseInt(String.valueOf(sA));
        Integer b = Integer.parseInt(String.valueOf(sB));
        Integer c = Integer.parseInt(String.valueOf(sC));
        if (((a + b) > c) && ((a + c) > b) && ((b + c) > a)) {
            System.out.println("三角形可能存在。");
        } else {
            System.out.println("三角形不可能存在。");
        }
    }

    public static void trafficLight() throws Exception {
        // 在此编写你的代码
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in)); // 输入流
        String sTime = reader.readLine();
        double t = Double.parseDouble(String.valueOf(sTime));
        double tModed = t % 5;

        if (tModed >= 0 & tModed < 3) {
            System.out.println("绿色");
        } else if (tModed >= 3 & tModed < 4) {
            System.out.println("黄色");
        } else {
            System.out.println("红色");
        }
    }
}
