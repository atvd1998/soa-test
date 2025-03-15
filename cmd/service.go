package cmd

import (
	"context"
	"fmt"
	"net/http"
	"soa-product-management/internal/config"
	"soa-product-management/internal/handler"
	"soa-product-management/internal/repository/store"
	"soa-product-management/internal/usecase"

	_ "soa-product-management/docs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

var service = &cobra.Command{
	Use:   "service",
	Short: "API Command of service",
	Long:  "API Command of service",
	Run: func(_ *cobra.Command, _ []string) {
		conf := config.MustLoad()

		fx.New(
			fx.StartTimeout(conf.App.StartTimeout),
			fx.StopTimeout(conf.App.StopTimeout),
			fx.Provide(
				NewEchoServer,
				handler.NewProductHanlder,
				usecase.NewProductUsecase,
				store.InitStorage,
			),
			fx.Supply(conf),
			fx.Invoke(func(*echo.Echo) {}),
		).Run()
	},
}

func NewEchoServer(lc fx.Lifecycle, productHandler handler.ProductHandler, conf *config.Config) *echo.Echo {
	e := echo.New()
	e.GET("/api/products", productHandler.GetListProductInformationHandler)
	e.POST("/api/products:generate-pdf", productHandler.GenerateProductInfromationPdf)
	e.GET("/api/statistics/products-per-category", productHandler.GetListProductStatisticPerCategory)
	e.GET("/api/statistics/products-per-supplier", productHandler.GetListProductStatisticPerSupplier)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(conf.App.HTTPAddr); err != nil && err != http.ErrServerClosed {
					fmt.Println("Error starting server:", err)
				}
			}()
			fmt.Println("Starting Echo server...")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping Echo server...")
			return e.Shutdown(ctx)
		},
	})

	return e
}
