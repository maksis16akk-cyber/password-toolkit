package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func checkStrength(pwd string) string {
	score := 0
	if len(pwd) >= 8 { score++ }
	hasDigit, hasUpper, hasLower, hasSpecial := false, false, false, false
	for _, r := range pwd {
		switch {
		case unicode.IsDigit(r): hasDigit = true
		case unicode.IsUpper(r): hasUpper = true
		case unicode.IsLower(r): hasLower = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r): hasSpecial = true
		}
	}
	if hasDigit { score++ }
	if hasUpper && hasLower { score++ }
	if hasSpecial { score++ }
	strength := []string{"Weak", "Medium", "Strong", "Very Strong"}
	if score == 0 { return "Weak" }
	return strength[score-1]
}

func generate(length int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+"
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, length)
	for i := range res {
		res[i] = chars[rand.Intn(len(chars))]
	}
	return string(res)
}

func main() {
	gen := flag.Int("g", 0, "generate password of given length")
	check := flag.String("c", "", "check password strength")
	flag.Parse()
	if *gen > 0 {
		pwd := generate(*gen)
		fmt.Printf("Generated: %s\nStrength: %s\n", pwd, checkStrength(pwd))
	} else if *check != "" {
		fmt.Printf("Strength: %s\n", checkStrength(*check))
	} else {
		fmt.Println("Use -g <length> or -c <password>")
	}
}
