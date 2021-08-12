package main


import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main() {
	// router := http router.New()
	// router.GET("/", Index)
	// log.Fatal(http.ListenAndServe(":8080", router))
	//fmt.Println(sum(5))
	//fmt.Println(sum(5))
	//fmt.Println(multiplier(5))
	//fmt.Println(addition(2, 3, 4))
	//fmt.Println(multiply(2,3,4,5))
	//fmt.Println(subtract(3,4))
	//fmt.Println(divide(10, 2.5))
	//fmt.Println(multiplyValues(2,3,4))
	//fmt.Println(factorial(5))
	//fmt.Println("........________________________ _____..........")
	//fmt.Println("........recursion test for subtraction..........")
	//fmt.Println(subtractAll(10-5-2))//3 correct
	//fmt.Println(subtractAll(10-50-20-10))//-70 correct
	//fmt.Println(subtractAll(100-50-30-20))// 0 correct
	//fmt.Println(subtractAll(-1-1-1-1))// -4 correct
	//fmt.Println(subtractAll(100-50-20-10))// 20
	//
	//fmt.Println("........________________________ _____..........")
	//fmt.Println("........recursion test for division..........")

//	fmt.Println(divideAll(4/4/2))


	//fmt.Println(multiplyValues(2.00, 8.10))
	m := orderedmap.NewOrderedMap()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set(123, true)

	m.Delete("qux")
	fmt.Println(m)

	for _, key := range m.Keys() {
		value, _ := m.Get(key)
		fmt.Println(key, value)
	}

	// Iterate through all elements from oldest to newest:
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}

	// You can also use Back and Prev to iterate in reverse:
	for el := m.Back(); el != nil; el = el.Prev() {
		fmt.Println(el.Key, el.Value)
	}

}

// func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	logger, _ := zap.NewProduction()
// 	logger.Info("Successfully performed http request")
// }
