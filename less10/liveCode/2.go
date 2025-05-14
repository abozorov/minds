package main

import "fmt"

type Student struct {
	Name   string
	Grades map[string]int
}

func AddGrade(student *Student, subject string, grade int) {
	student.Grades[subject] = grade
}

func AverageGrade(student Student) float64 {
	averageGrade := 0.0
	for _, v := range student.Grades {
		averageGrade += float64(v)
	}

	return averageGrade / float64(len(student.Grades))
}

func BestSubject(student Student) string {
	if len(student.Grades) == 0 {
		return fmt.Sprintf("У студента %s нет оценок!", student.Name)
	}

	var (
		best string
		mx   int
	)
	for k, v := range student.Grades {
		if mx < v {
			best, mx = k, v
		}
	}
	return best
}

func filterTopStudents(students []Student, minAverage float64) []Student {
	retrn := make([]Student, 0, len(students))
	for _, v := range students {
		if AverageGrade(v) >= minAverage {
			retrn = append(retrn, v)
		}
	}

	return retrn
}

func main() {
	students := []Student{
		{
			Name: "Иван Петров",
			Grades: map[string]int{
				"Математика":  85,
				"Физика":      78,
				"Информатика": 92,
			},
		},
		{
			Name: "Алина Смирнова",
			Grades: map[string]int{
				"Математика":  95,
				"Физика":      88,
				"Информатика": 90,
			},
		},
		{
			Name: "Дмитрий Кузнецов",
			Grades: map[string]int{
				"Математика":  65,
				"Физика":      70,
				"Информатика": 60,
			},
		},
		{
			Name: "Екатерина Новикова",
			Grades: map[string]int{
				"Математика":  78,
				"Физика":      82,
				"Информатика": 85,
			},
		},
	}

	AddGrade(&students[2], "Programming", 100)
	AddGrade(&students[0], "Programming", 70)
	AddGrade(&students[3], "Programming", 42)

	filterTopStudents()


}
