package beginnerlevel
// https://www.codechef.com/problems/SPCP2
import ("fmt" 
  "math")

func main() {
  var t int;
  var n, x float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &n);
    maxCap := 100 * x;
    if (maxCap > n ){
      fmt.Println(0);
    } else {
      if maxCap > n {
      
        needPlanes := math.Ceil((maxCap - n) / 100);
        fmt.Println(needPlanes);

      } else {
        needPlanes := math.Ceil((n - maxCap) / 100);
        fmt.Println(needPlanes);
      }

    }
    t--;
  }
}
