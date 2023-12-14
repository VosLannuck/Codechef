package beginnerlevel

// <https://www.codechef.com/problems/SPCP1

import  (
  "fmt"
)
func main() {
  var w, h int; 

  var cw, ch int;
  cw = 60; 
  ch = 130;
  fmt.Scan(&w , &h);

  if(cw <= w && ch >= h) {
    fmt.Println("YES");
  } else {
    fmt.Println("NO");
  }

}
