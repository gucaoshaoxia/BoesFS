# kbuild trick to avoid linker error. Can be omitted if a module is built.
obj- := dummy.o



# Tell kbuild to always build the programs
always := user_prog.o

EXTRA_CFLAGS += -I$(PWD)/include -I$(objtree)/samples/bpf

HOSTCFLAGS += -fPIC -I$(objtree)/usr/include -I$(PWD)/include
HOSTCFLAGS += -I$(objtree)/samples/bpf -I$(objtree)/tools/lib -I$(objtree)/tools/lib/bpf
HOSTLOADLIBES_sandfs += -lelf -lcap -lpthread

# Allows pointing LLC/CLANG to a LLVM backend with bpf support, redefine on cmdline:
#  make samples/bpf/ LLC=~/git/llvm/build/bin/llc CLANG=~/git/llvm/build/bin/clang
LLC ?= llc
CLANG ?= clang

# Trick to allow make to be run from this directory
all:
	$(MAKE) -C /lib/modules/`uname -r`/build $$PWD/

clean:
	$(MAKE) -C /lib/modules/`uname -r`/build M=$$PWD clean
	rm -f src/*.o xdp/*.o *.a *.so sandfs
	rm -f src/libbpf.c src/bpf_load.c

# Verify LLVM compiler tools are available and bpf target is supported by llc
.PHONY: verify_cmds verify_target_bpf $(CLANG) $(LLC)

verify_cmds: $(CLANG) $(LLC)
	@for TOOL in $^ ; do \
		if ! (which -- "$${TOOL}" > /dev/null 2>&1); then \
			echo "*** ERROR: Cannot find LLVM tool $${TOOL}" ;\
			exit 1; \
		else true; fi; \
	done

verify_target_bpf: verify_cmds
	@if ! (${LLC} -march=bpf -mattr=help > /dev/null 2>&1); then \
		echo "*** ERROR: LLVM (${LLC}) does not support 'bpf' target" ;\
		echo "   NOTICE: LLVM version >= 3.7.1 required" ;\
		exit 2; \
	else true; fi

$(src)/*.c: verify_target_bpf

# asm/sysreg.h - inline assembly used by it is incompatible with llvm.
# But, there is no easy way to fix it, so just exclude it since it is
# useless for BPF samples.
$(obj)/%.o: $(src)/%.c
	$(CLANG) $(NOSTDINC_FLAGS) $(LINUXINCLUDE) $(EXTRA_CFLAGS) \
		-D__KERNEL__ -D__ASM_SYSREG_H -Wno-unused-value -Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-O2 -emit-llvm -c $< -o -| $(LLC) -march=bpf -filetype=obj -o $@
