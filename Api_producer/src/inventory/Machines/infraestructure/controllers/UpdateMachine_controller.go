package machinecontrollers

import (
	"gym-system/src/inventory/Machines/application/useCases"
	"gym-system/src/inventory/Machines/infraestructure/adapters"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMachineController struct {
	useCase            machineusecases.UpdateMachine
	rabbitMQProducer   *adapters.RabbitMQProducer
}

func NewUpdateMachineController(useCase machineusecases.UpdateMachine, rabbitMQProducer *adapters.RabbitMQProducer) *UpdateMachineController {
	return &UpdateMachineController{
		useCase:          useCase,
		rabbitMQProducer: rabbitMQProducer,
	}
}

func (updateMachine *UpdateMachineController) Execute(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Id de máquina inválido"})
		return
	}

	var input struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Status string `json:"status"`
	}

	if err := g.ShouldBindJSON(&input); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateMachine.useCase.Execute(id, input.Name, input.Type, input.Status)

	err = updateMachine.rabbitMQProducer.PublishMachineStatus(id, input.Status)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el mensaje a RabbitMQ"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Máquina editada con éxito y mensaje enviado a RabbitMQ"})
}