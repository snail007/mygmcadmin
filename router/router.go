package router

import (
	"github.com/snail007/gmc"
	"mygmcadmin/controller"
)

func InitRouter(s *gmc.HTTPServer) {
	// acquire router object
	r := s.Router()

	// bind a controller, /demo is path of controller, after this you can visit http://127.0.0.1:7080/demo/hello
	// "hello" is full lower case name of controller method.
	//r.Controller("/demo", new(controller.Demo))
	//r.ControllerMethod("/", new(controller.Demo), "Index__")
	//r.ControllerMethod("/index.html", new(controller.Demo), "Index__")

	r.ControllerMethod("/", new(controller.Login), "Index_")
	r.Controller("/login", new(controller.Login))
	r.Controller("/main", new(controller.Main))
	r.Controller("/user", new(controller.User))

	// indicates router initialized
	s.Logger().Infof("router inited.")
}
