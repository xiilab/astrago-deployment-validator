# AstraGo Precheck Report

- 생성 시각: **2025-11-26 09:49:55**
- kube-vip: ****
- 체크 포트: `22, 6443, 2379, 10250`

---

## Node: 10.61.3.89 (role: master)

### Ping
- 결과: **OK**
```
Ping 성공
```

### SSH
- 결과: **OK**
```
SSH 접속 성공
```

### Firewall
- 결과: **OK**
```
[ufw]
Status: inactive
```

### Ports
- 결과: **OK**
```
[22] OK: TCP접속:true, listening:true
[6443] OK: TCP접속:true, listening:true
[2379] OK: TCP접속:false, listening:true
[10250] OK: TCP접속:true, listening:true
```

### Disk
- 결과: **OK**
```
Filesystem                        Type     Size  Used Avail Use% Mounted on
tmpfs                             tmpfs    3.2G  1.6M  3.2G   1% /run
/dev/mapper/ubuntu--vg-ubuntu--lv ext4      62G  7.5G   52G  13% /
tmpfs                             tmpfs     16G     0   16G   0% /dev/shm
tmpfs                             tmpfs    5.0M     0  5.0M   0% /run/lock
/dev/vda2                         ext4     2.0G  234M  1.6G  13% /boot
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/c5b86e87-d68e-4c41-9fc0-f2b17735cf0d/volumes/kubernetes.io~projected/kube-api-access-7kjmm
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/940dc81a-7510-497d-9851-27f044a54b25/volumes/kubernetes.io~projected/kube-api-access-hdbzp
tmpfs                             tmpfs     32G  4.0K   32G   1% /var/lib/k0s/kubelet/pods/d54e22cc-0c0f-48b9-b7ed-9d62e806eb64/volumes/kubernetes.io~projected/konnectivity-agent-token
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/d54e22cc-0c0f-48b9-b7ed-9d62e806eb64/volumes/kubernetes.io~projected/kube-api-access-zj6nm
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/4903e16e-63fc-47e4-a88d-4a80e2b731fd/volumes/kubernetes.io~projected/kube-api-access-qgcbg
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/a8e046b4276768fcd434fa8c36ff2291650896b617989493d30fb1a29610f2a7/shm
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/4e8a1e209b2f01b65307c1d944c45ef50fdb244a99226d0ea459bf6e2fe35f10/shm
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/a8e046b4276768fcd434fa8c36ff2291650896b617989493d30fb1a29610f2a7/rootfs
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/7acd24b78b6cf522c91322824201b7f6bacd065c2884acc01f3ec51cfa6ed682/shm
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/df34c72b72227ebbcb09d162b3ada38d7bbce5dd1e63eefe9666e46dc60a2038/shm
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/7acd24b78b6cf522c91322824201b7f6bacd065c2884acc01f3ec51cfa6ed682/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/4e8a1e209b2f01b65307c1d944c45ef50fdb244a99226d0ea459bf6e2fe35f10/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/df34c72b72227ebbcb09d162b3ada38d7bbce5dd1e63eefe9666e46dc60a2038/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/fb1d00e1d5c1d98f60d44860240afb20aa64b3775b19c1604b6c9ba78973e9fe/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/02f56b8d386ffcf8b1bc5ba114110e08f7beef6dd7b442bb7999cd410645e342/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/226a99983ab6f0e6bb07960266d35248dc56e7b5d2b52dae2f213f7bbb692b06/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/70b27912c16eb6b6be988a22c3c8d24759073e84e16b8a151c3224cda5850e32/rootfs
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/da02c5c1-f51f-4783-a9e8-c8c7b3cff527/volumes/kubernetes.io~projected/kube-api-access-ksmhc
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/c4d59a4fddce2db44d3f336189e02be28b5b7b88976d0dd7b00410ff0da81b12/shm
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/c4d59a4fddce2db44d3f336189e02be28b5b7b88976d0dd7b00410ff0da81b12/rootfs
overlay                           overlay   62G  7.5G   52G  13% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/6e53f53617946744efd9489b1a0e64787197e0b7f29d013d3a0de38101a717d5/rootfs
tmpfs                             tmpfs    3.2G  4.0K  3.2G   1% /run/user/0
```

### resolv.conf
- 결과: **OK**
```
# This is /run/systemd/resolve/stub-resolv.conf managed by man:systemd-resolved(8).
# Do not edit.
#
# This file might be symlinked as /etc/resolv.conf. If you're looking at
# /etc/resolv.conf and seeing this text, you have followed the symlink.
#
# This is a dynamic resolv.conf file for connecting local clients to the
# internal DNS stub resolver of systemd-resolved. This file lists all
# configured search domains.
#
# Run "resolvectl status" to see details about the uplink DNS servers
# currently in use.
#
# Third party programs should typically not access this file directly, but only
# through the symlink at /etc/resolv.conf. To manage man:resolv.conf(5) in a
# different way, replace this symlink by a static file or a different symlink.
#
# See man:systemd-resolved.service(8) for details about the supported modes of
# operation for /etc/resolv.conf.

nameserver 127.0.0.53
options edns0 trust-ad
search .
```

### sudo 권한
- 결과: **OK**
```
sudo OK
```

### NTP
- 결과: **OK**
```
Local time: Wed 2025-11-26 09:49:44 KST
           Universal time: Wed 2025-11-26 00:49:44 UTC
                 RTC time: Wed 2025-11-26 00:49:44
                Time zone: Asia/Seoul (KST, +0900)
System clock synchronized: yes
              NTP service: active
          RTC in local TZ: no
```

### CPU/Memory/GPU
- 결과: **OK**
```
Architecture:                       x86_64
CPU op-mode(s):                     32-bit, 64-bit
Address sizes:                      40 bits physical, 48 bits virtual
Byte Order:                         Little Endian
CPU(s):                             4
On-line CPU(s) list:                0-3
Vendor ID:                          GenuineIntel
Model name:                         Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
CPU family:                         6
Model:                              85
Thread(s) per core:                 1
Core(s) per socket:                 1
Socket(s):                          4
Stepping:                           4
BogoMIPS:                           4190.17
Flags:                              fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon rep_good nopl xtopology cpuid tsc_known_freq pni pclmulqdq vmx ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm mpx avx512f avx512dq rdseed adx smap clflushopt clwb avx512cd avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves arat umip pku ospke md_clear arch_capabilities
Virtualization:                     VT-x
Hypervisor vendor:                  KVM
Virtualization type:                full
L1d cache:                          128 KiB (4 instances)
L1i cache:                          128 KiB (4 instances)
L2 cache:                           16 MiB (4 instances)
L3 cache:                           64 MiB (4 instances)
NUMA node(s):                       1
NUMA node0 CPU(s):                  0-3
Vulnerability Gather data sampling: Unknown: Dependent on hypervisor status
Vulnerability Itlb multihit:        Not affected
Vulnerability L1tf:                 Mitigation; PTE Inversion; VMX flush not necessary, SMT disabled
Vulnerability Mds:                  Mitigation; Clear CPU buffers; SMT Host state unknown
Vulnerability Meltdown:             Mitigation; PTI
Vulnerability Mmio stale data:      Vulnerable: Clear CPU buffers attempted, no microcode; SMT Host state unknown
Vulnerability Retbleed:             Mitigation; IBRS
Vulnerability Spec rstack overflow: Not affected
Vulnerability Spec store bypass:    Mitigation; Speculative Store Bypass disabled via prctl and seccomp
Vulnerability Spectre v1:           Mitigation; usercopy/swapgs barriers and __user pointer sanitization
Vulnerability Spectre v2:           Mitigation; IBRS, IBPB conditional, STIBP disabled, RSB filling, PBRSB-eIBRS Not affected
Vulnerability Srbds:                Not affected
Vulnerability Tsx async abort:      Mitigation; Clear CPU buffers; SMT Host state unknown

total        used        free      shared  buff/cache   available
Mem:            31Gi       840Mi        28Gi       3.0Mi       1.9Gi        30Gi
Swap:          2.0Gi          0B       2.0Gi

nvidia-smi -L 실패
```

### Internet
- 결과: **OK**
```
ping -c 1 8.8.8.8: OK
ping -c 1 google.com: OK
```

### Swap
- 결과: **OK**
```
NAME      TYPE SIZE USED PRIO
/swap.img file   2G   0B   -2
```

### Kube VIP NIC
- 결과: **FAIL**
```
master node 또는 VIP 미지정
```

---

## Node: 10.61.3.85 (role: worker)

### Ping
- 결과: **OK**
```
Ping 성공
```

### SSH
- 결과: **OK**
```
SSH 접속 성공
```

### Firewall
- 결과: **OK**
```
[ufw]
Status: inactive
```

### Ports
- 결과: **FAIL**
```
[22] OK: TCP접속:true, listening:true
[6443] FAIL: TCP접속:false, listening:false
[2379] FAIL: TCP접속:false, listening:false
[10250] OK: TCP접속:true, listening:true
```

### Disk
- 결과: **OK**
```
Filesystem                        Type     Size  Used Avail Use% Mounted on
tmpfs                             tmpfs    3.2G  1.5M  3.2G   1% /run
/dev/mapper/ubuntu--vg-ubuntu--lv ext4      62G  9.2G   50G  16% /
tmpfs                             tmpfs     16G     0   16G   0% /dev/shm
tmpfs                             tmpfs    5.0M     0  5.0M   0% /run/lock
/dev/vda2                         ext4     2.0G  320M  1.5G  18% /boot
tmpfs                             tmpfs     32G  4.0K   32G   1% /var/lib/k0s/kubelet/pods/73e07f71-71b6-4220-9df9-bea03e7e543d/volumes/kubernetes.io~projected/konnectivity-agent-token
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/ce5cbe80-a1a6-4515-b86e-97725178dd29/volumes/kubernetes.io~projected/kube-api-access-ft4z5
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/73e07f71-71b6-4220-9df9-bea03e7e543d/volumes/kubernetes.io~projected/kube-api-access-cg4q5
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/1b57d655-7f05-4b5b-81cb-61ea1f6091f6/volumes/kubernetes.io~projected/kube-api-access-bckpv
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/3334b69d-1c2b-463e-bfcc-384f4afe5f0f/volumes/kubernetes.io~projected/kube-api-access-6zq9c
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/b9d95e4f75f122b48e14d4cb162cd9722ea0561aef34356543756bff91f87b3f/shm
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/39f231b18897f5c454d142da6fe9b583b8f55c2d71f899152260b23827953dd9/shm
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/006491a7da933742622c901bcbda5f00e974362d729f4b7dc46de53d78010cee/shm
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/60bcd1313a98ab99947a70e4c565f923a16d6f1644a0a9b3e2a2b2ded6fa18c8/shm
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/b9d95e4f75f122b48e14d4cb162cd9722ea0561aef34356543756bff91f87b3f/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/60bcd1313a98ab99947a70e4c565f923a16d6f1644a0a9b3e2a2b2ded6fa18c8/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/39f231b18897f5c454d142da6fe9b583b8f55c2d71f899152260b23827953dd9/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/006491a7da933742622c901bcbda5f00e974362d729f4b7dc46de53d78010cee/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/d9423981e4c51cedfa1e00bdc049b9940cbfbbfceb796b9437c496dcb811ca9d/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/ccf4a5e70b1ffa4006612dfd3609597c9389e83f55112e944d9ef8ea36fd6f17/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/82266d287bb8fbcecd1181b69fa18cc8dca6bd548187cbab1b2eb8e4e662a11a/rootfs
overlay                           overlay   62G  9.2G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/d41b8a7c35507ca9e8ef0bc1fb0e2fc0cea390f7b89d5d293d59bce8c2ba9429/rootfs
tmpfs                             tmpfs    3.2G  8.0K  3.2G   1% /run/user/0
```

### resolv.conf
- 결과: **OK**
```
# This is /run/systemd/resolve/stub-resolv.conf managed by man:systemd-resolved(8).
# Do not edit.
#
# This file might be symlinked as /etc/resolv.conf. If you're looking at
# /etc/resolv.conf and seeing this text, you have followed the symlink.
#
# This is a dynamic resolv.conf file for connecting local clients to the
# internal DNS stub resolver of systemd-resolved. This file lists all
# configured search domains.
#
# Run "resolvectl status" to see details about the uplink DNS servers
# currently in use.
#
# Third party programs should typically not access this file directly, but only
# through the symlink at /etc/resolv.conf. To manage man:resolv.conf(5) in a
# different way, replace this symlink by a static file or a different symlink.
#
# See man:systemd-resolved.service(8) for details about the supported modes of
# operation for /etc/resolv.conf.

nameserver 127.0.0.53
options edns0 trust-ad
search .
```

### sudo 권한
- 결과: **OK**
```
sudo OK
```

### NTP
- 결과: **OK**
```
Local time: Wed 2025-11-26 09:49:48 KST
           Universal time: Wed 2025-11-26 00:49:48 UTC
                 RTC time: Wed 2025-11-26 00:49:48
                Time zone: Asia/Seoul (KST, +0900)
System clock synchronized: yes
              NTP service: active
          RTC in local TZ: no
```

### CPU/Memory/GPU
- 결과: **OK**
```
Architecture:                       x86_64
CPU op-mode(s):                     32-bit, 64-bit
Address sizes:                      40 bits physical, 48 bits virtual
Byte Order:                         Little Endian
CPU(s):                             4
On-line CPU(s) list:                0-3
Vendor ID:                          GenuineIntel
Model name:                         Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
CPU family:                         6
Model:                              85
Thread(s) per core:                 1
Core(s) per socket:                 1
Socket(s):                          4
Stepping:                           4
BogoMIPS:                           4190.17
Flags:                              fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon rep_good nopl xtopology cpuid tsc_known_freq pni pclmulqdq vmx ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm mpx avx512f avx512dq rdseed adx smap clflushopt clwb avx512cd avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves arat umip pku ospke md_clear arch_capabilities
Virtualization:                     VT-x
Hypervisor vendor:                  KVM
Virtualization type:                full
L1d cache:                          128 KiB (4 instances)
L1i cache:                          128 KiB (4 instances)
L2 cache:                           16 MiB (4 instances)
L3 cache:                           64 MiB (4 instances)
NUMA node(s):                       1
NUMA node0 CPU(s):                  0-3
Vulnerability Gather data sampling: Unknown: Dependent on hypervisor status
Vulnerability Itlb multihit:        Not affected
Vulnerability L1tf:                 Mitigation; PTE Inversion; VMX flush not necessary, SMT disabled
Vulnerability Mds:                  Mitigation; Clear CPU buffers; SMT Host state unknown
Vulnerability Meltdown:             Mitigation; PTI
Vulnerability Mmio stale data:      Vulnerable: Clear CPU buffers attempted, no microcode; SMT Host state unknown
Vulnerability Retbleed:             Mitigation; IBRS
Vulnerability Spec rstack overflow: Not affected
Vulnerability Spec store bypass:    Mitigation; Speculative Store Bypass disabled via prctl and seccomp
Vulnerability Spectre v1:           Mitigation; usercopy/swapgs barriers and __user pointer sanitization
Vulnerability Spectre v2:           Mitigation; IBRS, IBPB conditional, STIBP disabled, RSB filling, PBRSB-eIBRS Not affected
Vulnerability Srbds:                Not affected
Vulnerability Tsx async abort:      Mitigation; Clear CPU buffers; SMT Host state unknown

total        used        free      shared  buff/cache   available
Mem:            31Gi       477Mi        29Gi       2.0Mi       1.3Gi        30Gi
Swap:          2.0Gi          0B       2.0Gi

GPU 0: Tesla V100-PCIE-32GB (UUID: GPU-01380529-7586-7a37-f65e-5ea008335644)
```

### Internet
- 결과: **OK**
```
ping -c 1 8.8.8.8: OK
ping -c 1 google.com: OK
```

### Swap
- 결과: **OK**
```
NAME      TYPE SIZE USED PRIO
/swap.img file   2G   0B   -2
```

### Kube VIP NIC
- 결과: **FAIL**
```
master node 또는 VIP 미지정
```

---

## Node: 10.61.3.86 (role: worker)

### Ping
- 결과: **OK**
```
Ping 성공
```

### SSH
- 결과: **OK**
```
SSH 접속 성공
```

### Firewall
- 결과: **OK**
```
[ufw]
Status: inactive
```

### Ports
- 결과: **FAIL**
```
[2379] FAIL: TCP접속:false, listening:false
[10250] OK: TCP접속:true, listening:true
[22] OK: TCP접속:true, listening:true
[6443] FAIL: TCP접속:false, listening:false
```

### Disk
- 결과: **OK**
```
Filesystem                        Type     Size  Used Avail Use% Mounted on
tmpfs                             tmpfs    3.2G  1.4M  3.2G   1% /run
/dev/mapper/ubuntu--vg-ubuntu--lv ext4      62G  9.1G   50G  16% /
tmpfs                             tmpfs     16G     0   16G   0% /dev/shm
tmpfs                             tmpfs    5.0M     0  5.0M   0% /run/lock
/dev/vda2                         ext4     2.0G  320M  1.5G  18% /boot
tmpfs                             tmpfs     32G  4.0K   32G   1% /var/lib/k0s/kubelet/pods/135fac05-e31d-4b30-8886-418ec9e9892f/volumes/kubernetes.io~projected/konnectivity-agent-token
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/135fac05-e31d-4b30-8886-418ec9e9892f/volumes/kubernetes.io~projected/kube-api-access-9r5wz
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/f0e225c5-6d7f-479c-b650-780ee366429c/volumes/kubernetes.io~projected/kube-api-access-f4qww
tmpfs                             tmpfs     32G   12K   32G   1% /var/lib/k0s/kubelet/pods/2d467b62-3579-4403-b476-44bc9e0b13c7/volumes/kubernetes.io~projected/kube-api-access-9m8rs
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/013ce3cfb5b06eecdabb853280bfede397ee7300fb3e7b48456925912134057e/shm
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/09cb5467ebd9c6a7caacc16a78db669099dfc6bfcd0b858b3b618dd3729b5213/shm
overlay                           overlay   62G  9.1G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/09cb5467ebd9c6a7caacc16a78db669099dfc6bfcd0b858b3b618dd3729b5213/rootfs
overlay                           overlay   62G  9.1G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/013ce3cfb5b06eecdabb853280bfede397ee7300fb3e7b48456925912134057e/rootfs
overlay                           overlay   62G  9.1G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/8de9434c8817b1673010e231b1d8a0919211550587dd35a8d26cb42c846d8f39/rootfs
overlay                           overlay   62G  9.1G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/e747db7804e6822071ef1cbad322910e61c7da63cd56956f7f931e818e42eac6/rootfs
shm                               tmpfs     64M     0   64M   0% /run/k0s/containerd/io.containerd.grpc.v1.cri/sandboxes/0a214edece6fe57d038f50f3be28d7ce74802d8bb3594cdcd6765c5dd910c067/shm
overlay                           overlay   62G  9.1G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/0a214edece6fe57d038f50f3be28d7ce74802d8bb3594cdcd6765c5dd910c067/rootfs
overlay                           overlay   62G  9.1G   50G  16% /run/k0s/containerd/io.containerd.runtime.v2.task/k8s.io/25004a10ab2b3d057c6e5417f5eef33d0285dd7936cd3d76163ed350ed3030b9/rootfs
tmpfs                             tmpfs    3.2G  8.0K  3.2G   1% /run/user/0
```

### resolv.conf
- 결과: **OK**
```
# This is /run/systemd/resolve/stub-resolv.conf managed by man:systemd-resolved(8).
# Do not edit.
#
# This file might be symlinked as /etc/resolv.conf. If you're looking at
# /etc/resolv.conf and seeing this text, you have followed the symlink.
#
# This is a dynamic resolv.conf file for connecting local clients to the
# internal DNS stub resolver of systemd-resolved. This file lists all
# configured search domains.
#
# Run "resolvectl status" to see details about the uplink DNS servers
# currently in use.
#
# Third party programs should typically not access this file directly, but only
# through the symlink at /etc/resolv.conf. To manage man:resolv.conf(5) in a
# different way, replace this symlink by a static file or a different symlink.
#
# See man:systemd-resolved.service(8) for details about the supported modes of
# operation for /etc/resolv.conf.

nameserver 127.0.0.53
options edns0 trust-ad
search .
```

### sudo 권한
- 결과: **OK**
```
sudo OK
```

### NTP
- 결과: **OK**
```
Local time: Wed 2025-11-26 09:49:52 KST
           Universal time: Wed 2025-11-26 00:49:52 UTC
                 RTC time: Wed 2025-11-26 00:49:52
                Time zone: Asia/Seoul (KST, +0900)
System clock synchronized: yes
              NTP service: active
          RTC in local TZ: no
```

### CPU/Memory/GPU
- 결과: **OK**
```
Architecture:                       x86_64
CPU op-mode(s):                     32-bit, 64-bit
Address sizes:                      40 bits physical, 48 bits virtual
Byte Order:                         Little Endian
CPU(s):                             4
On-line CPU(s) list:                0-3
Vendor ID:                          GenuineIntel
Model name:                         Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
CPU family:                         6
Model:                              85
Thread(s) per core:                 1
Core(s) per socket:                 1
Socket(s):                          4
Stepping:                           4
BogoMIPS:                           4190.17
Flags:                              fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon rep_good nopl xtopology cpuid tsc_known_freq pni pclmulqdq vmx ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 hle avx2 smep bmi2 erms invpcid rtm mpx avx512f avx512dq rdseed adx smap clflushopt clwb avx512cd avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves arat umip pku ospke md_clear arch_capabilities
Virtualization:                     VT-x
Hypervisor vendor:                  KVM
Virtualization type:                full
L1d cache:                          128 KiB (4 instances)
L1i cache:                          128 KiB (4 instances)
L2 cache:                           16 MiB (4 instances)
L3 cache:                           64 MiB (4 instances)
NUMA node(s):                       1
NUMA node0 CPU(s):                  0-3
Vulnerability Gather data sampling: Unknown: Dependent on hypervisor status
Vulnerability Itlb multihit:        Not affected
Vulnerability L1tf:                 Mitigation; PTE Inversion; VMX flush not necessary, SMT disabled
Vulnerability Mds:                  Mitigation; Clear CPU buffers; SMT Host state unknown
Vulnerability Meltdown:             Mitigation; PTI
Vulnerability Mmio stale data:      Vulnerable: Clear CPU buffers attempted, no microcode; SMT Host state unknown
Vulnerability Retbleed:             Mitigation; IBRS
Vulnerability Spec rstack overflow: Not affected
Vulnerability Spec store bypass:    Mitigation; Speculative Store Bypass disabled via prctl and seccomp
Vulnerability Spectre v1:           Mitigation; usercopy/swapgs barriers and __user pointer sanitization
Vulnerability Spectre v2:           Mitigation; IBRS, IBPB conditional, STIBP disabled, RSB filling, PBRSB-eIBRS Not affected
Vulnerability Srbds:                Not affected
Vulnerability Tsx async abort:      Mitigation; Clear CPU buffers; SMT Host state unknown

total        used        free      shared  buff/cache   available
Mem:            31Gi       468Mi        28Gi       2.0Mi       2.4Gi        30Gi
Swap:          2.0Gi          0B       2.0Gi

GPU 0: Tesla V100-PCIE-32GB (UUID: GPU-39d3c0a4-2961-5058-1dda-018d815308ff)
```

### Internet
- 결과: **OK**
```
ping -c 1 8.8.8.8: OK
ping -c 1 google.com: OK
```

### Swap
- 결과: **OK**
```
NAME      TYPE SIZE USED PRIO
/swap.img file   2G   0B   -2
```

### Kube VIP NIC
- 결과: **FAIL**
```
master node 또는 VIP 미지정
```

---
