## How to update libusb

1. clone libusb somewhere from official repo (say `~/libusb`)
2. `git checkout` to the latest stable version of libusb we used (to get patchset)
3. get list of current patches

   ```
   diff -ur  ~/libusb/libusb ~/cerberusd-go/usb/lowlevel/libusb/c/ > ~/cerberusd-patchset.diff
   ```
4. checkout the latest stable libusb
5. `mv ~/cerberusd-go/usb/lowlevel/libusb/c/ ~/cerberusd-go/usb/lowlevel/libusb/c_old`
6. `cp -r ~/libusb/libusb ~/cerberusd-go/usb/lowlevel/libusb/c`
7. `cp ~/libusb/AUTHORS ~/libusb/COPYING ~/cerberusd-go/usb/lowlevel/libusb/c`
8. `cp ~/cerberusd-go/usb/lowlevel/libusb/c_old/config.h ~/cerberusd-go/usb/lowlevel/libusb/c`
9. try to apply the patches from ~/cerberusd-patchset.diff to the new code (either manually or automatically)
11. delete unusued files, so far:

    ```
    usb/lowlevel/libusb/c/Makefile*
    usb/lowlevel/libusb/c/libusb-1.0.*
    usb/lowlevel/libusb/c/os/haiku*
    usb/lowlevel/libusb/c/os/linux_udev.c
    usb/lowlevel/libusb/c/os/netbsd_usb.c
    usb/lowlevel/libusb/c/os/null_usb.c
    usb/lowlevel/libusb/c/os/openbsd_usb.c
    usb/lowlevel/libusb/c/os/sunos_usb.c
    usb/lowlevel/libusb/c/os/sunos_usb.h
    ```

12. `rm -r ~/cerberusd-go/usb/lowlevel/libusb/c_old` when all is working fine

Note - you need to go build `go build -a` in order to "load" the new files
