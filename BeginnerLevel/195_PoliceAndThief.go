package beginnerlevel
// https://www.codechef.com/problems/POLTHIEF
import ("fmt"
        "math"
)

func main() {
  
  var t int ;
  var x, y float64;

  fmt.Scan(&t);


  for t > 0 {
    fmt.Scan(&x, &y);

    if(x == y ) {

      fmt.Println(0);
    } else if (x < y ) {
      fmt.Println(math.Ceil((math.Abs(x - y) / 2) * 2 ) );
    } else {
      fmt.Println( math.Ceil( math.Abs(x-y)  ));
    }
  
    t--;
  }

}
