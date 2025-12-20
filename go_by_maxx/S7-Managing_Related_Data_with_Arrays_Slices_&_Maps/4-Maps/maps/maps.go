package maps

import "fmt"

func Maps() {
	// map[KeyType]ValueType
	links := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Youtube":  "https://youtube.com",
	}

	fmt.Printf("links => %v\n", links)
	fmt.Printf("links[Google] => %v\n", links["Google"])
	fmt.Printf("links[Facebook] => %v\n", links["Facebook"])
	fmt.Printf("links[Youtube] => %v\n", links["Youtube"])

	fmt.Printf("Length of links map is %v", len(links)) // 3

	links["LinkedIn"] = "https://linkedin.com"
	fmt.Printf("links after adding LinkedIn => %v\n", links)
	fmt.Printf("Length of links map is %v\n", len(links)) // 4
	// append in not supported in maps like arrays or slices

	// to delete map item: the map from which we want to delete & the key of the item we want to delete
	delete(links, "Youtube")
	fmt.Printf("links after deleting Youtube => %v\n", links)
	fmt.Printf("Length of links map is %v\n", len(links)) // 3
}
