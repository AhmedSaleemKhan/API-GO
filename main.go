//main package
package main

//importing diffrent libraries for web requesting and building an API
import (
	"test/router"

	//using an aecho framework
	_ "github.com/lib/pq"
)

//main function
func main() {

	e := router.SetupRoute()

	e.Logger.Fatal(e.Start(":1323")) //request to server
}
