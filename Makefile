allcreational: singleton prototype factory abstractfactory builder

allstructural: adapter proxy flyweight facade decorator composite bridge

allbehavioral: iterator templatemethod visitor strategy observer command chainofresponsibility mediator state

all: allcreational allstructural allbehavioral 

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

composite:
	cd ./structural/composite && go run .

bridge:
	cd ./structural/bridge && go run .

# ======================
# behavioral

iterator:
	cd ./behavioral/iterator && go run .

templatemethod:
	cd ./behavioral/templatemethod && go run .

chainofresponsibility:
	cd ./behavioral/chainofresponsibility && go run .

memento:
	cd ./behavioral/memento && go run .

state:
	cd ./behavioral/state && go run .

command:
	cd ./behavioral/command && go run .

mediator:
	cd ./behavioral/mediator && go run .

observer:
	cd ./behavioral/observer && go run .

strategy:
	cd ./behavioral/strategy && go run .

visitor:
	cd ./behavioral/visitor && go run .
