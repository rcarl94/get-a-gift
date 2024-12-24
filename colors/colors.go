package colors

import "fmt"

const red = "\033[31m"
const green = "\033[32m"
const reset = "\033[0m"

func Red(stringToColor string) string {
  return fmt.Sprintf("%s%s%s", red, stringToColor, reset)
}

func Green(stringToColor string) string {
  return fmt.Sprintf("%s%s%s", green, stringToColor, reset)
}
