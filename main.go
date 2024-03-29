package main

import (
	"github.com/snail007/gmc"
	gcore "github.com/snail007/gmc/core"
	"mygmcadmin/initialize"
)

func main() {

	// 1. create an default app to run.
	app := gmc.New.AppDefault()

	// 2. add a http server service to app.
	app.AddService(gcore.ServiceItem{
		Service: gmc.New.HTTPServer(app.Ctx()).(*gmc.HTTPServer),
		AfterInit: func(s *gcore.ServiceItem) (err error) {
			// do some initialize after http server initialized.
			err = initialize.Initialize(s.Service.(*gmc.HTTPServer))
			return
		},
	})

	// 3. run the app
	if e := gmc.Err.Stack(app.Run()); e != "" {
		app.Logger().Fatal(e)
	}
}
