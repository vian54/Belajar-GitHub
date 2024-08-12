package service

import (
    "github.com/ricnah/workit-be/config"
    "github.com/ricnah/workit-be/service/middlewares"
    "github.com/ricnah/workit-be/service/modules/product"
    "github.com/ricnah/workit-be/service/modules/product/handler" // Tambahkan impor handler
    "github.com/ricnah/workit-be/service/modules/product/repository"
    "github.com/ricnah/workit-be/service/modules/product/usecase"
    userModule "github.com/ricnah/workit-be/service/modules/user"
    userrepo "github.com/ricnah/workit-be/service/modules/user/repository"
    usercase "github.com/ricnah/workit-be/service/modules/user/usecase"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func setRoutes(cfg *config.Config) (r *gin.Engine, err error) {
    r = gin.New()

    
    userRepo := userrepo.UserCreateRepository(cfg.DB)
    userCase := usercase.UserCreateUsecase(userRepo)

    productRepo := repository.NewProductRepository(cfg.DB)
    productUsecase := usecase.NewProductUsecase(productRepo)
    productHandler := handler.NewProductHandler(productUsecase)

    r.Use(corsConfig())
    r.Use(middlewares.ActivityLogger())

    api := r.Group("/api")
    v1 := api.Group("/v1")
    {
        userModule.InitRoutes(v1, userCase, cfg)
        
        product.InitRoutes(v1, productHandler) // Gunakan productHandler
    }

    r.GET("/", func(c *gin.Context) {
        c.Redirect(302, "/public")
    })

    r.Static("/public", "./view/.output/public")

    r.NoRoute(func(c *gin.Context) {
        c.File("./view/.output/public/index.html")
    })

    return
}

func corsConfig() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    })
}

func Start(cfg *config.Config) (err error) {
    eng, err := setRoutes(cfg)
    if err != nil {
        return err
    }

    eng.Run(cfg.App.Host + ":" + cfg.App.Port)
    return
}
