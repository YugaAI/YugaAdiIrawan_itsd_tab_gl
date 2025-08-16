package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) GetEmployeByNIK(c *gin.Context) {
	nik := c.Param("nik")

	employee, err := h.service.GetEmployeeByNIK(nik)
	if err != nil {
		c.JSON(404, ResponseMessage{
			Status:  false,
			Message: "nik :" + nik + " tidak di temukan",
		})
		return
	}

	c.JSON(200, ResponseMessage{Status: true, Data: employee})
}
