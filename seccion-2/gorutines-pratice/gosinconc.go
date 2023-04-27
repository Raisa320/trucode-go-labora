package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	userName := fetchUser()
	likes := fetchUserLikes(userName)
	match := fetchUserMatch(userName)

	fmt.Println("Likes:", likes)
	fmt.Println("Match:", match)
	fmt.Println("took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func fetchUserLikes(username string) int {
	time.Sleep(time.Millisecond * 150)
	return 11
}

func fetchUserMatch(username string) string {
	time.Sleep(time.Microsecond * 100)
	return "ANA"
}
