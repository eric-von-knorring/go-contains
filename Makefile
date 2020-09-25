
include config.mk

all: options clean build

options:
	@echo Build options:
	@echo " TARGET   = ${TARGET}"
	@echo " BINARY   = ${BINARY}"

build:
	@echo Building binary
	@go build -o ./${TARGET}/${BINARY}

clean:
	@echo Cleaning
	rm -rf ./${TARGET}/

install: all
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f ./${TARGET}/${BINARY} ${DESTDIR}${PREFIX}/bin
	chmod 755 ${DESTDIR}${PREFIX}/bin/${BINARY}
	mkdir -p ${DESTDIR}${MANPREFIX}/man1
	sed "s/VERSION/${VERSION}/g" < contains.1 > ${DESTDIR}${MANPREFIX}/man1/contains.1
	chmod 644 ${DESTDIR}${MANPREFIX}/man1/contains.1

uninstall:
	rm -f ${DESTDIR}${PREFIX}/bin/${BINARY}\
		${DESTDIR}${MANPREFIX}/man1/contains.1

.PHONY: all options clean dist install uninstall
