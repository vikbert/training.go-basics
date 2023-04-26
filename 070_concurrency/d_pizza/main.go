package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Ingredient struct {
	name       string
	prepTime   int
	isPrepared bool
}

func (ing *Ingredient) prepare() {
	log.Printf("%s is being prepared\n", ing.name)
	time.Sleep(time.Duration(ing.prepTime) * time.Second)
	ing.isPrepared = true
}

type Pizza struct {
	name        string
	ingredients []*Ingredient
	bakingTime  int
}

type Oven struct {
	prepTime   int
	isPrepared bool
}

func (o *Oven) prepare() {
	log.Println("Oven is heating up...")
	time.Sleep(time.Duration(o.prepTime) * time.Second)
	o.isPrepared = true
}

func (o *Oven) bake(pizza *Pizza) error {
	log.Printf("Trying to bake Pizza %s\n", pizza.name)
	if !o.isPrepared {
		return errors.New("oven is not prepared")
	}
	for _, ing := range pizza.ingredients {
		if !ing.isPrepared {
			return fmt.Errorf("ingredient %q is not prepared", ing.name)
		}
	}
	time.Sleep(time.Duration(pizza.bakingTime) * time.Second)
	log.Printf("Pizza %s is done!\n", pizza.name)
	return nil
}

var (
	dough1        = &Ingredient{"Dough", 4, false}
	tomatoSauce1  = &Ingredient{"Tomato Sauce", 1, false}
	gratedCheese1 = &Ingredient{"Grated Cheese", 2, false}
	dough2        = &Ingredient{"Dough", 4, false}
	tomatoSauce2  = &Ingredient{"Tomato Sauce", 1, false}
	gratedCheese2 = &Ingredient{"Grated Cheese", 2, false}
	salami2       = &Ingredient{"Salami", 1, false}
	onions2       = &Ingredient{"Onions", 3, false}
)

var (
	pizzaMargarita = &Pizza{
		"Margarita",
		[]*Ingredient{dough1, tomatoSauce1, gratedCheese1},
		5,
	}
	pizzaSalamiSpecial = &Pizza{
		"Salami Speciale",
		[]*Ingredient{dough2, tomatoSauce2, gratedCheese2, salami2, onions2},
		6,
	}
)

var oven = Oven{8, false}

func main() {
	// we want to measure time until both pizza are done
	startTime := time.Now()

	// prep oven
	oven.prepare()

	// prep ingredients
	for _, p := range []*Pizza{pizzaMargarita, pizzaSalamiSpecial} {
		for _, ing := range p.ingredients {
			ing.prepare()
		}
	}

	// lets bake some pizza
	err := oven.bake(pizzaMargarita)
	handleError("baking pizza", err)
	err = oven.bake(pizzaSalamiSpecial)
	handleError("baking pizza", err)

	// measure
	log.Printf("All pizza done after %d seconds\n", int32(time.Since(startTime).Seconds()))
}

func handleError(context string, err error) {
	if err != nil {
		log.Printf("Got error %s: %s", context, err.Error())
		os.Exit(2)
	}
}
