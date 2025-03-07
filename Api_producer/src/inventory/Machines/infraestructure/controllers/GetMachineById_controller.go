package machinecontrollers

import (
	machineusecases "gym-system/src/inventory/Machines/application/useCases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMachineByIdController struct {
	useCase machineusecases.GetMachineById
}

func NewMachineByIdController(useCase machineusecases.GetMachineById) *GetMachineByIdController {
	return &GetMachineByIdController{useCase: useCase}
}

func (getMachine *GetMachineByIdController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Id no valido"})
		return
	}

	machine, err := getMachine.useCase.Execute(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el equipo"})
		return
	}

	if machine == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Maquina no encontrado"})
		return
	}

	g.JSON(http.StatusOK, machine)
}
