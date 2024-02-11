### Problem statement 2: Drop packets only for a given process

<b> PS: </b> Write an eBPF code to allow traffic only at a specific TCP port (default 4040) for a given process name (for e.g, "myprocess"). All the traffic to all other ports for only that process should be dropped.

<b>Solution:</b></br>
Approch is using hash map to get the PID of the process and then using the PID to get the process name and then checking the process name with the given process name and then checking the port number and then dropping the packet if the process name and port number is not matching with the given process name and port number.

<b>Extra Information:</b>
- Before running this script, you need to install the github.com/iovisor/gobpf library:
  ```
  go get github.com/iovisor/gobpf
  ```
  Compile and run the Go script:
  ```
  go run main.go
  ```
- The eBPF program drops TCP packets destined for port 4040.
<b>Help:</b>
- https://www.datadoghq.com/blog/ebpf-guide/
- ChatGpt ( AI Tool )
- Google Search 👨‍💻