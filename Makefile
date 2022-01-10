all: singleton

# ======================
# creational

.PHONY: singleton prototype factory

singleton:
	cd ./creational/singleton && go run .

prototype:
	cd ./creational/prototype && go run .

factory:
	cd ./creational/factory && go run .