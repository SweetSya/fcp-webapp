package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	target := c.Param("id")
	id, err := strconv.Atoi(target)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid task ID"))
		return
	}
	var reqBody = model.Task{}

	if err = c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("Something went wrong while updating task!"))
		return
	}
	reqBody.ID = id
	if err = t.taskService.Update(id, &reqBody); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("Something went wrong while updating task!"))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("update task success"))
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	target := c.Param("id")
	id, err := strconv.Atoi(target)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("Invalid task ID"))
		return
	}
	if err = t.taskService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("Something went wrong while updating task!"))
		return
	}
	c.JSON(http.StatusOK, model.NewSuccessResponse("delete task success"))
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	lists, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	target := c.Param("id")
	id, err := strconv.Atoi(target)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("Invalid task ID"))
		return
	}
	lists, err := t.taskService.GetTaskCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}
