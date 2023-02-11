package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tangx/srv-sensetive-text/pkg/version"
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
	Use:   "validator",
	Short: fmt.Sprintf("version: %s", version.Version),
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
