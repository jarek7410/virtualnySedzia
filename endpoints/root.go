package endpoints

import (
	"github.com/gin-gonic/gin"
	"log"
	"virtualnySedziaServer/endpoints/controller"
	"virtualnySedziaServer/securiry"
)

type Routs struct {
	//db   *gorm.DB
	r *gin.Engine
	//repo *database.Repo
}

func NewRouts() *Routs {
	//if repo == nil {
	//	//because i can
	//	log.Fatalln("Database needed!!")
	//}
	return &Routs{
		//repo: repo,
		//db:   repo.DB,
		r: gin.Default(),
	}
}
func (r *Routs) AddAuthPaths() {

	authRoutes := r.r.Group("/auth/user")
	{
		// registration route
		authRoutes.POST("/register", controller.Register)
		// login route
		authRoutes.POST("/login", controller.Login)

	}
}
func (r *Routs) ServeApplication() {
	adminRoutes := r.r.Group("/admin")
	{
		adminRoutes.Use(securiry.JWTadminAuth())

		adminRoutes.GET("/users", controller.GetUsers)
		adminRoutes.GET("/user/:id", controller.GetUser)
		adminRoutes.PUT("/user/:id", controller.UpdateUser)

		//maybe in future?
		//adminRoutes.POST("/user/role", controller.CreateRole)
		//adminRoutes.GET("/user/roles", controller.GetRoles)
		//adminRoutes.GET("/user/role/:id", controller.GetRole)
		//adminRoutes.PUT("/user/role/:id", controller.UpdateRole)
	}

	publicRoutes := r.r.Group("/api")
	{

		publicRoutes.GET("/user/:id", controller.GetUser)
		publicRoutes.GET("/issue", controller.GetIssueQuery)
		publicRoutes.GET("/issue/:id", controller.GetIssue)
		publicRoutes.POST("/issue/", controller.PostIssue)
		publicRoutes.GET("/issue/:id/comment", controller.GetIssueComments)
		publicRoutes.GET("/comment/:id", controller.GetComment)
		publicRoutes.POST("/comment", controller.PostComment)
		publicRoutes.GET("/user/actions/:id", controller.GerUserActions)
	}

	protectedRoutes := r.r.Group("/api")
	{
		protectedRoutes.Use(securiry.JWTAuthAnonymous())

		protectedRoutes.DELETE("/issue/:id", controller.DeleteIssue)
		protectedRoutes.DELETE("/comment/:id", controller.DeleteComment)

		protectedRoutes.GET("/users", controller.GetUsersPublic)
		protectedRoutes.GET("/me", controller.GetMyUser)
		protectedRoutes.POST("/me", controller.ChangeMyUser)
		protectedRoutes.POST("/me/password", controller.ChangeMyUserPassword)
	}

}

func (r *Routs) AddPaths() {
	//base := NewBasicForRouts(r.db)

	r.r.GET("/", controller.HalfCheck)
	r.r.GET("/tea", controller.Coffee)
	r.r.GET("/coffee", controller.Coffee)
}

func (r *Routs) Start(port string) {
	portString := ":" + port
	if err := r.r.Run(portString); err != nil {
		log.Fatalln("gin do not start on", port)
	}
}
