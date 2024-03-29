package routes

import (
	"distriai-index-solana/handlers"
	"distriai-index-solana/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	mailbox := engine.Group("/mailbox")
	{
		mailbox.POST("/subscribe", handlers.Subscribe)
		mailbox.POST("/unsubscribe", handlers.Unsubscribe)
	}
	user := engine.Group("/user")
	{
		user.POST("/login", handlers.Login)
	}
	machine := engine.Group("/machine")
	{
		machine.POST("/filter", handlers.MachineFilter)
		machine.POST("/market", handlers.MachineMarket)
		machine.POST("/mine", handlers.MachineMine)
	}
	order := engine.Group("/order")
	{
		order.POST("/mine", handlers.OrderMine)
		order.POST("/all", handlers.OrderAll)
	}
	reward := engine.Group("/reward")
	{
		reward.POST("/total", handlers.RewardTotal)
		reward.POST("/claimable/list", handlers.RewardClaimableList)
		reward.POST("/period/list", handlers.RewardPeriodList)
		reward.POST("/machine/list", handlers.RewardMachineList)
	}
	log := engine.Group("/log")
	{
		log.POST("/add", handlers.LogAdd)
		log.POST("/list", handlers.LogList)
	}
	model := engine.Group("/model")
	{
		model.POST("/list", handlers.ModelList)
		model.GET("/:Id", handlers.ModelGet)
	}
	modelAuth := engine.Group("/model", middleware.Jwt())
	{
		modelAuth.POST("/create", handlers.ModelCreate)
		modelAuth.POST("/presign", handlers.ModelPresign)
	}
	engine.POST("/faucet", handlers.Faucet)
}
