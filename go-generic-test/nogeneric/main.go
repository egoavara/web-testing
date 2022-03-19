package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name         string
	cLanguage    float32
	javaLanguage float32
}
type StudentMean struct {
	name string
	mean float32
}
type ArrayStudentMean []StudentMean

func (a ArrayStudentMean) Len() int           { return len(a) }
func (a ArrayStudentMean) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArrayStudentMean) Less(i, j int) bool { return a[i].mean < a[j].mean }

func main() {
	var students []Student = []Student{
		{name: "나", cLanguage: 100, javaLanguage: 100},
		{name: "철수", cLanguage: 90, javaLanguage: 80},
		{name: "영희", cLanguage: 90, javaLanguage: 95},
	}
	//
	var above90 = ArrayStudentMean(make([]StudentMean, 0, len(students)))
	for _, student := range students {
		var studentMean = StudentMean{
			name: student.name,
			mean: (student.cLanguage + student.javaLanguage) / 2,
		}
		if studentMean.mean >= 90 {
			above90 = append(above90, studentMean)
		}
	}
	sort.Sort(above90)
	//
	for _, data := range above90 {
		fmt.Printf("%s : %f\n", data.name, data.mean)
	}
}
