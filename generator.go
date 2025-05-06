package main
import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	var id_str string = ""
	var name string = ""

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	
	var id_end bool = false
	for i := 0; i < len(input); i++ {
	  if !id_end {
	    if input[i] == ' ' || input[i] == '.' {
	      id_end = true
	      for i + 1 < len(input) && input[i + 1] == ' ' {
	        i++
	      }
	    } else {
    	  if input[i] < '0' || input[i] > '9' {
    	    fmt.Print("Input error")
    	    return
    	  }
    	  id_str += string(input[i])
	    }
	  } else {
	    name += string(input[i])
	  }
	}
	
	if len(id_str) > 4 {
	   fmt.Print("Input error")
	   return
	}
	
	for len(id_str) < 4 {
	  id_str = "0" + id_str
	}

	var line string = "|"
	line += id_str
	line += "|"
	line += name
	line += "|[Link](https://leetcode.com/problems/"
	para_nums := 0
	var name_leetcode, name_github string = "", ""
	for _, ch := range name {
		if ch >= 'A' && ch <= 'Z' {
			name_leetcode += string(ch - 'A' + 'a')
			name_github += string(ch - 'A' + 'a')
		} else if ch == ' ' {
			name_leetcode += "-"
			name_github += "_"
		} else if ch == '(' {
			para_nums++
		} else if ch == ')' {
			para_nums--
		} else {
			name_leetcode += string(ch)
			name_github += string(ch)
		}
		if para_nums < 0 {
			fmt.Print("Input error")
			return
		}
	}
	if para_nums != 0 {
		fmt.Print("Input error")
		return
	}
	line += name_leetcode
	line += "/description/)|[Code](https://github.com/jerrykcode/leetcode-go/blob/main/Problems/"
	line += id_str
	line += "_"
	line += name_github
	line += ".go)|"
	fmt.Print(line)
}