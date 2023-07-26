package utils

func ToClassName(s string) string{
    bs := []byte(s)
    if len(bs) == 0 {
        return ""
    }
    bs[0] = byte(bs[0] - 32)
    return string(bs)
}
