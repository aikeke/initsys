package main
import(
	"fmt"
	"os"
	"initsys"
)

func main(){
	 nt, ip := os.Args[1], os.Args[2]
        s := initsys.NetInit(nt, ip)
        fmt.Println(s)
}
