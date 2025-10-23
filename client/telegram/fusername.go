package telegram

import "fmt"

func FormatUsername(us string) string {
	return fmt.Sprintf("@%v", us)
}
