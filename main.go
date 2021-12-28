package main
import "fmt"
import "time"
import "github.com/karinahallberg/forking_time/lctime"


func main()  {
	lctime.SetLocale("fi_FI")
	fmt.Println(lctime.Strftime("%e %B, %Y", time.Now()))
}