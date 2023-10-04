package beginnerlevel;

import "fmt";
func main() {

var str  string = "";
  fmt.Scan(&str);

  for i:= 0 ; i < len(str); i++ {
    if i == 6 {
      fmt.Printf("%c", str[i]);
    }
  }

}
