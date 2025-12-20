package main

import "fmt"

func main() {

	// Anti pattern (architecturally wrong approach, even though it works but should not be used like this)
	// user := map[string]interface{}{
	// 	"id":   1,
	// 	"name": "Alice",
	// }

	// should be a struct
	type User struct {
		id   int
		name string
	}

	userMap := map[int]User{}

	for i := 0; i <= 10; i++ {
		userMap[i] = User{
			id:   i,
			name: fmt.Sprintf("User%d", i),
		}
	}

	fmt.Printf("userMap => %v\n", userMap)
	// map[0:{0 User0} 1:{1 User1} 2:{2 User2} 3:{3 User3} 4:{4 User4} 5:{5 User5} 6:{6 User6} 7:{7 User7} 8:{8 User8} 9:{9 User9} 10:{10 User10}]
	fmt.Println(userMap[2].name) // User2

	/*
	   STRUCTS vs MAPS IN GO — ARCHITECTURAL NOTE

	   Context:
	   --------
	   In Go, structs and maps are NOT interchangeable.
	   They solve different problems and choosing the wrong one leads to fragile code.

	   ----------------------------------------------------------------
	   ANTI-PATTERN (Common JS Habit — Avoid in Go):
	   ----------------------------------------------------------------

	   user := map[string]interface{}{
	       "id":   1,
	       "name": "Alice",
	   }

	   Why this is problematic:
	   - No compile-time guarantees (typos like "nmae" compile fine)
	   - Requires unsafe type assertions everywhere
	   - Refactoring is dangerous (no compiler help)
	   - IDE autocomplete & static analysis are lost
	   - Domain meaning is hidden behind strings

	   This creates "stringly-typed" data.
	   It works, but it scales poorly and fails silently.

	   ----------------------------------------------------------------
	   IDIOMATIC GO APPROACH (Preferred):
	   ----------------------------------------------------------------

	   type User struct {
	       ID   int
	       Name string
	   }

	   user := User{
	       ID:   1,
	       Name: "Alice",
	   }

	   Benefits:
	   - Strong compile-time type safety
	   - Clear domain intent
	   - Faster field access (no hashing)
	   - Safe refactoring
	   - Excellent tooling, JSON, DB support

	   If the data has a known shape → ALWAYS prefer structs.

	   ----------------------------------------------------------------
	   WHEN map[string]interface{} IS ACCEPTABLE:
	   ----------------------------------------------------------------
	   Use it ONLY when the schema is unknown at compile time, such as:
	   - Parsing arbitrary / external JSON
	   - Middleware or proxy layers
	   - Logging dynamic metadata
	   - Temporary ingestion of unstructured data

	   Even in these cases:
	   - Convert map → struct as early as possible
	   - Do NOT let maps leak into core business logic

	   ----------------------------------------------------------------
	   RULE OF THUMB (MEMORIZE THIS):
	   ----------------------------------------------------------------
	   - If data has MEANING → use a struct
	   - If data has DYNAMIC KEYS → use a map
	   - If unsure → default to struct

	   map[string]interface{} is raw material, not a domain model.

	   Go is opinionated by design.
	   It rewards clarity and punishes ambiguity.
	*/

}
