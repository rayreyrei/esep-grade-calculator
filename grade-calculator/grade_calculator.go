package esepunittests

//import "fmt"

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	//fmt.Println(numericalGrade)

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		newAssignment := Grade{Name: name, Grade: grade, Type: gradeType}
		//fmt.Println(newAssignment)
		gc.assignments = append(gc.assignments, newAssignment)
		//fmt.Println(gc.assignments)
		//gc.assignments = append(gc.assignments, Grade{
		//	Name:  name,
		//	Grade: grade,
		//	Type:  Assignment,
		//})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	//fmt.Println(gc.assignments)
	assignment_average := computeAverage(gc.assignments)
	//fmt.Println(assignment_average)
	exam_average := computeAverage(gc.exams)
	// Changed from exams to essays
	essay_average := computeAverage(gc.essays)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	sum := 0
	//fmt.Printf("%+v\n", grades)
	//fmt.Println(grades[0].Grade)

	// For Loop now calculates the grade value
	for i := 0; i < len(grades); i++ {
		//fmt.Println(grade)
		//fmt.Println(name)
		sum += grades[i].Grade
		//fmt.Println(sum)
	}

	return sum / len(grades)
}
