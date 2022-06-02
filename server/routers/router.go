package routers

import (
	"github.com/astaxie/beego"
	"tie-explorer-go/controllers"
)

func init() {
	// Default url
	beego.Router("/", &controllers.DefaultController{}, "*:Default")

	api := beego.NewNamespace("/api",
		// Dashboard API
		beego.NSRouter("/dashboard", &controllers.DashboardController{}, "*:Info"),
		beego.NSRouter("/search", &controllers.DashboardController{}, "*:Search"),

		// Block API
		beego.NSRouter("/blocks", &controllers.BlockController{}, "*:List"),
		beego.NSRouter("/block/:number", &controllers.BlockController{}, "*:Info"),

		// Transaction API
		beego.NSRouter("/transactions", &controllers.TransactionController{}, "*:List"),
		beego.NSRouter("/transaction/:hash", &controllers.TransactionController{}, "*:Info"),

		// Token API
		beego.NSRouter("/tokens", &controllers.TokenController{}, "*:List"),
		beego.NSRouter("/token/:address", &controllers.TokenController{}, "*:Info"),

		// Address API
		beego.NSRouter("/address/:address", &controllers.AddressController{}, "*:Info"),
	)

	bridge := beego.NewNamespace("/bridge",
		// Dashboard API
		beego.NSRouter("/chain_health", &controllers.BridgeController{}, "*:ChainHealth"),
		beego.NSRouter("/token_basics", &controllers.BridgeController{}, "*:TokenBasics"),
		beego.NSRouter("/token_map", &controllers.BridgeController{}, "*:TokenMap"),
		beego.NSRouter("/get_fee", &controllers.BridgeController{}, "*:GetFee"),
		beego.NSRouter("/expect_time", &controllers.BridgeController{}, "*:ExpectTime"),
		beego.NSRouter("/transaction_of_hash", &controllers.BridgeController{}, "*:Transaction"),
		beego.NSRouter("/transactions_of_address", &controllers.BridgeController{}, "*:Transactions"),
		beego.NSRouter("/set_exchange", &controllers.BridgeController{}, "*:SetExchange"),
	)

	beego.AddNamespace(api, bridge)
}
