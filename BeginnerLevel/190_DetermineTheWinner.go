// https://www.codechef.com/problems/WINNERR
package beginnerlevel
import ("fmt"
        "math"
)
func main() {
  
  var t int ;
  var pa, pb, qa, qb float64;
  fmt.Scan(&t);

  for t > 0 {

    fmt.Scan(&pa, &pb, &qa, &qb);
    var maxP = math.Max(pa, pb);
    var maxQ = math.Max(qa, qb);

    if(maxP == maxQ) {
      fmt.Println("TIE"); 
    } else if (maxP > maxQ) {
      fmt.Println("Q");
    } else {
      fmt.Println("P");
    }
    t--;
  }
  
  
}
