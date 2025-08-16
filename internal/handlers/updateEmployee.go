package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	nik := c.Param("nik")

	var req struct {
		Nama string `json:"nama"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, ResponseMessage{
			Status:  false,
			Message: "Invalid request",
		})
		return
	}

	oldName, err := h.service.UpdateEmployeeByName(nik, req.Nama)
	if err != nil {
		c.JSON(404, ResponseMessage{Status: false,
			Message: fmt.Sprintf("Nik: %s tidak ditemukan", nik)})
		return
	}
	c.JSON(200, ResponseSuccessUpdate{
		Message: fmt.Sprintf("Perubahan Nama: %s menjadi %s Berhasil", oldName, req.Nama),
	})
}
