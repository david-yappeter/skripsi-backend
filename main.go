package main

import (
	"myapp/cmd"
)

// @title		Mortal Health - Clinic Pilot API
// @version	0.0.1
// @host		cp-api.mortalhealth.com
// @BasePath	/
func main() {

	// client, err := gotiktok.New(
	// 	"6bgdj91pdnm9v",
	// 	"6cfbf8374b80618ca0d1b5eafd87ca0e23554e57",
	// 	"202309",
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.SearchProducts(
	// 	context.Background(),
	// 	gotiktok.CommonParam{
	// 		AccessToken: "ROW_T7RD0wAAAAAjuAWBFh7OoD5X1P3Y_MSshmDMZG4rcBi48ay1C4wMwLRWz0ZiBR2yKAPor8EMpLUvnjsWFSTx4fWFn8Me9B2B1RaFihkHEQtnSHt9KemECjG16XnJ3yWL9EICvihPoLIihfjDZyIl0Nb56kf-GMr9lf5Pq7qBN2VaEFq_8fsZVg",
	// 		ShopCipher:  "ROW_ij-EHgAAAAAFH7_LWApa2DADTZh6ANIA",
	// 		ShopId:      "7495591168837323491",
	// 	},
	// 	gotiktok.CursorPaginationParam{
	// 		PageSize: 50,
	// 	},
	// 	gotiktok.SearchProductRequest{
	// 		Status: nil,
	// 	},
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// t, _ := json.MarshalIndent(resp, " ", "  ")
	// fmt.Printf("%+v\n", string(t))

	// for _, category := range resp.Categories {
	// 	if category.ParentId != "0" {
	// 		fmt.Printf("%+v\n", category)
	// 	}
	// }
	// return

	cmd.Execute()
}
