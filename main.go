package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
)

var filepath string


func main() {

	var exists bool
	filepath, exists = os.LookupEnv("FILE_PATH_TXN")
	// filepath = "/Users/abc/Desktop/scripts/main.bash" // for local testing
	// exists = true
	if !exists {
		log.Println("File Path not found.")
	}

	e := echo.New()

	// define your routes
	e.GET("/", scriptHandler)
	e.POST("/src", txnScriptPostHandler)
	e.GET("/test", testHandler)
	e.GET("/healthz", pinger)
	e.Logger.Fatal(e.Start(":8080"))

}

func pinger(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, "OK", " ")
}

func testHandler(c echo.Context) error {
	cmd := exec.CommandContext(context.Background(), "sh", "/app/script.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err.Error(), "  ")
	}
	return c.JSONPretty(http.StatusOK, string(out), " ")
}

func txnScriptPostHandler(c echo.Context) error {

	bodyMap := make(map[string]string)

	body := c.Request().Body
	if body != nil {
		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Println("Bad request", err.Error())
			return c.JSONPretty(http.StatusBadRequest, err.Error(), " ")
		}
		err = json.Unmarshal(reqBody, &bodyMap)
		if err != nil {
			log.Println("unmarshalling error", err.Error())
			return c.JSONPretty(http.StatusInternalServerError, err.Error(), " ")
		}
		log.Println(string(reqBody))
	} else {
		log.Println("No body")
	}
	go ExecuteThis(filepath, bodyMap)
	return c.JSONPretty(http.StatusOK, "OK", " ")
}

func scriptHandler(c echo.Context) error {
	go ExecuteThis(filepath, make(map[string]string))
	return c.JSONPretty(http.StatusOK, "OK", " ")
}
