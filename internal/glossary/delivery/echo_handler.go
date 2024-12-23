package delivery

import (
	"github.com/arvaliullin/go-restful-service/internal/domain"
	"github.com/arvaliullin/go-restful-service/internal/glossary/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GlossaryHandler struct {
	UC *usecase.GlossaryUsecase
}

func NewGlossaryHandler(e *echo.Echo, uc *usecase.GlossaryUsecase) {
	handler := &GlossaryHandler{UC: uc}

	e.GET("/terms", handler.GetAll)
	e.GET("/terms/:term", handler.GetTerm)
	e.POST("/terms", handler.CreateTerm)
	e.PUT("/terms/:term", handler.UpdateTerm)
	e.DELETE("/terms/:term", handler.DeleteTerm)
}

// GetAll получает список всех терминов.
// @Summary Получение всех терминов
// @Description Возвращает список всех терминов из глоссария.
// @Tags Глоссарий
// @Produce json
// @Success 200 {array} domain.GlossaryTerm
// @Failure 500 {object} string "Ошибка сервера"
// @Router /terms [get]
func (h *GlossaryHandler) GetAll(c echo.Context) error {
	terms, err := h.UC.GetAllTerms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, terms)
}

// GetTerm получает информацию о конкретном термине.
// @Summary Получение информации о термине
// @Description Возвращает информацию о термине по ключевому слову.
// @Tags Глоссарий
// @Produce json
// @Param term path string true "Ключевое слово"
// @Success 200 {object} domain.GlossaryTerm
// @Failure 404 {object} string "Термин не найден"
// @Router /terms/{term} [get]
func (h *GlossaryHandler) GetTerm(c echo.Context) error {
	term := c.Param("term")
	result, err := h.UC.GetTerm(term)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// CreateTerm добавляет новый термин.
// @Summary Добавление нового термина
// @Description Добавляет новый термин в глоссарий с описанием.
// @Tags Глоссарий
// @Accept json
// @Produce json
// @Param term body domain.GlossaryTerm true "Термин"
// @Success 201 {object} domain.GlossaryTerm
// @Failure 400 {object} string "Неверные данные"
// @Failure 500 {object} string "Ошибка сервера"
// @Router /terms [post]
func (h *GlossaryHandler) CreateTerm(c echo.Context) error {
	var term domain.GlossaryTerm
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.UC.AddTerm(term); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, term)
}

// UpdateTerm обновляет существующий термин.
// @Summary Обновление термина
// @Description Обновляет информацию о термине в глоссарии.
// @Tags Глоссарий
// @Accept json
// @Produce json
// @Param term body domain.GlossaryTerm true "Термин"
// @Success 200 {object} domain.GlossaryTerm
// @Failure 400 {object} string "Неверные данные"
// @Failure 500 {object} string "Ошибка сервера"
// @Router /terms/{term} [put]
func (h *GlossaryHandler) UpdateTerm(c echo.Context) error {
	var term domain.GlossaryTerm
	if err := c.Bind(&term); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.UC.UpdateTerm(term); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, term)
}

// DeleteTerm удаляет термин из глоссария.
// @Summary Удаление термина
// @Description Удаляет термин из глоссария по ключевому слову.
// @Tags Глоссарий
// @Produce json
// @Param term path string true "Ключевое слово"
// @Success 204 "Термин успешно удален"
// @Failure 500 {object} string "Ошибка сервера"
// @Router /terms/{term} [delete]
func (h *GlossaryHandler) DeleteTerm(c echo.Context) error {
	term := c.Param("term")
	if err := h.UC.DeleteTerm(term); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
