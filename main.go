package main

import (
	"log"

	"car-rental/internal/handler"
	"car-rental/internal/infrastructure"
	"car-rental/internal/repository"
	"car-rental/internal/router"
	"car-rental/internal/service"

	"github.com/gin-gonic/gin"

	// _ "car-rental/docs"

	"github.com/joho/godotenv"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			FINAL PROJECT MY GRAM DTS
// @version		2.0
// @description	Final Project My Gram
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	// requirement technical:
	// [x] middleware untuk recover ketika panic
	// [x] mengecheck basic auth
	server()
}


func server() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	g := gin.Default()
	g.Use(gin.Recovery())
	gorm := infrastructure.NewGormPostgres()

	customersGroup := g.Group("/customers")
	customerRepo := repository.NewCustomersQuery(gorm)
	customersvc := service.NewCustomerService(customerRepo)
	customerHdl := handler.NewCustomerHandler(customersvc)
	customerRouter := router.NewCustomerRouter(customersGroup, customerHdl)
	customerRouter.Mount()

	carsGroup := g.Group("/cars")
	carRepo := repository.NewCarsQuery(gorm)
	carsvc := service.NewCarservice(carRepo)
	carHdl := handler.NewCarHandler(carsvc)
	carRouter := router.NewCarRouter(carsGroup, carHdl)
	carRouter.Mount()

	driversGroup := g.Group("/drivers")
	driverRepo := repository.NewDriversQuery(gorm)
	driversvc := service.NewDriverservice(driverRepo)
	driverHdl := handler.NewDriverHandler(driversvc)
	driverRouter := router.NewDriverRouter(driversGroup, driverHdl)
	driverRouter.Mount()

	bookingTypesGroup := g.Group("/bookingtypes")
	bookingTypeRepo := repository.NewBookingTypesQuery(gorm)
	bookingTypesvc := service.NewBookingTypeservice(bookingTypeRepo)
	bookingTypeHdl := handler.NewBookingTypeHandler(bookingTypesvc)
	bookingTypeRouter := router.NewBookingTypeRouter(bookingTypesGroup, bookingTypeHdl)
	bookingTypeRouter.Mount()

	driversIncentiveGroup := g.Group("/drivers-incentive")
	driverIncentiveRepo := repository.NewDriversIncentiveQuery(gorm)

	bookingsGroup := g.Group("/bookings")
	bookingRepo := repository.NewBookingsQuery(gorm)
	bookingsvc := service.NewBookingservice(bookingRepo, carRepo, customerRepo, driverRepo, driverIncentiveRepo)
	bookingHdl := handler.NewBookingHandler(bookingsvc, customersvc, carsvc, driversvc, bookingTypesvc)
	bookingRouter := router.NewBookingRouter(bookingsGroup, bookingHdl)
	bookingRouter.Mount()

	
	driversIncentivevc := service.NewDriversIncentiveervice(driverIncentiveRepo)
	driverIncentiveHdl := handler.NewDriverIncentiveHandler(driversIncentivevc, bookingsvc)
	driverIncentiveRouter := router.NewDriverIncentiveRouter(driversIncentiveGroup, driverIncentiveHdl)
	driverIncentiveRouter.Mount()

	

	membershipsGroup := g.Group("/memberships")
	membershipRepo := repository.NewMembershipQuery(gorm)
	membershipsvc := service.NewMembershipservice(membershipRepo)
	membershipHdl := handler.NewMembershipHandler(membershipsvc)
	membershipRouter := router.NewMembershipRouter(membershipsGroup, membershipHdl)
	membershipRouter.Mount()

	
	
	// photosGroup := g.Group("/photos")
	// photoRepo := repository.NewPhotoQuery(gorm) 
	// photoSvc := service.NewPhotoService(photoRepo, userRepo)   
	// photoHdl := handler.NewPhotoHandler(photoSvc)   
	// photoRouter := router.NewPhotoRouter(photosGroup, photoHdl)
	// photoRouter.Mount()
	// commentsGroup := g.Group("/comments")
	// commentRepo := repository.NewCommentQuery(gorm) 
	// commentSvc := service.NewCommentService(commentRepo, userRepo, photoRepo)   
	// commentHdl := handler.NewCommentHandler(commentSvc)   
	// commentRouter := router.NewCommentRouter(commentsGroup, commentHdl)
	// commentRouter.Mount()
	// socialMediasGroup := g.Group("/socialmedias")
	// socialMediaRepo := repository.NewSocialMediaQuery(gorm) 
	// socialMediaSvc := service.NewSocialMediaService(socialMediaRepo, userRepo)   
	// socialMediaHdl := handler.NewSocialMediaHandler(socialMediaSvc)   
	// socialMediaRouter := router.NewSocialMediaRouter(socialMediasGroup, socialMediaHdl)
	// socialMediaRouter.Mount()

	// mount
	
	// swagger
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}

