package beginnerlevel;


import "fmt"
func main() {
  var t int;
  var a, b int;
  fmt.Scan(&t);
  
  for i:= 0 ; i<t; i++ {
    fmt.Scan(&a, &b);

    if a + b <= 6 {
      fmt.Println("NO");
    }else {
      fmt.Println("YES");  
    }
  }
}
