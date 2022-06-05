package httpstatus

func Get_http_status(student_name string) (int, error) {
	if student_name == "" {
		return 403, nil
	}

	return 200, nil
}

func Create_http_status(student_name string, id string, point string) (int, error) {
	if id == "" || student_name == "" || point == "" {
		return 403, nil
	}

	return 201, nil
}

func Update_http_status(id string, point string) (int, error) {
	if id == "" || point == "" {
		return 400, nil
	}
	return 200, nil
}

func Delete_http_status(id int) (int, error) {
	if id <= 0 {
		return 400, nil
	}

	return 204, nil
}
