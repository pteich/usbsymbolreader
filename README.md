# USB Symbol Scanner Reader 

This Go commandline utility connects to any barcode scanner (genereic symbol scanner) that runs in
IBM Hand Held mode, IBM Table Top mode or OPOS Hand Held mode. It uses an forked version of the Go USB HID library https://github.com/karalabe/hid.

It runs on OS X (tested with Mojave), Windows (tested with Windows 10 64Bit) and should run on Linux (untested).

No drivers are needed for symbol scanners. It's tested with Zebra DS3678 scanner.
