PREFIX=GOPATH=$(PWD)
COMPILER=go
INDEXTRACKER_OUTNAME=indextracker
TRACKERMANAGER_OUTNAME=./cmd/trackermanager

all: submodules indextracker trackermanager
	$(info trackermanager and indexmanager already compiled; not recompiling. use 'make clear && make' to create new binaries)	
indextracker:
	$(PREFIX) $(COMPILER) build -o $(INDEXTRACKER_OUTNAME)
trackermanager:
	$(PREFIX) $(COMPILER) build $(TRACKERMANAGER_OUTNAME)
submodules:
	git submodule init
	git submodule update

clear:
	rm indextracker trackermanager
