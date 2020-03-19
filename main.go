/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import (
	"flag"
	"log"
	app "tgo-shorten/application"
)

func main() {
	a := app.NewApp()
	var port string
	flag.StringVar(&port, "port", "8080", "web port")
	flag.Parse()
	log.Printf("App run at 127.0.0.1:" + port)
	a.Run("127.0.0.1:" + port)
}
