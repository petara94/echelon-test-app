package main

import (
	"encoding/json"
	"executer_server/executer"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func main() {
	machine, err := executer.AutoStartMachine()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/os", func(c *gin.Context) {
		c.JSON(200, machine)
	})

	router.GET("/exec", func(c *gin.Context) {
		bodyRaw, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, executer.NewBadExecResult(nil, err.Error()))
			return
		}

		var body executer.RequestBody
		err = json.Unmarshal(bodyRaw, &body)
		if err != nil {
			c.JSON(400, executer.NewBadExecResult(nil, err.Error()))
			return
		}

		if body.OS != machine.OS {
			c.JSON(400, executer.NewBadExecResult(&body, executer.ERROR_EXEC_OS))
			return
		}

		result, err := machine.Exec(body.CMD, body.Stdin)
		if err != nil {
			c.JSON(400, executer.NewBadExecResult(&body, err.Error()))
			return
		}
		c.JSON(200, result)
	})

	err = router.Run(":" + executer.PORT)
	if err != nil {
		log.Fatal(err)
	}

}
