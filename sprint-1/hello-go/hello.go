/*Rich Goluszka
 *Eric Pogue
 *CPSC 360-1: Hello World in Go
 *1/26/2021
 *
 *from cl.cam.ac.uk/~mgk25/ucs/quotes.html
 *curly apostrophes are available as special characters
 *through unicode U+2018 and U+2019. They are added to
 *strings using the \u escape character
 */

package main

import "fmt"

func main() {
	greeting := "Rich Goluszka" + "\u2019" + "s Hello World"
	fmt.Println(greeting)
}
