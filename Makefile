all: singleton prototype factory abstractfactory builder

# ======================
# creational

.PHONY: singleton prototype factory abstractfactory builder

singleton:
	cd ./creational/singleton && go run .

prototype:
	cd ./creational/prototype && go run .

factory:
	cd ./creational/factory && go run .

abstractfactory:
	cd ./creational/abstractfactory && go run .

builder:
	cd ./creational/builder && go run .

# ======================
# structural

.PHONY: adapter

adapter:
	cd ./structural/adapter && go run .
