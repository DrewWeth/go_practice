import java.util.*;
public class CountBits{
  public static void p(int x){
    System.out.println(x);
  }

  public static int countBitsEvenFaster(int x){
    // count bits of each 2-bit chunk
    x = x - ((x >> 1) & 0x55555555);
    p(x);

    // count bits of each 4-bit chunk
    x = (x & 0x33333333) + ((x >> 2) & 0x33333333);
    p(x);

    // count bits of each 8-bit chunk
    x = x + (x >> 4);
    p(x);
    // mask out junk
    x &= 0xF0F0F0F;
    // add all four 8-bit chunks
    return (x * 0x01010101) >> 24;
  }

  public static void main(String[] args){
    System.out.println(countBitsEvenFaster(27));
  }

}
