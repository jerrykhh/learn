
import java.util.ArrayList;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**
 *
 * @author jerrykwok
 */
public class CreditCard {

    private String num;
    private String exp_date;
    private String cvv;

    public CreditCard(String num, String exp_date, String cvv) {
        this.num = num;
        this.exp_date = exp_date;
        this.cvv = cvv;
    }

    public void setNum(String num) {
        this.num = num;
    }

    public void setExp_date(String exp_date) {
        this.exp_date = exp_date;
    }

    public void setCvv(String cvv) {
        this.cvv = cvv;
    }



    public static String genArrayForNum(ArrayList<CreditCard> list) {
        String result = "[";
        for (CreditCard card : list) {
            result += "'" + card.num + "', ";
        }
        result = result.substring(0, result.length() - 2);
        result += "]";
        return result;
    }

    public static String genArrayForExp_date(ArrayList<CreditCard> list) {
        String result = "[";
        for (CreditCard card : list) {
            result += "'" + card.exp_date + "', ";
        }
        result = result.substring(0, result.length() - 2);
        result += "]";
        return result;
    }
    
    
    public static String genArrayForCvv(ArrayList<CreditCard> list) {
        String result = "[";
        for (CreditCard card : list) {
            result += "'" + card.cvv + "', ";
        }
        result = result.substring(0, result.length() - 2);
        result += "]";
        return result;
    }
    
    
}
