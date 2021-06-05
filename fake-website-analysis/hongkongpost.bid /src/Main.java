/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author jerrykwok
 */
import java.util.Scanner;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;

public class Main {
    public static void main(String[] args) throws FileNotFoundException{
        Scanner in = new Scanner(new File("data.txt"));
        ArrayList<CreditCard> cards = new ArrayList<>();
        int count = 0;
        while(in.hasNext()){
            String[] data = in.nextLine().split(", ");
            cards.add(new CreditCard(data[1], data[2], data[3]));
            count++;
        }
        System.out.println("length: " + count);
        System.out.println(CreditCard.genArrayForNum(cards));
        System.out.println(CreditCard.genArrayForExp_date(cards));
        System.out.println(CreditCard.genArrayForCvv(cards));
    }
    
   
}
