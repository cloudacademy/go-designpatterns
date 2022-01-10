all: singleton

# ======================
# creational

.PHONY: singleton prototype factory builder

singleton:
	cd ./creational/singleton && go run .

prototype:
	cd ./creational/prototype && go run .

factory:
	cd ./creational/factory && go run .

builder:
	cd ./creational/builder && go run .