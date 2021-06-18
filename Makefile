.PHONY: build clean e2e-test test

subdirs=cmd/packet-diagram

build: 
	for dir in $(subdirs); do $(MAKE) -C "$$dir" $@; done

test:
	for dir in $(subdirs); do $(MAKE) -C "$$dir" $@; done

e2e-test:
	for dir in $(subdirs); do $(MAKE) -C "$$dir" $@; done

clean:
	for dir in $(subdirs); do $(MAKE) -C "$$dir" $@; done
