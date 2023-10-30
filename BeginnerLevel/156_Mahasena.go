package beginnerlevel
// https://www.codechef.com/problems/AMR15A
import "fmt"
func main () {


  var n, a int;
  fmt.Scan(&n);
  var oddsTotal, evenTotal int;

  for n > 0 {

    fmt.Scan(&a);
    if(a % 2 == 0 ) {
      evenTotal += 1;
    }else {
      oddsTotal += 1;
    }

    n--;  
  }

  if(oddsTotal > evenTotal) {

    fmt.Println("NOT READY");
  } else {
    fmt.Println("READY FOR BATTLE");
  }
  

  
  
}

