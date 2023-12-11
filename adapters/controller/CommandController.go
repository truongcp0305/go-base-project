package controller

import (
	"go-project/adapters/incoming"
	"go-project/usecase/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommandController struct {
	cmdService service.CommandService
}

func NewCommandController(cmd service.CommandService) CommandController {
	return CommandController{
		cmdService: cmd,
	}
}

// @Summary List Process
// @Description Get all process runing in computer
// @Tags Command
// @Success 200 {object} outgoing.ListProcessOutgoing
// @Failure 500 {object} outgoing.ModelInternalErr
// @Router /cmd/list-process [get]
func (cC *CommandController) HandleListProcess(c *gin.Context) {
	listPr, err := cC.cmdService.ListProcess()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": listPr})
}

// @Summary Kill Process
// @Description Kill a process is runing
// @Tags Command
// @Param Body body incoming.KillProcessParams true "query params"
// @Success 200 {object} outgoing.KillProcessOutgoing
// @Failure 400 {object} outgoing.ModelBadRequestErr
// @Failure 500 {object} outgoing.ModelInternalErr
// @Router /cmd/kill-process [post]
func (cC *CommandController) HandleKillProcess(c *gin.Context) {
	var params incoming.KillProcessParams
	var err error
	if err = c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if params.Pid != "" {
		err = cC.cmdService.KillByPid(params.Pid)
	} else if params.Name != "" {
		err = cC.cmdService.KillByName(params.Name)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "process killed!"})
}

// @Summary Execute Script
// @Description Execute command line in a text script file
// @Tags Command
// @Param Body body incoming.ExcecuteScriptParams true "query params"
// @Success 200 {object} outgoing.ExecuteScriptOutgoing
// @Failure 400 {object} outgoing.ModelBadRequestErr
// @Failure 500 {object} outgoing.ModelInternalErr
// @Router /cmd/execute-script [post]
func (cC *CommandController) HandleExecuteScript(c *gin.Context) {
	var params incoming.ExcecuteScriptParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if params.Path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	_, err := cC.cmdService.ExcuteScript(params.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "excute done!"})
}

// @Summary Find File
// @Description Find and open file by path
// @Tags Command
// @Param Body body incoming.GetFileParams true "query params"
// @Success 200 {object} outgoing.OpenFileSucessOutgoing
// @Failure 400 {object} outgoing.ModelBadRequestErr
// @Failure 500 {object} outgoing.OpenFileFailedOutgoing
// @Router /cmd/open-file [post]
func (cC *CommandController) HandleOpenFile(c *gin.Context) {
	var params incoming.GetFileParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if params.Path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	_, err := cC.cmdService.OpenFilePath(params.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find file with path"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "opening file"})
}
