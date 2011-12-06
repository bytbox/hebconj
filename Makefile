include ${GOROOT}/src/Make.inc

TARG = hebconj
GOFILES = hebconj.go

include ${GOROOT}/src/Make.cmd

fmt:
	gofmt -w ${GOFILES}

