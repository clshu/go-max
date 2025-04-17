package maps

import "fmt"

func maps() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Reddit": "https://www.reddit.com",
		"GitHub": "https://www.github.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Facebook"] = "https://www.facebook.com"
	fmt.Println(websites)
	delete(websites, "Reddit")
	fmt.Println(websites)
}
