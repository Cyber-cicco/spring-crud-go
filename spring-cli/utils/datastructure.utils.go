package utils

func Map[T, U any](ts []T, f func(T) U) []U {
    us := make([]U, len(ts))
    for i := range ts {
        us[i] = f(ts[i])
    }
    return us
}

func CopyMap[K , V comparable](m map[K]V) map[K]V{
    result := make(map[K]V)
    for k, v := range m {
        result[k] = v
    }
    return result
}
