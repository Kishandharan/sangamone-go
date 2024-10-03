package main
import("fmt")

func main(){
	for j := 3; j <= 20; j++ {
		for i := 1; i <= 10; i++ {
	       fmt.Println(j, "x", i, "=", i*j)	
	    }
	    fmt.Println()		
	}
	
}