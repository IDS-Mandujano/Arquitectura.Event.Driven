package equipamentcontrollers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gym-system/src/inventory/Equipments/application/useCases"
)

type CreateEquipamentController struct {
    useCase equipamentusecases.CreateEquipament
}

func NewCreateEquipamentController(useCase equipamentusecases.CreateEquipament) *CreateEquipamentController {
    return &CreateEquipamentController{useCase: useCase}
}

func (ce_c *CreateEquipamentController) Execute(c *gin.Context) {

    var requestBody struct {
        Name string `json:"name"`
        Category string `json:"category"`
        Condition string `json:"condition"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":"Datos invalidos"})
        return
    }

    ce_c.useCase.Execute(requestBody.Name, requestBody.Category, requestBody.Condition)

    c.JSON(http.StatusOK, gin.H{"message": "Equipo agregado exitosamente"})
}