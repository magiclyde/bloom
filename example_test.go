/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2022/3/28 下午4:14
 * @note:
 */

package bloom

import "fmt"

// Build a blacklist of shady websites.
func Example_basics() {
	blacklist := New(10000)

	url := []byte("https://rascal.com")
	blacklist.Add(url)

	// Test for membership.
	if blacklist.Test(url) {
		fmt.Println(string(url), "seems to be shady.")
	} else {
		fmt.Println(string(url), "has not yet been added to our blacklist.")
	}
	// Output: https://rascal.com seems to be shady.
}
