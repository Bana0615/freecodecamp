package main

import "fmt"

/* Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping. */

func main() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21 // overwrites 24

	ages = map[string]int{
		"John": 37,
		"Mary": 21,
	}
	fmt.Println(len(ages))

	/* Map Mutations
	   m[key] = elem
	   -- Get an element
	   elem = m[key]
	   -- Delete an element
	   delete(m, key)
	   -- Check if a key exists
	   elem, ok := m[key]
	*/
}
