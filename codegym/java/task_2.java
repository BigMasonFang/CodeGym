// package codegym.java;

// cal the circumference
public class task_2 {
    public static void main(String[] args) {
        printCircleCircumference(5);
        System.out.println(convertCelsiusToFahrenheit(41));
        Person person = new Person();
        Man man = new Man();
        Woman woman = new Woman();
        man.wife = woman;
        woman.husband = man;
    }

    public static void printCircleCircumference(int radius) {
        double pi = 3.14;
        System.out.println(2 * pi * radius);
    }

    public static double convertCelsiusToFahrenheit(int celsius) {
        return celsius * 9 / 5 + 32;
    }

    public static class Person {
        int age;
        int weight;
        int money;
        int height;
        String name;
    }

    public static class Man extends Person {
        public Woman wife;
    }

    public static class Woman extends Person {
        public Man husband;
    }
}
