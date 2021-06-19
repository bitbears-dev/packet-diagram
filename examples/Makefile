.PHONY: build clean e2e-test test

pdfiles=$(wildcard *.pd)
svgfiles=$(patsubst %.pd,%.svg,$(pdfiles))

build: $(svgfiles)
	for dir in $(subdirs); do $(MAKE) -C "$$dir" $@; done

clean:
	rm -f *.svg

generate-examples: $(svgfiles)

%.svg: %.pd
	../cmd/packet-diagram/packet-diagram -i $< > $@
