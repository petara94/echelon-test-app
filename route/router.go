package route

import (
	"echelon-test-app/executer"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"sync"
)

// InitRouter Иницализация серрвера, роутинг
func InitRouter() *gin.Engine {
	router := gin.Default()

	rg := router.Group("api/v1")
	rg.GET("/os", func(c *gin.Context) {
		c.JSON(200, executer.MainMachine)
	})

	rg.GET("/exec", RouteExec)
	rg.GET("/remote-execution", RouteExecAll)
	rg.GET("/remote-execution-async", RouteAsyncExecAll)

	return router
}

// RouteExec Выполнение команды
func RouteExec(c *gin.Context) {
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

	if body.OS != executer.MainMachine.OS {
		c.JSON(400, executer.NewBadExecResult(&body, executer.ERROR_EXEC_OS))
		return
	}

	result, err := executer.MainMachine.Exec(body.CMD, body.Stdin)

	if err != nil {
		c.JSON(400, executer.NewBadExecResult(&body, err.Error()))
		return
	}
	c.JSON(200, result)
}

// RouteExecAll Выполнение массива команд
func RouteExecAll(c *gin.Context) {
	bodyRaw, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(400, executer.NewBadExecResult(nil, err.Error()))
		return
	}

	var body []executer.RequestBody
	err = json.Unmarshal(bodyRaw, &body)
	if err != nil {
		c.JSON(400, executer.NewBadExecResult(nil, err.Error()))
		return
	}

	Results := make([]interface{}, 0)

	for _, task := range body {
		if task.OS != executer.MainMachine.OS {
			Results = append(Results, executer.NewBadExecResult(&task, executer.ERROR_EXEC_OS))
			continue
		}

		ans, err := executer.MainMachine.Exec(task.CMD, task.Stdin)

		if err != nil {
			Results = append(Results, executer.NewBadExecResult(&task, err.Error()))
			continue
		}

		Results = append(Results, ans)
	}

	c.JSON(200, Results)
}

// RouteAsyncExecAll Асинхронное выполнение массива команд
func RouteAsyncExecAll(c *gin.Context) {
	bodyRaw, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(400, executer.NewBadExecResult(nil, err.Error()))
		return
	}

	var body []executer.RequestBody
	err = json.Unmarshal(bodyRaw, &body)
	if err != nil {
		c.JSON(400, executer.NewBadExecResult(nil, err.Error()))
		return
	}

	Results := make([]interface{}, 0)
	wg := sync.WaitGroup{}
	wg.Add(len(body))
	for _, task := range body {
		go func(task executer.RequestBody, wg *sync.WaitGroup) {
			defer wg.Done()
			if task.OS != executer.MainMachine.OS {
				Results = append(Results, executer.NewBadExecResult(&task, executer.ERROR_EXEC_OS))
				return
			}

			ans, err := executer.MainMachine.Exec(task.CMD, task.Stdin)

			if err != nil {
				Results = append(Results, executer.NewBadExecResult(&task, err.Error()))
				return
			}
			Results = append(Results, ans)
		}(task, &wg)
	}

	wg.Wait()

	c.JSON(200, Results)
}
