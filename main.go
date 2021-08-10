package main

import (
	"os"

	"github.com/mbaraa/asu_forms/data"
)

func main() {
	formsStore := new(data.HardCodeSource)
	foo, _ := formsStore.Get("SocietyService")

	foo.ModifyFieldContent("StudentName", "مهاماد براء بشار المصري")
	foo.ModifyFieldContent("StudentId", "201910560")
	foo.ModifyFieldContent("AcademicAdvisor", "أ. هديل احمد")
	foo.ModifyFieldContent("Major", "هندسة البرمجيات")
	foo.ModifyFieldContent("Date", "08/04/2021")
	foo.ModifyFieldContent("Semester", "20212")
	foo.ModifyFieldContent("ActivityGoal", "تعزيز مهارة البرمجة عند طلاب كلية تكنلوجيا المعلومات")
	foo.ModifyFieldContent("TargetedPersonnel", "طلاب السنة الاولى")
	foo.ModifyFieldContent("ActivityTitle", "Junior Programming Contest3")
	foo.ModifyFieldContent("DeservedPoints", "2")

	final, err := foo.MakeForm()
	if err != nil {
		panic(err)
	}
	fin, _ := os.Create("foobar.png")
	fin.Write(final)
	fin.Close()

	return

}
