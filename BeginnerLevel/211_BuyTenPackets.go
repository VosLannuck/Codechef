package beginnerlevel

import "fmt"

func main() {

  var t int;
  var x , y float64;
  fmt.Scan(&t);
  for t > 0 {
    fmt.Scan(&x, &y);
    var yPerCigarrete = y / 4;
    var xPerCigarrete = x / 2;

    if(yPerCigarrete < xPerCigarrete) {
      fmt.Println(y * 2 + (x));
    } else {
      fmt.Println(3 * x + (y));
    }
    t--;
  }
}
