package main

func elemDistinto(arrA, arrB []int, ini, fin int) int {
	if ini == fin {
		return arrA[ini]
	}

	medio := (ini + fin) / 2

	if arrA[medio] == arrB[medio] {
		return elemDistinto(arrA, arrB, medio+1, fin)
	}

	return elemDistinto(arrA, arrB, ini, medio)
}
