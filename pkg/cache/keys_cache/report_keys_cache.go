package keys

import "github.com/hung-senbox/senbox-cache-service/helper"

func StudentReportPercentage(studentID string) string {
	return string(helper.ReportServicePrefix) + "student_report_percentage:" + studentID
}

func TeacherReportPercentage(teacherID string) string {
	return string(helper.ReportServicePrefix) + "teacher_report_percentage:" + teacherID
}
