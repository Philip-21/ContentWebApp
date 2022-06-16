package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

//this response file contains things that will be used in various parts of the application interms of error handling

func ClientError(c *gin.Context, status int) {
	//write to the info log
	app.InfoLog.Println("Client error with status of", status)
	c.JSON(http.StatusInternalServerError, status)
}

func ServerError(c *gin.Context, err error) {
	//getting the trace of the error
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack()) //err.Error() is the error message debug.Stack()this gives the detailed information about the error that took place
	//the error log writes it to the terminal in development
	//in production the error log writes it in a log file ,put a directive for the user to check email ,see the error message in the log file and fix the error

	//writing error log to the terminal window
	app.ErrorLog.Println(trace)
	c.JSON(http.StatusInternalServerError, trace)

}
