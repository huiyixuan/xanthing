package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"xanthing/config"
	"xanthing/internal/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var httpServerCmd = &cobra.Command{
	Use:   "http",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)
		port := config.GetConfig("http_port")
		r := gin.New()
		route.SetRoute(r)
		srv := &http.Server{
			Addr:    ":" + cast.ToString(port),
			Handler: r,
		}
		fmt.Println("listening on ", port)
		//
		go func() {
			err := srv.ListenAndServe()
			if err != nil {
				log.Fatalf(err.Error())
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit

		log.Println("Shutdown Server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("server shutdown")
		}
		log.Println("server exiting")
	},
}
