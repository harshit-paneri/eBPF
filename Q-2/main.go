package main

import (
	"fmt"
	"os"
	"unsafe"

	bpf "github.com/iovisor/gobpf/bcc"
)

const program = `
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/sched.h>

BPF_HASH(pid_filter, u32, char);

int block_port(struct xdp_md *ctx) {
    // Get the packet data and the end of the packet data
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;

    // Get the ethernet header, IP header, and TCP header
    struct iphdr *ip = data + sizeof(struct ethhdr);
    struct tcphdr *tcph = data + sizeof(struct ethhdr) + sizeof(struct iphdr);

    u32 pid = bpf_get_current_pid_tgid();
    char comm[TASK_COMM_LEN];

    bpf_get_current_comm(&comm, sizeof(comm));

    bpf_trace_printk("Debug message: variable = %d\n", comm);

    if (tcph->dest != 4040) {
        // Drop traffic for other ports
        if (pid_filter.lookup(&pid) && bpf_strncmp(comm, "myprocess", TASK_COMM_LEN) == 0 {
            return XDP_DROP;
        }
    }

    return XDP_PASS;
}
`

func main() {
	m := bpf.NewModule(program, []string{})
	defer m.Close()

	fn, err := m.Load("block_port", bpf.BPF_PROG_TYPE_XDP, 1, 65536)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BPF program: %v\n", err)
		os.Exit(1)
	}

	err = m.AttachXDP("lo", fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach BPF program: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("eBPF program attached. Press Ctrl+C to stop.")

	// Simulate adding a PID to the filter table
	pid := uint32(5878)
	key := unsafe.Pointer(&pid)
	value := unsafe.Pointer(&pid)
	m.UpdateElement("pid_filter", key, value)

	select {}
}
