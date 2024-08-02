package main
import "fmt"
import "time"

func main(){
	fmt.Println("Namaste\n");
	time.Sleep( 1*time.Second );
	for i:=0;i<10;i++ {
		time.Sleep( 1*time.Second );
		fmt.Println("Namaste\n");
	}
}

