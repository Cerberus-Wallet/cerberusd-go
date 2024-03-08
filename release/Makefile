.PHONY: windows linux macos copy gpg clean

all: windows linux macos copy

linux:
	make -C linux all

macos:
	make -C macos all

windows:
	make -C windows all

copy:
	rm -rf installers
	mkdir installers
	cp macos/build/cerberus-bridge*.pkg* installers
	cp windows/build/cerberus-bridge*.exe* installers
	cp linux/build/*.deb linux/build/*.rpm installers
	rm -rf binaries
	mkdir binaries
	cp macos/build/cerberusd binaries/cerberusd-darwin-universal
	cp windows/build/cerberusd-32b.exe binaries/cerberusd-windows-386.exe
	cp windows/build/cerberusd-64b.exe binaries/cerberusd-windows-amd64.exe
	cp linux/build/cerberusd-linux-* binaries/

clean:
	rm -rf installers binaries macos/build linux/build windows/build
