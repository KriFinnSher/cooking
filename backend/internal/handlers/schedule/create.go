package schedule

import (
	"cooking/backend/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Request struct {
	EventName string `json:"event_name"`
	EventDate string `json:"event_date"`
	Location  string `json:"location"`
}

type Response struct {
	ID        int    `json:"id"`
	EventName string `json:"event_name"`
	EventDate string `json:"event_date"`
	Location  string `json:"location"`
	ChefID    int    `json:"chef_id"`
}

type Handler struct {
	ScheduleUseCase *usecase.ScheduleUseCase
	ChefUseCase     *usecase.ChefUseCase
}

func NewScheduleHandler(scheduleUseCase *usecase.ScheduleUseCase, chefUseCase *usecase.ChefUseCase) *Handler {
	return &Handler{ScheduleUseCase: scheduleUseCase, ChefUseCase: chefUseCase}
}

func (h *Handler) CreateEvent(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	username := ctx.Get("username").(string)
	chef, err := h.ChefUseCase.GetChefInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "chef not found")
	}

	mark, _ := strconv.Atoi(req.EventDate)
	err = h.ScheduleUseCase.CreateEvent(ctx.Request().Context(), req.EventName, mark, req.Location, chef.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "Event created successfully")
}
