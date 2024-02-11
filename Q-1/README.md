### Problem statement 1: Drop packets using eBPF

<b> PS: </b> Write an eBPF code to drop the TCP packets on a port

<b>Solution:</b>
I have written a simple eBPF code to drop the TCP packets on a port. The code is file ``ebpf-drop.c``

Extra Information:
- The eBPF code is written in C language and it is compiled using clang compiler.
- The eBPF code is loaded into the kernel using the bpf() system call.
- The eBPF code is used to drop the TCP packets on a port.
- To test this solution you can use netcat to send the TCP packets to the port and check the packets are dropped or not.
- Create a script to run this code.<br>
Script: 
```
compile:
	clang -I/usr/include/x86_64-linux-gnu -O2 -g -Wall -target bpf -c xdp_drop.c -o xdp_drop.o

load:
	sudo ip link set lo xdpgeneric obj xdp_drop.o sec tcp_drop

unload:
	sudo ip link set lo xdpgeneric off

status:
	sudo ./xdp-tools/xdp-loader/xdp-loader status
  ```

<b>Help:</b>
- https://www.tigera.io/learn/guides/ebpf/ebpf-xdp/
- ChatGpt ( AI Tool )
- Google Search 👨‍💻