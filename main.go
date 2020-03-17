/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import app "TgoShorten/application"

func main() {
	a := app.NewApp()
	a.Run("127.0.0.1:8080")
}
