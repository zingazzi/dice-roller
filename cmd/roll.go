/*
Copyright © 2022 MARCO ZINGONI
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var dice int
var number int
var t string

type IData struct {
	number int
	amount int
}

var roll = &cobra.Command{
	Use:   "dice",
	Short: "Roll a dice",
	Long:  `Roll a dice and get a random number between 1 and dice type`,

	Run: func(cmd *cobra.Command, args []string) {
		res := 0
		for i := 0; i < number; i++ {
			val := rollDice(dice)
			fmt.Println("Rolling a dice...", val)
			if t == "sum" {
				res += val
			}
			if t == "best" {
				if val > res {
					res = val
				}
			}
			if t == "worst" {
				if res == 0 {
					res = val
				}
				if val < res {
					res = val
				}
			}
		}
		if number > 1 {
			fmt.Println("Result:", res)
		}
	},
}

var massive = &cobra.Command{
	Use:   "Massive roling",
	Short: "Roll a dice huge amount of times",
	Long:  `Roll a dice huge amount of times and get stats`,

	Run: func(cmd *cobra.Command, args []string) {
		res := []IData{}
		n := 100000
		for i := 0; i < n; i++ {
			val := rollDice(dice)
			find := false
			for j := 0; j < len(res); j++ {
				if res[j].number == val {
					res[j].amount++
					find = true
				}
			}
			if find == false {
				res = append(res, IData{val, 1})
			}
		}
		for i := 0; i < len(res)-1; i++ {
			fmt.Println(res[i])
		}
	},
}

func Execute() {
	err := roll.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	roll.Flags().IntVarP(&dice, "dice", "d", 6, "Possible dice types: 4, 6, 8, 10, 12, 20, 100")
	roll.Flags().IntVarP(&number, "number", "n", 1, "Number of dice")
	roll.Flags().StringVarP(&t, "type", "t", "sum", "Possible values: [sum, best, worst]")
}

func rollDice(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	roll := rand.Intn(max-min+1) + min
	return roll
}
