package main

import (
	"net/http"

	"ibfd.org/app/cfg"
	log "ibfd.org/app/log4u"
	"ibfd.org/app/mem"
	"ibfd.org/app/resource"
	"ibfd.org/app/route"
	admt "ibfd.org/app/table/adm"
	"ibfd.org/app/uc/adm"
)

func main() {
	config := cfg.NewConfig(version)
	defer config.CloseLog()
	log.Infof("Starting %s on %s\n", config.AppName(), config.Server())
	// TODO uncomment the database connection here
	// replace ??? with appropriate name
	// ??? := db.NewDB(config.GetDbDef())
	// defer ???.Close()
	at := admt.NewTable(mem.NewFakeAdminDB())

	pr := resource.NewProtector(adm.New(at), config.BasicAuthRealm())
	rb := route.NewRouteBuilder(config.AllowCORS(), pr, config.AppName(), config.IsLogDebug())
	rb.Add("Home", "GET", "/", resource.HomeHandler(config.HomePage()))
	// Safe
	rb.AddSafe("SafeHome", "GET", "/safe", resource.HomeHandler(config.HomePage()))

	log.Fatal(http.ListenAndServe(config.Server().String(), rb.Router()))
}
