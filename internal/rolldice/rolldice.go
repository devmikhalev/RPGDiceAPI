package rolldice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type Params struct {
	Faces    int  `uri:"faces" binding:"required"`
	Amount   int  `uri:"amount" binding:"required"`
	Modifier *int `uri:"mod" binding:"required"` // pointer because a zero value must be processed
}

type Roll struct {
	Result   int
	Notation string
	Dices    []int
}

func newRoll(params Params) *Roll {
	var roll Roll
	sum := 0
	roll.Dices = make([]int, 0, params.Amount)

	for i := 0; i < params.Amount; i++ {
		face := rand.Intn(params.Faces-1) + 1
		roll.Dices = append(roll.Dices, face)
		sum += face
	}

	if *params.Modifier == 0 {
		roll.Notation = fmt.Sprintf("%vd%v", params.Amount, params.Faces)
	} else if *params.Modifier < 0 {
		roll.Notation = fmt.Sprintf("%vd%v%v", params.Amount, params.Faces, *params.Modifier)
	} else {
		roll.Notation = fmt.Sprintf("%vd%v+%v", params.Amount, params.Faces, *params.Modifier)
	}

	roll.Result = sum + *params.Modifier

	return &roll
}

func MakeRoll(context *gin.Context) {
	var params Params
	if err := context.ShouldBindUri(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	roll := newRoll(params)

	context.JSON(http.StatusOK, gin.H{"Result": roll.Result, "Notation": roll.Notation, "Dices": roll.Dices})
}
