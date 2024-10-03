package main
import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	str1 := ""
	f1, err := os.OpenFile("table.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil{return}
	start := 0
	end := 0

	fmt.Print("Enter start number: ")
	fmt.Scanln(&start)
	fmt.Print("Enter end number: ")
	fmt.Scanln(&end)

	for j := start; j <= end; j++ {
		for i := 1; i <= 10; i++ {
	       str1 += (strconv.Itoa(j)+"x"+strconv.Itoa(i)+"="+strconv.Itoa((i*j)))+"\n"	
	    }
	    str1 += "\n"		
	}

	f1.Write([]byte(str1))

	f1.Close()
	
}