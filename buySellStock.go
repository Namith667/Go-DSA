package main
import "fmt"


func maxProfit(prices []int) int{

	minPrice := prices[0]
	maxProfit:= 0

	for _,price := range prices{
		if(price<minPrice){
			minPrice = price
		}else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
 return maxProfit
}

func main(){

	prices := []int{7,1,5,3,6,4}

	fmt.Println(maxProfit(prices))
}