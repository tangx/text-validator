package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func main() {
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

type Response struct {
	Code    int
	Message string
}

func httpserve(port int) {
	r := gin.Default()

	r.GET("/validate/text", ValidateText)

	listen := fmt.Sprintf(":%d", port)
	r.Run(listen)
}

var root = &cobra.Command{
	Use: "validator",
	Run: func(cmd *cobra.Command, args []string) {
		InitValidator(dict)
		httpserve(port)
	},
}

var (
	port = 8081
	dict = "./dict.txt"
)

func init() {
	root.Flags().IntVarP(&port, "port", "", 8081, "服务器端口")
	root.Flags().StringVarP(&dict, "dict", "", "./dict.txt", "词典")
}
