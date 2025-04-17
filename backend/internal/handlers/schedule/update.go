package schedule

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateEvent(ctx echo.Context) error {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "invalid event ID")
	}

	var req Request
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	username := ctx.Get("username").(string)
	chef, err := h.ChefUseCase.GetChefInfo(ctx.Request().Context(), username)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusUnauthorized, "chef not found")
	}

	mark, _ := strconv.Atoi(req.EventDate)
	err = h.ScheduleUseCase.UpdateEventDetails(ctx.Request().Context(), eventID, req.EventName, mark, req.Location, chef.ID)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Event updated successfully")
}
