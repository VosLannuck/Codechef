package begineerLevel;

import ("fmt")
// https://www.codechef.com/problems/LASTLEVELS
func main() {
  
  var x, y,z int;
  var t int;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y, &z);
    totalCompletion := x * y;
    takingBreak := x / 3 ;
    if (x % 3 == 0 ) {
      fmt.Println(totalCompletion + (takingBreak - 1) * z);
    } else {

      fmt.Println(totalCompletion + takingBreak * z);
    }

    t--;
  }
}
