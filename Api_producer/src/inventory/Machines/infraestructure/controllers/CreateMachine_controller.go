package machinecontrollers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gym-system/src/inventory/Machines/application/useCases"
)

type CreateMachineController struct {
    useCase machineusecases.CreateMachine
}

func NewCreateMachineController(useCase machineusecases.CreateMachine) *CreateMachineController {
    return &CreateMachineController{useCase: useCase}
}

func (ce_c *CreateMachineController) Execute(c *gin.Context) {

    
    var requestBody struct {
        Name string `json:"name"`
        Type string `json:"type"`
        Status string `json:"status"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":"Datos invalidos"})
        return
    }

    ce_c.useCase.Execute(requestBody.Name, requestBody.Type, requestBody.Status)

    c.JSON(http.StatusOK, gin.H{"message": "Maquina agregado exitosamente"})
}