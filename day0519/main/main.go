package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	lenS := len(s)
	max := s[lenS-1]
	fast := 0
	slow := 0

	for ; fast < len(g); fast++ {
		if g[fast] > max || slow >= lenS {
			break
		}
		for ; slow < len(s); slow++ {
			if s[slow] >= g[fast] {
				slow++
				break
			}
		}
	}
	return fast
}

func findContentChildrenII(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	cookieSum := 0
	res := 0
	var pickCookie func(count int)
	pickCookie = func(count int) {
		fmt.Println("count:", count, " s:", s)
		for i := 0; i < len(s); i++ {
			tmp := s[i]
			if tmp == count {
				cookieSum -= s[i]
				s = append(s[:i], s[i+1:]...)
				return
			} else if i == len(s)-1 || s[i+1] > count {
				cookieSum -= tmp
				s = append(s[:i], s[i+1:]...)
				if count-tmp < 0 {
					return
				}
				pickCookie(count - tmp)
				return
			}
		}
		return
	}
	for i := 0; i < len(s); i++ {
		cookieSum += s[i]
	}
	for ; res < len(g); res++ {
		if g[res] > cookieSum {
			break
		}
		pickCookie(g[res])
	}
	return res
}

func wiggleMaxLength(nums []int) int {
	numsL := len(nums)
	if numsL <= 1 {
		return numsL
	}
	slow := 1
	res := 1
	flagN := true

	for ; slow < numsL; slow++ {
		if nums[slow] != nums[slow-1] {
			if nums[slow] > nums[slow-1] {
				flagN = true
			} else {
				flagN = false
			}
			res++
			break
		}
	}
	if slow == numsL-1 {
		return res
	}
	for i := slow + 1; i < numsL; i++ {
		if flagN {
			if nums[i] < nums[i-1] {
				flagN = false
				res++
			}
		} else {
			if nums[i] > nums[i-1] {
				flagN = true
				res++
			}
		}
	}
	return res
}

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func generatePassword(seed, appName, username string, includeLowercase, includeUppercase, includeSpecialChars, includeDigits bool) string {
	// Concatenate seed, appName, and username
	inputStr := seed + appName + username

	// Perform hash function
	hash := sha256.New()
	hash.Write([]byte(inputStr))
	hashValue := hash.Sum(nil)

	// Set the random seed based on the hash value
	rand.Seed(int64(toInt(hashValue)))

	// Define character pools for each character type
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specialChars := "!@#$%^&*()_"

	// Generate the password with required character types
	var password strings.Builder

	// Check if lowercase letters are required
	if includeLowercase {
		password.WriteByte(lowercaseLetters[rand.Intn(len(lowercaseLetters))])
	}

	// Check if uppercase letters are required
	if includeUppercase {
		password.WriteByte(uppercaseLetters[rand.Intn(len(uppercaseLetters))])
	}

	// Check if digits are required
	if includeDigits {
		password.WriteByte(digits[rand.Intn(len(digits))])
	}

	// Check if special characters are required
	if includeSpecialChars {
		password.WriteByte(specialChars[rand.Intn(len(specialChars))])
	}

	// Generate the remaining characters
	characterPool := ""
	if includeLowercase {
		characterPool += lowercaseLetters
	}
	if includeUppercase {
		characterPool += uppercaseLetters
	}
	if includeDigits {
		characterPool += digits
	}
	if includeSpecialChars {
		characterPool += specialChars
	}

	for i := password.Len(); i < 16; i++ {
		password.WriteByte(characterPool[rand.Intn(len(characterPool))])
	}

	return password.String()
}

func toInt(bytes []byte) int {
	result := 0
	for _, b := range bytes {
		result = result*256 + int(b)
	}
	return result
}

func main() {
	// Define command line flags
	seed := flag.String("seed", "", "Seed for password generation")
	appName := flag.String("appName", "", "Application name")
	username := flag.String("username", "", "Username")
	includeLowercase := flag.Bool("includeLowercase", true, "Include lowercase letters")
	includeUppercase := flag.Bool("includeUppercase", true, "Include uppercase letters")
	includeSpecialChars := flag.Bool("includeSpecialChars", true, "Include special characters")
	includeDigits := flag.Bool("includeDigits", true, "Include digits")

	// Parse command line flags
	flag.Parse()

	// Generate password
	password := generatePassword(*seed, *appName, *username, *includeLowercase, *includeUppercase, *includeSpecialChars, *includeDigits)
	fmt.Println(password)
}
