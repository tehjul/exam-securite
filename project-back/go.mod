module main

go 1.12

require (
	github.com/gorilla/mux v1.8.0
	github.com/rs/cors v1.8.0
)

replace (
	github.com/charleslgn/exam-securite/project-back/connection => "./connection"
	github.com/charleslgn/exam-securite/project-back/customhttp => "./customhttp"
)