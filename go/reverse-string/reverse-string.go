package reverse

func String(str string) string {
    reversed := ""

    for _, char := range str {
        reversed = string(char) + reversed
    }

    return reversed
}
