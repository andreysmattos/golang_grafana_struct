package main

import "golang_grafana_struct/passage"

func main() {
	passageService, _ := InitializePassageService()

	query := passage.GetPassageQuery{UID: "123"}
	passageService.GetPassage(query)
}
