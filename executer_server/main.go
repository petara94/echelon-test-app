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
			c.JSON(400, executer.ExecResult{})
			return
		}

		var body executer.RequestBody
		err = json.Unmarshal(bodyRaw, &body)
		if err != nil {
			c.JSON(400, executer.ExecResult{})
			return
		}

		result, err := machine.Exec(body.CMD, body.Stdin)
		if err != nil {
			c.JSON(400, executer.ExecResult{})
			return
		}
		c.JSON(200, result)
	})

	port := ""

	if machine.OS == executer.LINUX_OS {
		port = executer.DEFAULT_LINUX_PORT
	}
	if machine.OS == executer.WINDOWS_OS {
		port = executer.DEFAULT_WIN_PORT
	}

	if port == "" {
		log.Fatal("unsupported system")
	}

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
