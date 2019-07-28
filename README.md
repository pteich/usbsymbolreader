# USB Symbol Scanner Reader 

This Go program connects to any barcode scanner (generic symbol scanner) that runs in
IBM Hand Held mode, IBM Table Top mode or OPOS Hand Held mode. It uses an forked version of the Go USB HID library https://github.com/karalabe/hid.

It runs on OS X (tested with Mojave), Windows (tested with Windows 10 64Bit) and should run on Linux (untested).

No drivers are needed for symbol scanners. It's tested with Zebra DS3678 scanner.

While running it detects new devices or detached devices. It connects to the frist symbol scanner found.
