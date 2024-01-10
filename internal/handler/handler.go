package handler

import (
	"fmt"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/handler/request"
	"github.com/yafiesetyo/dating-apps-srv/internal/handler/response"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
)

type Handler struct {
	validator           *validator.Validate
	datingProfileGetter interfaces.UViewDatingProfile
	userProfileGetter   interfaces.UViewProfile
	loginner            interfaces.ULogin
	registerer          interfaces.UCreateUser
	profileSwiper       interfaces.USwipe
	purchaser           interfaces.UPurchase
}

func NewHandler(
	validator *validator.Validate,
	datingProfileGetter interfaces.UViewDatingProfile,
	userProfileGetter interfaces.UViewProfile,
	login interfaces.ULogin,
	register interfaces.UCreateUser,
	profileSwiper interfaces.USwipe,
	purchaser interfaces.UPurchase,
) *Handler {
	return &Handler{
		datingProfileGetter: datingProfileGetter,
		userProfileGetter:   userProfileGetter,
		loginner:            login,
		registerer:          register,
		profileSwiper:       profileSwiper,
		purchaser:           purchaser,
		validator:           validator,
	}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var (
		req request.Login
		ctx = c.Context()
	)

	if err := c.BodyParser(&req); err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	if err := h.validator.StructCtx(ctx, &req); err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	token, err := h.loginner.Do(ctx, entity.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var (
		req request.Register
		ctx = c.Context()
	)

	if err := c.BodyParser(&req); err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	if err := h.validator.StructCtx(ctx, &req); err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	if err := h.registerer.Do(ctx, req.ToEntity()); err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Created",
	})
}

func (h *Handler) ViewUserProfile(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		id  = c.Locals("id").(float64)
	)

	user, err := h.userProfileGetter.Do(ctx, entity.User{
		ID: uint(id),
	})
	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	response := response.Profile{
		ID:           user.ID,
		Name:         user.Name,
		Gender:       string(user.Gender),
		ImageUrl:     user.ImageUrl,
		DOB:          user.DOB,
		POB:          user.POB,
		Religion:     string(user.Religion),
		Description:  user.Description,
		Hobby:        user.Hobby,
		VerifiedUser: user.VerifiedUser,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *Handler) ViewDatingProfile(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		id  = c.Locals("id").(float64)
	)

	user, err := h.datingProfileGetter.Do(ctx, entity.User{
		ID: uint(id),
	})
	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	response := response.Profile{
		ID:           user.ID,
		Name:         user.Name,
		Gender:       string(user.Gender),
		ImageUrl:     user.ImageUrl,
		DOB:          user.DOB,
		POB:          user.POB,
		Religion:     string(user.Religion),
		Description:  user.Description,
		Hobby:        user.Hobby,
		VerifiedUser: user.VerifiedUser,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *Handler) Swipe(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		id  = c.Locals("id").(float64)
	)

	action := c.Query("action")
	if constants.SwipeAction(action) == "" {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"message": "invalid action",
			})
	}

	likedUserId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	if err := h.profileSwiper.Do(ctx, uint(id), uint(likedUserId), constants.SwipeAction(action)); err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Ok",
	})
}

func (h *Handler) Purchase(c *fiber.Ctx) error {
	var (
		ctx = c.Context()
		id  = c.Locals("id").(float64)
	)

	feature := constants.GetFeature(c.Query("feature"))
	fmt.Println(feature)
	if feature < 1 {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"message": "invalid feature",
			})
	}

	user := entity.User{
		ID: uint(id),
	}
	if feature == constants.VerifiedUser {
		user.VerifiedUser = true
	} else if feature == constants.UnlimitedSwipe {
		user.UnlimitedSwipe = true
	}

	if err := h.purchaser.Do(ctx, user); err != nil {
		return c.
			Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Ok",
	})
}
