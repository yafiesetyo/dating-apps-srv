package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yafiesetyo/dating-apps-srv/config"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/cache"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/db"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories"
	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/usecase"
	"github.com/yafiesetyo/dating-apps-srv/internal/bcyrpt"
	"github.com/yafiesetyo/dating-apps-srv/internal/handler"
	"github.com/yafiesetyo/dating-apps-srv/internal/logger"

	jwtware "github.com/gofiber/contrib/jwt"
)

func init() {
	config.Init()

	logger.Init("local")
}

func main() {
	db := db.New(config.Cfg)
	cache := cache.New(config.Cfg)
	validator := validator.New()
	p := bcyrpt.New()

	rspu := repositories.NewSetPremiumUser(db)
	ruc := repositories.NewUserCreator(db)
	rufbi := repositories.NewUserFinderById(db)
	rufbu := repositories.NewUserFinderByUsername(db)
	rupf := repositories.NewUserProfileFinder(db)
	rcus := repositories.NewUserSwipeCreator(db)

	ucdpg := usecase.NewDatingProfileGetter(cache, rufbi, rupf)
	ucupg := usecase.NewUserProfileGetter(rufbi)
	ucl := usecase.NewLogin(rufbu, p)
	ucr := usecase.NewRegister(rufbu, ruc, p)
	ucps := usecase.NewSwipeProfile(cache, rufbi, rcus, rupf)
	ucp := usecase.NewPurchase(rufbi, rspu)

	handler := handler.NewHandler(
		validator,
		ucdpg,
		ucupg,
		ucl,
		ucr,
		ucps,
		ucp,
	)

	app := fiber.New()
	v1 := app.Group("/v1")

	v1.Post("/login", handler.Login)
	v1.Post("/register", handler.Register)
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(config.Cfg.JWT.SecretKey),
			JWTAlg: jwtware.HS256,
		},
		SuccessHandler: successHandler,
	}))
	v1.Get("/profile", handler.ViewUserProfile)
	v1.Get("/dating", handler.ViewDatingProfile)
	v1.Post("/swipe/:id", handler.Swipe)
	v1.Post("/purchase", handler.Purchase)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":" + config.Cfg.Port); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}

func successHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(http.StatusUnauthorized).
			JSON(fiber.Map{
				"message": "unauthorized",
				"error":   "failed to parse claim",
			})
	}

	c.Locals("id", id)
	return c.Next()
}
