#!/bin/zsh

qemu-system-aarch64 -m 1024 -smp 2 -cpu host -M virt,highmem=off \
-nographic \
-accel hvf \
-kernel openwrt-22.03.5-armvirt-64-Image \
-drive file=openwrt-22.03.5-armvirt-64-rootfs-squashfs.img,format=raw,if=virtio \
-append root=/dev/vda \
-device virtio-net,netdev=net0 -netdev user,id=net0,net=192.168.1.0/24,hostfwd=tcp:127.0.0.1:1122-192.168.1.1:22 \
-device virtio-net,netdev=net1 -netdev user,id=net1,net=192.0.2.0/24
