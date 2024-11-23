package utility

import "fmt"

func ResultFormat(status int,resp string){
	fmt.Println("\n")
	fmt.Println("STATUS",status)
	fmt.Println("\n")
	fmt.Println(" RESPONSE: ",string(resp))
	fmt.Println("\n")

}