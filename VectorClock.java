import java.util.*;


public class VectorClock{
  public HashMap<String,Integer> vc;
  public VectorClock(){
    this.vc = new HashMap<String, Integer>();
  }

  public static void main(String [] args){
    HashMap<String, Integer> vc1 = new HashMap<String, Integer>();
    vc1.put("a", 4);
    vc1.put("b", 5);
    HashMap<String, Integer> vc2 = new HashMap<String, Integer>();
    vc2.put("a", 4);
    vc2.put("b", 2);
    vc2.put("c", 1);

    int result = compareVectorClock(vc1, vc2);
    System.out.println(result);
  }


  public static int compareVectorClock(HashMap<String,Integer> vc1, HashMap<String,Integer> vc2){
    int tempResult = 0;
    boolean inComp = false;
    boolean isFirst = true;

    int c1;
    int c2;

    HashSet<String> keys = new HashSet(vc1.keySet());
    keys.addAll(vc2.keySet());

    for(String key : keys) {
      System.out.println(key);
      if(vc1.containsKey(key)){
        c1 = vc1.get(key);
      }else{
        c1 = 0;
      }
      if(vc2.containsKey(key)){
        c2 = vc2.get(key);
      }else{
        c2 = 0;
      }
      if(tempResult == 0){
        if(c1 > c2){
          tempResult = 1; 
        }else if(c2 > c1){
          tempResult = -1;
        }
      }else{
        if(c1 > c2 && tempResult == -1){
          return 2;
        }else if(c2 > c1 && tempResult == 1){
          return 2;
        }
      }
    }
    return tempResult;
  }
}
