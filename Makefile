SUBDIRS := $(wildcard modules/*)
TOPTARGETS := all clean

.PHONY: $(TOPTARGETS) $(SUBDIRS)
.DEFAULT_GOAL := all

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)
