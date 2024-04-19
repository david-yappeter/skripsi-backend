package main

import (
	"myapp/cmd"

	goShopee "github.com/david-yappeter/go-shopee-v2"
)

// @title		Mortal Health - Clinic Pilot API
// @version	0.0.1
// @host		cp-api.mortalhealth.com
// @BasePath	/
func main() {
	goShopee.NewClient(goShopee.App{
		PartnerID:   0,
		PartnerKey:  "",
		RedirectURL: "",
		APIURL:      "",
		Client:      &goShopee.Client{},
	}, goShopee.WithRetry(3), goShopee.WithLogger(&goShopee.LeveledLogger{
		Level: goShopee.LevelDebug,
	}))

	return

	cmd.Execute()
}
