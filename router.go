package main

import (
	"context"
	"log"

	account "github.com/fox-one/f1db/account"

	config "github.com/fox-one/f1db/config"
	ctrl "github.com/fox-one/f1db/controller"
	"github.com/gin-gonic/gin"
)

func serve(ctx context.Context, pk string) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello.",
		})
	})
	router.POST("/register", ctrl.RegisterHandler(ctx, pk))
	router.POST("/login", ctrl.LoginHandler(ctx))
	router.GET("/quota", account.AuthRequired(), ctrl.QuotaHandler(ctx))
	router.POST("/records", account.AuthRequired(), ctrl.NewRecordHandler)
	router.GET("/records/:hash", ctrl.GetRecordHandler)
	router.POST("/records/:hash/keep", account.AuthRequired(), ctrl.KeepRecordHandler(pk))
	router.GET("/snapshots/:snapshot_id", ctrl.GetSnapshotHandler(ctx))
	log.Printf("Server starts on %s\n", config.GetConfig().Server.Host)
	router.Run(config.GetConfig().Server.Host)
}
