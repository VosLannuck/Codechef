package beginnerlevel
// https://www.codechef.com/problems/MAXTASTE
import ("fmt"
      "math")


func main() {

  var t int;
  var a,b,c,d float64;
  fmt.Scan(&t);
  
  for t > 0 {
    fmt.Scan(&a, &b, &c ,&d);
    var max_dish_1 float64 = math.Max(a, b);
    var max_dish_2 float64 = math.Max(c, d);

    fmt.Println(max_dish_1 + max_dish_2);
    t--;
  }
  
  



  
}
