package beginnerlevel

// https://www.codechef.com/problems/FLIPCARDS
import ("fmt"
      "math"
)

func main() {

  var t int;
  var n , x float64;

  fmt.Scan(&t);

  for t > 0 {
    fmt.Scan(&n, &x);
    if(x ==0 || n ==0 || x == n ) {
      fmt.Println(0);
    } else {
      var faceDown float64 = n - x;
      var faceUp float64 = n-faceDown;
      var min = math.Min(faceDown, faceUp);
      fmt.Println(min);
    }
    t--;
  }

}
