allcreational: singleton prototype factory abstractfactory builder

allstructural: adapter proxy flyweight facade decorator

all: allcreational allstructural

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

proxy:
	cd ./structural/proxy && go run .

flyweight:
	cd ./structural/flyweight && go run .

facade:
	cd ./structural/facade && go run .

decorator:
	cd ./structural/decorator && go run .	