package main

import (
	"fmt"
	"sort"
)

type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge2 []user3

func (a ByAge2) Len() int           { return len(a) }
func (a ByAge2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge2) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast2 []user3

func (l ByLast2) Len() int           { return len(l) }
func (l ByLast2) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast2) Less(i, j int) bool { return l[i].Last < l[j].Last }

func main() {
	u1 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user3{u1, u2, u3}

	fmt.Println(users)

	// Hands-on exercise #5 code:
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("----------------")

	sort.Sort(ByAge2(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("----------------")

	sort.Sort(ByLast2(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}
