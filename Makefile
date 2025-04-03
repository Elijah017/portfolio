CMDDIR=./cmd
BINDIR=./bin

CMDS=$(shell find $(CMDDIR) -mindepth 1 -maxdepth 1 -type d -printf "%f\n")

.PHONY: all clean $(CMDS)
all: $(CMDS)

$(CMDS):
	go build -o $(BINDIR)/$@ $(CMDDIR)/$@/main.go

clean:
	rm -f $(BINDIR)/*
