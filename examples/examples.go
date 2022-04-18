package examples

import (
	"fmt"
	"log"

	"github.com/open-source/hasher"
)

func main() {
	argon2ID := hasher.NewArgon2ID()

	plain := "inanzzz"

	hash1, err := argon2ID.Hash(plain)
	if err != nil {
		log.Fatal(err)
	}
	ok1, err := argon2ID.Verify(plain, hash1)
	if err != nil {
		log.Fatal(err)
	}

	hash2, err := argon2ID.Hash(plain)
	if err != nil {
		log.Fatal(err)
	}
	ok2, err := argon2ID.Verify(plain, hash2)
	if err != nil {
		log.Fatal(err)
	}

	hash3, err := argon2ID.Hash(plain)
	if err != nil {
		log.Fatal(err)
	}
	ok3, err := argon2ID.Verify(plain, hash3+"uuu")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PLAIN:", plain)
	fmt.Println("HASH 1:", hash1)
	fmt.Println("VALID 1:", ok1)
	fmt.Println("HASH 2:", hash2)
	fmt.Println("VALID 2:", ok2)
	fmt.Println("HASH 3:", hash3)
	fmt.Println("VALID 3:", ok3)
}
