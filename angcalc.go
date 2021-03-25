package main

import (
	"fmt"
	"os"

	calc "github.com/Klevino1/angcalc/calculations"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "angcalc",
		Usage:                "Консольный калькулятор угловых размеров.",
		EnableBashCompletion: true,
		Copyright:            "Klevino1, 2021",
		Flags: []cli.Flag{
			&cli.Float64Flag{ // Дистанция до объекта
				Name:    "distance",
				Usage:   "Дистанция до объекта. Базовое значение - 1.",
				Aliases: []string{"l"},
				Value:   1.0,
			},
			&cli.Float64Flag{
				Name:    "angle", // Угловой размер объекта
				Usage:   "Угловой размер объекта в градусах. Базовое значение - 1°.",
				Aliases: []string{"a", "alpha"},
				Value:   1.0,
			},
			&cli.Float64Flag{
				Name:    "size", // Линейный размер
				Usage:   "Линейный размер объекта. Базовое значение - 1.",
				Aliases: []string{"d"},
				Value:   1.0,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "dis",
				Usage: "Вычисление дистанции до объекта, исходя из его линейного размера и углового размера.",
				Action: func(c *cli.Context) error {
					fmt.Println("Расстояние до объекта равно %.3f\n", calc.Lcalc(c.Float64("distance"), c.Float64("angle")))
					return nil
				},
			},
		},
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Klevino1",
				Email: "animeissin1337@tutanota.com",
			},
		},
	}
	app.Run(os.Args)
}
