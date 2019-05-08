// Package thoughts provides happy or sad wishes
package thoughts

import "fmt"

// Happy Thoughts
func Happy(s string) string {
	return fmt.Sprintf("Happy day to you %v", s)
}

// Sad Thoughts
func Sad(s string) string {
        return fmt.Sprintf("Sad day to you %v", s)
}
