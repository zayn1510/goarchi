migrate:
	go run app/migrate/migrate.go
model:
	go run cmd/make.go model $(name)
controller:
	go run cmd/make.go controller $(name)
service:
	go run cmd/make.go service $(name)
request:
	go run cmd/make.go request $(name)
resource:
	go run cmd/make.go resource $(name)
seeder:
	go run cmd/make.go seeder $(name)
create-seed:
	go run cmd/make.go create-seed $(name)
all:
	go run cmd/make.go all $(name)