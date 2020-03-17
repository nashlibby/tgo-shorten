/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import (
	app "TgoShorten/application"
	"log"
)

func main() {
	a := app.NewApp()
	log.Printf("App run at 127.0.0.1:8080")
	a.Run("127.0.0.1:8080")
}
