package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	/* SHA-1 */
	s1 := sha1.New()
	io.WriteString(s1, "ABCDE")
	fmt.Printf("%x\n", s1.Sum(nil))

	/* SHA-256 */
	s256 := sha256.New()
	io.WriteString(s256, "ABCDE")
	fmt.Printf("%x\n", s256.Sum(nil))

	/* SHA-512 */
	s512 := sha512.New()
	io.WriteString(s512, "ABCDE")
	fmt.Printf("%x\n", s512.Sum(nil))
}
