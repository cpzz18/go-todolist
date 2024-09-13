package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

)

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUpdateTodosController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request UpdateRequest
		if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		_, err := db.Exec(
			"UPDATE todos SET title = ?, description = ? WHERE id = ?",
			request.Title,
			request.Description,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.NoContent(http.StatusNoContent)
	})
}
