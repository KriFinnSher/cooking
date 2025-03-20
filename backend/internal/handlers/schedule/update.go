package schedule

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateEvent(ctx echo.Context) error {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid event ID")
	}

	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	username := ctx.Get("username").(string)
	chef, err := h.ChefUseCase.GetChefInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "chef not found")
	}

	err = h.ScheduleUseCase.UpdateEventDetails(ctx.Request().Context(), eventID, req.EventName, req.EventDate, req.Location, chef.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Event updated successfully")
}
