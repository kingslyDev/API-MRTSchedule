package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kingslyDev/API-MRTSchedule/modules/stations"
)

func main()  {
	
	initiateRouter()
	
}

func initiateRouter(){
	var {
		router = gin.Default();
		api = router.Group("v1/api")
	}
	

	stations.Initiate(api)

	router.Run(":8080");
}