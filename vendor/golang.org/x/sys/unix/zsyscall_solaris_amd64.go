
// +build solaris,amd64

package unix

import (
	"syscall"
	"unsafe"
)

//go:cgo_import_dynamic libc_pipe pipe "libc.so"
//go:cgo_import_dynamic libc_getsockname getsockname "libsocket.so"
//go:cgo_import_dynamic libc_getcwd getcwd "libc.so"
//go:cgo_import_dynamic libc_getgroups getgroups "libc.so"
//go:cgo_import_dynamic libc_setgroups setgroups "libc.so"
//go:cgo_import_dynamic libc_wait4 wait4 "libc.so"
//go:cgo_import_dynamic libc_gethostname gethostname "libc.so"
//go:cgo_import_dynamic libc_utimes utimes "libc.so"
//go:cgo_import_dynamic libc_utimensat utimensat "libc.so"
//go:cgo_import_dynamic libc_fcntl fcntl "libc.so"
//go:cgo_import_dynamic libc_futimesat futimesat "libc.so"
//go:cgo_import_dynamic libc_accept accept "libsocket.so"
//go:cgo_import_dynamic libc___xnet_recvmsg __xnet_recvmsg "libsocket.so"
//go:cgo_import_dynamic libc___xnet_sendmsg __xnet_sendmsg "libsocket.so"
//go:cgo_import_dynamic libc_acct acct "libc.so"
//go:cgo_import_dynamic libc___makedev __makedev "libc.so"
//go:cgo_import_dynamic libc___major __major "libc.so"
//go:cgo_import_dynamic libc___minor __minor "libc.so"
//go:cgo_import_dynamic libc_ioctl ioctl "libc.so"
//go:cgo_import_dynamic libc_poll poll "libc.so"
//go:cgo_import_dynamic libc_access access "libc.so"
//go:cgo_import_dynamic libc_adjtime adjtime "libc.so"
//go:cgo_import_dynamic libc_chdir chdir "libc.so"
//go:cgo_import_dynamic libc_chmod chmod "libc.so"
//go:cgo_import_dynamic libc_chown chown "libc.so"
//go:cgo_import_dynamic libc_chroot chroot "libc.so"
//go:cgo_import_dynamic libc_close close "libc.so"
//go:cgo_import_dynamic libc_creat creat "libc.so"
//go:cgo_import_dynamic libc_dup dup "libc.so"
//go:cgo_import_dynamic libc_dup2 dup2 "libc.so"
//go:cgo_import_dynamic libc_exit exit "libc.so"
//go:cgo_import_dynamic libc_fchdir fchdir "libc.so"
//go:cgo_import_dynamic libc_fchmod fchmod "libc.so"
//go:cgo_import_dynamic libc_fchmodat fchmodat "libc.so"
//go:cgo_import_dynamic libc_fchown fchown "libc.so"
//go:cgo_import_dynamic libc_fchownat fchownat "libc.so"
//go:cgo_import_dynamic libc_fdatasync fdatasync "libc.so"
//go:cgo_import_dynamic libc_flock flock "libc.so"
//go:cgo_import_dynamic libc_fpathconf fpathconf "libc.so"
//go:cgo_import_dynamic libc_fstat fstat "libc.so"
//go:cgo_import_dynamic libc_fstatvfs fstatvfs "libc.so"
//go:cgo_import_dynamic libc_getdents getdents "libc.so"
//go:cgo_import_dynamic libc_getgid getgid "libc.so"
//go:cgo_import_dynamic libc_getpid getpid "libc.so"
//go:cgo_import_dynamic libc_getpgid getpgid "libc.so"
//go:cgo_import_dynamic libc_getpgrp getpgrp "libc.so"
//go:cgo_import_dynamic libc_geteuid geteuid "libc.so"
//go:cgo_import_dynamic libc_getegid getegid "libc.so"
//go:cgo_import_dynamic libc_getppid getppid "libc.so"
//go:cgo_import_dynamic libc_getpriority getpriority "libc.so"
//go:cgo_import_dynamic libc_getrlimit getrlimit "libc.so"
//go:cgo_import_dynamic libc_getrusage getrusage "libc.so"
//go:cgo_import_dynamic libc_gettimeofday gettimeofday "libc.so"
//go:cgo_import_dynamic libc_getuid getuid "libc.so"
//go:cgo_import_dynamic libc_kill kill "libc.so"
//go:cgo_import_dynamic libc_lchown lchown "libc.so"
//go:cgo_import_dynamic libc_link link "libc.so"
//go:cgo_import_dynamic libc___xnet_llisten __xnet_llisten "libsocket.so"
//go:cgo_import_dynamic libc_lstat lstat "libc.so"
//go:cgo_import_dynamic libc_madvise madvise "libc.so"
//go:cgo_import_dynamic libc_mkdir mkdir "libc.so"
//go:cgo_import_dynamic libc_mkdirat mkdirat "libc.so"
//go:cgo_import_dynamic libc_mkfifo mkfifo "libc.so"
//go:cgo_import_dynamic libc_mkfifoat mkfifoat "libc.so"
//go:cgo_import_dynamic libc_mknod mknod "libc.so"
//go:cgo_import_dynamic libc_mknodat mknodat "libc.so"
//go:cgo_import_dynamic libc_mlock mlock "libc.so"
//go:cgo_import_dynamic libc_mlockall mlockall "libc.so"
//go:cgo_import_dynamic libc_mprotect mprotect "libc.so"
//go:cgo_import_dynamic libc_msync msync "libc.so"
//go:cgo_import_dynamic libc_munlock munlock "libc.so"
//go:cgo_import_dynamic libc_munlockall munlockall "libc.so"
//go:cgo_import_dynamic libc_nanosleep nanosleep "libc.so"
//go:cgo_import_dynamic libc_open open "libc.so"
//go:cgo_import_dynamic libc_openat openat "libc.so"
//go:cgo_import_dynamic libc_pathconf pathconf "libc.so"
//go:cgo_import_dynamic libc_pause pause "libc.so"
//go:cgo_import_dynamic libc_pread pread "libc.so"
//go:cgo_import_dynamic libc_pwrite pwrite "libc.so"
//go:cgo_import_dynamic libc_read read "libc.so"
//go:cgo_import_dynamic libc_readlink readlink "libc.so"
//go:cgo_import_dynamic libc_rename rename "libc.so"
//go:cgo_import_dynamic libc_renameat renameat "libc.so"
//go:cgo_import_dynamic libc_rmdir rmdir "libc.so"
//go:cgo_import_dynamic libc_lseek lseek "libc.so"
//go:cgo_import_dynamic libc_setegid setegid "libc.so"
//go:cgo_import_dynamic libc_seteuid seteuid "libc.so"
//go:cgo_import_dynamic libc_setgid setgid "libc.so"
//go:cgo_import_dynamic libc_sethostname sethostname "libc.so"
//go:cgo_import_dynamic libc_setpgid setpgid "libc.so"
//go:cgo_import_dynamic libc_setpriority setpriority "libc.so"
//go:cgo_import_dynamic libc_setregid setregid "libc.so"
//go:cgo_import_dynamic libc_setreuid setreuid "libc.so"
//go:cgo_import_dynamic libc_setrlimit setrlimit "libc.so"
//go:cgo_import_dynamic libc_setsid setsid "libc.so"
//go:cgo_import_dynamic libc_setuid setuid "libc.so"
//go:cgo_import_dynamic libc_shutdown shutdown "libsocket.so"
//go:cgo_import_dynamic libc_stat stat "libc.so"
//go:cgo_import_dynamic libc_statvfs statvfs "libc.so"
//go:cgo_import_dynamic libc_symlink symlink "libc.so"
//go:cgo_import_dynamic libc_sync sync "libc.so"
//go:cgo_import_dynamic libc_times times "libc.so"
//go:cgo_import_dynamic libc_truncate truncate "libc.so"
//go:cgo_import_dynamic libc_fsync fsync "libc.so"
//go:cgo_import_dynamic libc_ftruncate ftruncate "libc.so"
//go:cgo_import_dynamic libc_umask umask "libc.so"
//go:cgo_import_dynamic libc_uname uname "libc.so"
//go:cgo_import_dynamic libc_umount umount "libc.so"
//go:cgo_import_dynamic libc_unlink unlink "libc.so"
//go:cgo_import_dynamic libc_unlinkat unlinkat "libc.so"
//go:cgo_import_dynamic libc_ustat ustat "libc.so"
//go:cgo_import_dynamic libc_utime utime "libc.so"
//go:cgo_import_dynamic libc___xnet_bind __xnet_bind "libsocket.so"
//go:cgo_import_dynamic libc___xnet_connect __xnet_connect "libsocket.so"
//go:cgo_import_dynamic libc_mmap mmap "libc.so"
//go:cgo_import_dynamic libc_munmap munmap "libc.so"
//go:cgo_import_dynamic libc___xnet_sendto __xnet_sendto "libsocket.so"
//go:cgo_import_dynamic libc___xnet_socket __xnet_socket "libsocket.so"
//go:cgo_import_dynamic libc___xnet_socketpair __xnet_socketpair "libsocket.so"
//go:cgo_import_dynamic libc_write write "libc.so"
//go:cgo_import_dynamic libc___xnet_getsockopt __xnet_getsockopt "libsocket.so"
//go:cgo_import_dynamic libc_getpeername getpeername "libsocket.so"
//go:cgo_import_dynamic libc_setsockopt setsockopt "libsocket.so"
//go:cgo_import_dynamic libc_recvfrom recvfrom "libsocket.so"

//go:linkname procpipe libc_pipe
//go:linkname procgetsockname libc_getsockname
//go:linkname procGetcwd libc_getcwd
//go:linkname procgetgroups libc_getgroups
//go:linkname procsetgroups libc_setgroups
//go:linkname procwait4 libc_wait4
//go:linkname procgethostname libc_gethostname
//go:linkname procutimes libc_utimes
//go:linkname procutimensat libc_utimensat
//go:linkname procfcntl libc_fcntl
//go:linkname procfutimesat libc_futimesat
//go:linkname procaccept libc_accept
//go:linkname proc__xnet_recvmsg libc___xnet_recvmsg
//go:linkname proc__xnet_sendmsg libc___xnet_sendmsg
//go:linkname procacct libc_acct
//go:linkname proc__makedev libc___makedev
//go:linkname proc__major libc___major
//go:linkname proc__minor libc___minor
//go:linkname procioctl libc_ioctl
//go:linkname procpoll libc_poll
//go:linkname procAccess libc_access
//go:linkname procAdjtime libc_adjtime
//go:linkname procChdir libc_chdir
//go:linkname procChmod libc_chmod
//go:linkname procChown libc_chown
//go:linkname procChroot libc_chroot
//go:linkname procClose libc_close
//go:linkname procCreat libc_creat
//go:linkname procDup libc_dup
//go:linkname procDup2 libc_dup2
//go:linkname procExit libc_exit
//go:linkname procFchdir libc_fchdir
//go:linkname procFchmod libc_fchmod
//go:linkname procFchmodat libc_fchmodat
//go:linkname procFchown libc_fchown
//go:linkname procFchownat libc_fchownat
//go:linkname procFdatasync libc_fdatasync
//go:linkname procFlock libc_flock
//go:linkname procFpathconf libc_fpathconf
//go:linkname procFstat libc_fstat
//go:linkname procFstatvfs libc_fstatvfs
//go:linkname procGetdents libc_getdents
//go:linkname procGetgid libc_getgid
//go:linkname procGetpid libc_getpid
//go:linkname procGetpgid libc_getpgid
//go:linkname procGetpgrp libc_getpgrp
//go:linkname procGeteuid libc_geteuid
//go:linkname procGetegid libc_getegid
//go:linkname procGetppid libc_getppid
//go:linkname procGetpriority libc_getpriority
//go:linkname procGetrlimit libc_getrlimit
//go:linkname procGetrusage libc_getrusage
//go:linkname procGettimeofday libc_gettimeofday
//go:linkname procGetuid libc_getuid
//go:linkname procKill libc_kill
//go:linkname procLchown libc_lchown
//go:linkname procLink libc_link
//go:linkname proc__xnet_llisten libc___xnet_llisten
//go:linkname procLstat libc_lstat
//go:linkname procMadvise libc_madvise
//go:linkname procMkdir libc_mkdir
//go:linkname procMkdirat libc_mkdirat
//go:linkname procMkfifo libc_mkfifo
//go:linkname procMkfifoat libc_mkfifoat
//go:linkname procMknod libc_mknod
//go:linkname procMknodat libc_mknodat
//go:linkname procMlock libc_mlock
//go:linkname procMlockall libc_mlockall
//go:linkname procMprotect libc_mprotect
//go:linkname procMsync libc_msync
//go:linkname procMunlock libc_munlock
//go:linkname procMunlockall libc_munlockall
//go:linkname procNanosleep libc_nanosleep
//go:linkname procOpen libc_open
//go:linkname procOpenat libc_openat
//go:linkname procPathconf libc_pathconf
//go:linkname procPause libc_pause
//go:linkname procPread libc_pread
//go:linkname procPwrite libc_pwrite
//go:linkname procread libc_read
//go:linkname procReadlink libc_readlink
//go:linkname procRename libc_rename
//go:linkname procRenameat libc_renameat
//go:linkname procRmdir libc_rmdir
//go:linkname proclseek libc_lseek
//go:linkname procSetegid libc_setegid
//go:linkname procSeteuid libc_seteuid
//go:linkname procSetgid libc_setgid
//go:linkname procSethostname libc_sethostname
//go:linkname procSetpgid libc_setpgid
//go:linkname procSetpriority libc_setpriority
//go:linkname procSetregid libc_setregid
//go:linkname procSetreuid libc_setreuid
//go:linkname procSetrlimit libc_setrlimit
//go:linkname procSetsid libc_setsid
//go:linkname procSetuid libc_setuid
//go:linkname procshutdown libc_shutdown
//go:linkname procStat libc_stat
//go:linkname procStatvfs libc_statvfs
//go:linkname procSymlink libc_symlink
//go:linkname procSync libc_sync
//go:linkname procTimes libc_times
//go:linkname procTruncate libc_truncate
//go:linkname procFsync libc_fsync
//go:linkname procFtruncate libc_ftruncate
//go:linkname procUmask libc_umask
//go:linkname procUname libc_uname
//go:linkname procumount libc_umount
//go:linkname procUnlink libc_unlink
//go:linkname procUnlinkat libc_unlinkat
//go:linkname procUstat libc_ustat
//go:linkname procUtime libc_utime
//go:linkname proc__xnet_bind libc___xnet_bind
//go:linkname proc__xnet_connect libc___xnet_connect
//go:linkname procmmap libc_mmap
//go:linkname procmunmap libc_munmap
//go:linkname proc__xnet_sendto libc___xnet_sendto
//go:linkname proc__xnet_socket libc___xnet_socket
//go:linkname proc__xnet_socketpair libc___xnet_socketpair
//go:linkname procwrite libc_write
//go:linkname proc__xnet_getsockopt libc___xnet_getsockopt
//go:linkname procgetpeername libc_getpeername
//go:linkname procsetsockopt libc_setsockopt
//go:linkname procrecvfrom libc_recvfrom

var (
	procpipe,
	procgetsockname,
	procGetcwd,
	procgetgroups,
	procsetgroups,
	procwait4,
	procgethostname,
	procutimes,
	procutimensat,
	procfcntl,
	procfutimesat,
	procaccept,
	proc__xnet_recvmsg,
	proc__xnet_sendmsg,
	procacct,
	proc__makedev,
	proc__major,
	proc__minor,
	procioctl,
	procpoll,
	procAccess,
	procAdjtime,
	procChdir,
	procChmod,
	procChown,
	procChroot,
	procClose,
	procCreat,
	procDup,
	procDup2,
	procExit,
	procFchdir,
	procFchmod,
	procFchmodat,
	procFchown,
	procFchownat,
	procFdatasync,
	procFlock,
	procFpathconf,
	procFstat,
	procFstatvfs,
	procGetdents,
	procGetgid,
	procGetpid,
	procGetpgid,
	procGetpgrp,
	procGeteuid,
	procGetegid,
	procGetppid,
	procGetpriority,
	procGetrlimit,
	procGetrusage,
	procGettimeofday,
	procGetuid,
	procKill,
	procLchown,
	procLink,
	proc__xnet_llisten,
	procLstat,
	procMadvise,
	procMkdir,
	procMkdirat,
	procMkfifo,
	procMkfifoat,
	procMknod,
	procMknodat,
	procMlock,
	procMlockall,
	procMprotect,
	procMsync,
	procMunlock,
	procMunlockall,
	procNanosleep,
	procOpen,
	procOpenat,
	procPathconf,
	procPause,
	procPread,
	procPwrite,
	procread,
	procReadlink,
	procRename,
	procRenameat,
	procRmdir,
	proclseek,
	procSetegid,
	procSeteuid,
	procSetgid,
	procSethostname,
	procSetpgid,
	procSetpriority,
	procSetregid,
	procSetreuid,
	procSetrlimit,
	procSetsid,
	procSetuid,
	procshutdown,
	procStat,
	procStatvfs,
	procSymlink,
	procSync,
	procTimes,
	procTruncate,
	procFsync,
	procFtruncate,
	procUmask,
	procUname,
	procumount,
	procUnlink,
	procUnlinkat,
	procUstat,
	procUtime,
	proc__xnet_bind,
	proc__xnet_connect,
	procmmap,
	procmunmap,
	proc__xnet_sendto,
	proc__xnet_socket,
	proc__xnet_socketpair,
	procwrite,
	proc__xnet_getsockopt,
	procgetpeername,
	procsetsockopt,
	procrecvfrom syscallFunc
)

func pipe(p *[2]_C_int) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procpipe)), 1, uintptr(unsafe.Pointer(p)), 0, 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procgetsockname)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getcwd(buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetcwd)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getgroups(ngid int, gid *_Gid_t) (n int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procgetgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func setgroups(ngid int, gid *_Gid_t) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procsetgroups)), 2, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func wait4(pid int32, statusp *_C_int, options int, rusage *Rusage) (wpid int32, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procwait4)), 4, uintptr(pid), uintptr(unsafe.Pointer(statusp)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0)
	wpid = int32(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func gethostname(buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procgethostname)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procutimes)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func utimensat(fd int, path string, times *[2]Timespec, flag int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procutimensat)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), uintptr(flag), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procfcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func futimesat(fildes int, path *byte, times *[2]Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procfutimesat)), 3, uintptr(fildes), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(times)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procaccept)), 3, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_recvmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_sendmsg)), 3, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func acct(path *byte) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procacct)), 1, uintptr(unsafe.Pointer(path)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func __makedev(version int, major uint, minor uint) (val uint64) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&proc__makedev)), 3, uintptr(version), uintptr(major), uintptr(minor), 0, 0, 0)
	val = uint64(r0)
	return
}

func __major(version int, dev uint64) (val uint) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&proc__major)), 2, uintptr(version), uintptr(dev), 0, 0, 0, 0)
	val = uint(r0)
	return
}

func __minor(version int, dev uint64) (val uint) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&proc__minor)), 2, uintptr(version), uintptr(dev), 0, 0, 0, 0)
	val = uint(r0)
	return
}

func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procioctl)), 3, uintptr(fd), uintptr(req), uintptr(arg), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func poll(fds *PollFd, nfds int, timeout int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procpoll)), 3, uintptr(unsafe.Pointer(fds)), uintptr(nfds), uintptr(timeout), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Access(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procAccess)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Adjtime(delta *Timeval, olddelta *Timeval) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procAdjtime)), 2, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chmod(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChmod)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procChroot)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Close(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procClose)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Creat(path string, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procCreat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Dup(fd int) (nfd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procDup)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	nfd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Dup2(oldfd int, newfd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procDup2)), 2, uintptr(oldfd), uintptr(newfd), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Exit(code int) {
	sysvicall6(uintptr(unsafe.Pointer(&procExit)), 1, uintptr(code), 0, 0, 0, 0, 0)
	return
}

func Fchdir(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchdir)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchmod(fd int, mode uint32) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchmod)), 2, uintptr(fd), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchmodat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(flags), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchown(fd int, uid int, gid int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchown)), 3, uintptr(fd), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFchownat)), 5, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), uintptr(flags), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fdatasync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFdatasync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Flock(fd int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFlock)), 2, uintptr(fd), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fpathconf(fd int, name int) (val int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFpathconf)), 2, uintptr(fd), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fstat(fd int, stat *Stat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFstat)), 2, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fstatvfs(fd int, vfsstat *Statvfs_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFstatvfs)), 2, uintptr(fd), uintptr(unsafe.Pointer(vfsstat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getdents(fd int, buf []byte, basep *uintptr) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetdents)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getgid() (gid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetgid)), 0, 0, 0, 0, 0, 0, 0)
	gid = int(r0)
	return
}

func Getpid() (pid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	return
}

func Getpgid(pid int) (pgid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpgid)), 1, uintptr(pid), 0, 0, 0, 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getpgrp() (pgid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetpgrp)), 0, 0, 0, 0, 0, 0, 0)
	pgid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Geteuid() (euid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGeteuid)), 0, 0, 0, 0, 0, 0, 0)
	euid = int(r0)
	return
}

func Getegid() (egid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGetegid)), 0, 0, 0, 0, 0, 0, 0)
	egid = int(r0)
	return
}

func Getppid() (ppid int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procGetppid)), 0, 0, 0, 0, 0, 0, 0)
	ppid = int(r0)
	return
}

func Getpriority(which int, who int) (n int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procGetpriority)), 2, uintptr(which), uintptr(who), 0, 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getrusage(who int, rusage *Rusage) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGetrusage)), 2, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Gettimeofday(tv *Timeval) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procGettimeofday)), 1, uintptr(unsafe.Pointer(tv)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Getuid() (uid int) {
	r0, _, _ := rawSysvicall6(uintptr(unsafe.Pointer(&procGetuid)), 0, 0, 0, 0, 0, 0, 0)
	uid = int(r0)
	return
}

func Kill(pid int, signum syscall.Signal) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procKill)), 2, uintptr(pid), uintptr(signum), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLchown)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(uid), uintptr(gid), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Link(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Listen(s int, backlog int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_llisten)), 2, uintptr(s), uintptr(backlog), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procLstat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Madvise(b []byte, advice int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMadvise)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(advice), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mkdir(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkdir)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkdirat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mkfifo(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkfifo)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mkfifoat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMkfifoat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMknod)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mknodat(dirfd int, path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMknodat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(dev), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMlock)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mlockall(flags int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMlockall)), 1, uintptr(flags), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Mprotect(b []byte, prot int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMprotect)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(prot), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Msync(b []byte, flags int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMsync)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Munlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMunlock)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Munlockall() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procMunlockall)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procNanosleep)), 2, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procOpen)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procOpenat)), 4, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), uintptr(mode), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pathconf(path string, name int) (val int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPathconf)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(name), 0, 0, 0, 0)
	val = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pause() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPause)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPread)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procPwrite)), 4, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procread)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procReadlink)), 3, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), uintptr(len(buf)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Rename(from string, to string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(from)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(to)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRename)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRenameat)), 4, uintptr(olddirfd), uintptr(unsafe.Pointer(_p0)), uintptr(newdirfd), uintptr(unsafe.Pointer(_p1)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Rmdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procRmdir)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proclseek)), 3, uintptr(fd), uintptr(offset), uintptr(whence), 0, 0, 0)
	newoffset = int64(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setegid(egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetegid)), 1, uintptr(egid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Seteuid(euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSeteuid)), 1, uintptr(euid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setgid(gid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetgid)), 1, uintptr(gid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Sethostname(p []byte) (err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSethostname)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setpgid(pid int, pgid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetpgid)), 2, uintptr(pid), uintptr(pgid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setpriority(which int, who int, prio int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSetpriority)), 3, uintptr(which), uintptr(who), uintptr(prio), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setregid(rgid int, egid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetregid)), 2, uintptr(rgid), uintptr(egid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setreuid(ruid int, euid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetreuid)), 2, uintptr(ruid), uintptr(euid), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setrlimit(which int, lim *Rlimit) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetrlimit)), 2, uintptr(which), uintptr(unsafe.Pointer(lim)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setsid() (pid int, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetsid)), 0, 0, 0, 0, 0, 0, 0)
	pid = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Setuid(uid int) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procSetuid)), 1, uintptr(uid), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Shutdown(s int, how int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procshutdown)), 2, uintptr(s), uintptr(how), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procStat)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Statvfs(path string, vfsstat *Statvfs_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procStatvfs)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(vfsstat)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Symlink(path string, link string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(link)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSymlink)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Sync() (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procSync)), 0, 0, 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Times(tms *Tms) (ticks uintptr, err error) {
	r0, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procTimes)), 1, uintptr(unsafe.Pointer(tms)), 0, 0, 0, 0, 0)
	ticks = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procTruncate)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Fsync(fd int) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFsync)), 1, uintptr(fd), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Ftruncate(fd int, length int64) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procFtruncate)), 2, uintptr(fd), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Umask(mask int) (oldmask int) {
	r0, _, _ := sysvicall6(uintptr(unsafe.Pointer(&procUmask)), 1, uintptr(mask), 0, 0, 0, 0, 0)
	oldmask = int(r0)
	return
}

func Uname(buf *Utsname) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procUname)), 1, uintptr(unsafe.Pointer(buf)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Unmount(target string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(target)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procumount)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUnlink)), 1, uintptr(unsafe.Pointer(_p0)), 0, 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Unlinkat(dirfd int, path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUnlinkat)), 3, uintptr(dirfd), uintptr(unsafe.Pointer(_p0)), uintptr(flags), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Ustat(dev int, ubuf *Ustat_t) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUstat)), 2, uintptr(dev), uintptr(unsafe.Pointer(ubuf)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func Utime(path string, buf *Utimbuf) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procUtime)), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(buf)), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_bind)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_connect)), 3, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procmmap)), 6, uintptr(addr), uintptr(length), uintptr(prot), uintptr(flag), uintptr(fd), uintptr(pos))
	ret = uintptr(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func munmap(addr uintptr, length uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procmunmap)), 2, uintptr(addr), uintptr(length), 0, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_sendto)), 6, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = e1
	}
	return
}

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_socket)), 3, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&proc__xnet_socketpair)), 4, uintptr(domain), uintptr(typ), uintptr(proto), uintptr(unsafe.Pointer(fd)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procwrite)), 3, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0, 0, 0)
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&proc__xnet_getsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, _, e1 := rawSysvicall6(uintptr(unsafe.Pointer(&procgetpeername)), 3, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procsetsockopt)), 5, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procrecvfrom)), 6, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = e1
	}
	return
}
