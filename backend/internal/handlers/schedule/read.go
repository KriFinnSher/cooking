package schedule

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GetEvent(ctx echo.Context) error {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid event ID")
	}

	eventDetails, err := h.ScheduleUseCase.GetEventDetails(ctx.Request().Context(), eventID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "event not found")
	}
	mark := strconv.Itoa(eventDetails.EventDate)

	response := Response{
		ID:        eventDetails.ID,
		EventName: eventDetails.EventName,
		EventDate: mark,
		Location:  eventDetails.Location,
		ChefID:    eventDetails.ChefID,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllEvents(ctx echo.Context) error {

	events, err := h.ScheduleUseCase.GetAllEvents(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	var response []Response

	for _, event := range events {
		mark := strconv.Itoa(event.EventDate)
		response = append(response, Response{
			ID:        event.ID,
			EventName: event.EventName,
			EventDate: mark,
			Location:  event.Location,
			ChefID:    event.ChefID,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}
