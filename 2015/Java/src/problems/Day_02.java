package problems;

import utils.Crono;
import utils.FileHandler;

import java.io.IOException;
import java.util.ArrayList;
import java.util.StringTokenizer;

/**
 * Problem 02
 */
public class Day_02 {

    private static ArrayList<String> lines;
    private static Crono crono;

    public static void main(String[] args) throws IOException {
        crono = new Crono();
        crono.start();
        lines = (ArrayList<String>) FileHandler.getNotEmptyLines("../inputfiles/day02.txt");
        System.out.println("[Day 02] File parsed in " + crono.stop().toMillis() + " miliseconds");

        crono.start();
        System.out.print("[Day 02] Problem 1: " + problem_01());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");

        crono.start();
        System.out.print("[Day 02] Problem 2: " + problem_02());
        System.out.println(" (" + crono.stop().toMillis() + " miliseconds)");
    }

    public static int problem_01() throws IOException {
        int result = 0, smallestSide, l, w, h;

        for (String line : lines) {
            StringTokenizer strTok = new StringTokenizer(line, "x\n\r");

            l = Integer.parseInt(strTok.nextToken());
            w = Integer.parseInt(strTok.nextToken());
            h = Integer.parseInt(strTok.nextToken());

            smallestSide = l * w;
            if (w * h < smallestSide) smallestSide = w * h;
            if (h * l < smallestSide) smallestSide = h * l;

            result = result + 2 * l * w + 2 * w * h + 2 * h * l + smallestSide;
        }
        return result;
    }

    public static int problem_02() throws IOException {
        int result = 0, l, w, h;
        int shortestSides[] = new int[2];

        for (String line : lines) {
            StringTokenizer strTok = new StringTokenizer(line, "x\n\r");

            l = Integer.parseInt(strTok.nextToken());
            w = Integer.parseInt(strTok.nextToken());
            h = Integer.parseInt(strTok.nextToken());

            shortestSides[0] = l;
            shortestSides[1] = w;
            if (h < l && l >= w) shortestSides[0] = h;
            if (h < w && w > l) shortestSides[1] = h;

            result += 2 * (shortestSides[0] + shortestSides[1]) + l * w * h;

        }
        return result;

    }


}
