package utils

import (
	"fmt"
	"math/rand"
)

// All functions must start with capital letters to be exported
func GenerateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomContent() string {
	contents := []string{
		"This is really interesting...",
		"I've been thinking about this for a while...",
		"What are your thoughts on this?",
		"Has anyone else experienced this?",
		"Looking for advice on this matter.",
	}
	return contents[rand.Intn(len(contents))]
}

func GenerateRandomTitle() string {
	titles := []string{
		"Just found this interesting thing",
		"What do you think about this?",
		"Amazing discovery!",
		"Need help with this",
		"First time posting here",
	}
	return titles[rand.Intn(len(titles))]
}
func GenerateUsername(index int) string {
	return fmt.Sprintf("user_%d", index)
}
