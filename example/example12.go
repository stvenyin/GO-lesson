import (
	"fmt"
	"runtime"
)

func main() {

	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello.%s\n", name)
		}(name)
		runtime.Gosched()
	}

}
