package main
import("fmt")

func main(){
	start := 0
	end := 0

	fmt.Print("Enter start number: ")
	fmt.Scanln(&start)
	fmt.Print("Enter end number: ")
	fmt.Scanln(&end)

	for j := start; j <= end; j++ {
		for i := 1; i <= 10; i++ {
	       fmt.Println(j, "x", i, "=", i*j)	
	    }
	    fmt.Println()		
	}
	
}