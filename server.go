package main

import (
	"RestApi-Echo/config"
	"RestApi-Echo/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

func main() {
	config.ConnectDB()
	route := echo.New()
	route.POST("user/create_user", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		contentType := c.Request().Header.Get("Content-type")
		if contentType == "application/json" {
			fmt.Println("Request from json")
		} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
			file, err := c.FormFile("name")
			if err != nil {
				fmt.Println("Name null")
			} else {
				src, err := file.Open()
				if err != nil {
					return err
				}
				defer src.Close()
				dst, err := os.Create(file.Filename)
				if err != nil {
					return err
				}
				defer dst.Close()
				if _, err = io.Copy(dst, src); err != nil {
					return err
				}

				user.Name = file.Filename
				fmt.Println("File found, will save")
			}
		}
		response := new(Response)
		if user.CreateUser() != nil { // method create user
			response.ErrorCode = 10
			response.Message = "Failed create user data"
		} else {
			response.ErrorCode = 0
			response.Message = "Success create user data"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:name", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		response := new(Response)
		if user.UpdateUser(c.Param("name")) != nil { //method update data
			response.ErrorCode = 10
			response.Message = "Failed update user data"
		} else {
			response.ErrorCode = 0
			response.Message = "Success update user data"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_user/:name", func(c echo.Context) error {
		user, _ := model.GetOneByName(c.Param("name")) // method get by emaik
		response := new(Response)

		if user.DeleteUser() != nil { // method update user
			response.ErrorCode = 10
			response.Message = "failed remove data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Success remove data user"
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/search_user", func(c echo.Context) error {
		response := new(Response)
		users, err := model.GetAll(c.QueryParam("keywords")) //method get all
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Failed get data user"
		} else {
			response.ErrorCode = 0
			response.Message = "Success get data user"
			response.Data = users
		}
		return c.JSON(http.StatusOK, response)
	})

	route.Start(":9000")
}
