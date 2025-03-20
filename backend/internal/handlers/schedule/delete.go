package schedule

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteEvent(ctx echo.Context) error {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid event ID")
	}

	username := ctx.Get("username").(string)
	_, err = h.ChefUseCase.GetChefInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "chef not found")
	}

	err = h.ScheduleUseCase.RemoveEvent(ctx.Request().Context(), eventID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Event deleted successfully")
}
