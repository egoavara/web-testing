package main

import (
	"fmt"
)

type Iter[T any] func() (T, bool)

func FromArray[T any](array []T) Iter[T] {
	var index = -1
	var end T
	return func() (T, bool) {
		if index+1 < len(array) {
			index += 1
			return array[index], true
		}
		return end, false
	}
}

func Filter[T any](origin Iter[T], fn func(T) bool) Iter[T] {
	var end T
	return func() (T, bool) {
		for data, ok := origin(); ok; data, ok = origin() {
			if fn(data) {
				return data, true
			}
		}
		return end, false
	}
}

func Map[T any, E any](origin Iter[T], fn func(T) E) Iter[E] {
	var end E
	return func() (E, bool) {
		for data, ok := origin(); ok; data, ok = origin() {
			return fn(data), true
		}
		return end, false
	}
}
func Sorted[T any](origin Iter[T], fn func(a, b T) int) Iter[T] {
	var collect = Collect(origin)
	var n = len(collect)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if fn(collect[j-1], collect[j]) > 0 {
				collect[j-1], collect[j] = collect[j], collect[j-1]
			}
			j = j - 1
		}
	}
	return FromArray(collect)
}
func Collect[T any](origin Iter[T]) []T {
	var collect = make([]T, 0, 10)
	for data, ok := origin(); ok; data, ok = origin() {
		collect = append(collect, data)
	}
	return collect
}
func ForEach[T any](origin Iter[T], fn func(T)) {
	for data, ok := origin(); ok; data, ok = origin() {
		fn(data)
	}
}

type Student struct {
	name         string
	cLanguage    float32
	javaLanguage float32
}
type StudentMean struct {
	name string
	mean float32
}

func main() {
	var students []Student = []Student{
		{name: "나", cLanguage: 100, javaLanguage: 100},
		{name: "철수", cLanguage: 90, javaLanguage: 80},
		{name: "영희", cLanguage: 90, javaLanguage: 95},
	}
	ForEach(
		Sorted(
			Filter(
				Map(
					FromArray(students),
					func(t Student) StudentMean {
						return StudentMean{
							name: t.name,
							mean: (t.cLanguage + t.javaLanguage) / 2,
						}
					},
				),
				func(t StudentMean) bool {
					return t.mean >= 90
				},
			),
			func(a, b StudentMean) int {
				if a.mean == b.mean {
					return 0
				} else if a.mean < b.mean {
					return -1
				} else {
					return 1
				}
			},
		),
		func(t StudentMean) {
			fmt.Printf("%s : %f\n", t.name, t.mean)
		},
	)
}
