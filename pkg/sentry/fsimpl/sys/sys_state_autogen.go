// automatically generated by stateify.

package sys

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *dirRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.dirRefs"
}

func (x *dirRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (x *dirRefs) beforeSave() {}

func (x *dirRefs) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
}

func (x *dirRefs) afterLoad() {}

func (x *dirRefs) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
}

func (x *kcovInode) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.kcovInode"
}

func (x *kcovInode) StateFields() []string {
	return []string{
		"InodeAttrs",
		"InodeNoopRefCount",
		"InodeNotDirectory",
		"InodeNotSymlink",
		"implStatFS",
	}
}

func (x *kcovInode) beforeSave() {}

func (x *kcovInode) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeAttrs)
	m.Save(1, &x.InodeNoopRefCount)
	m.Save(2, &x.InodeNotDirectory)
	m.Save(3, &x.InodeNotSymlink)
	m.Save(4, &x.implStatFS)
}

func (x *kcovInode) afterLoad() {}

func (x *kcovInode) StateLoad(m state.Source) {
	m.Load(0, &x.InodeAttrs)
	m.Load(1, &x.InodeNoopRefCount)
	m.Load(2, &x.InodeNotDirectory)
	m.Load(3, &x.InodeNotSymlink)
	m.Load(4, &x.implStatFS)
}

func (x *kcovFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.kcovFD"
}

func (x *kcovFD) StateFields() []string {
	return []string{
		"FileDescriptionDefaultImpl",
		"NoLockFD",
		"vfsfd",
		"inode",
		"kcov",
	}
}

func (x *kcovFD) beforeSave() {}

func (x *kcovFD) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.FileDescriptionDefaultImpl)
	m.Save(1, &x.NoLockFD)
	m.Save(2, &x.vfsfd)
	m.Save(3, &x.inode)
	m.Save(4, &x.kcov)
}

func (x *kcovFD) afterLoad() {}

func (x *kcovFD) StateLoad(m state.Source) {
	m.Load(0, &x.FileDescriptionDefaultImpl)
	m.Load(1, &x.NoLockFD)
	m.Load(2, &x.vfsfd)
	m.Load(3, &x.inode)
	m.Load(4, &x.kcov)
}

func (x *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.FilesystemType"
}

func (x *FilesystemType) StateFields() []string {
	return []string{}
}

func (x *FilesystemType) beforeSave() {}

func (x *FilesystemType) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *FilesystemType) afterLoad() {}

func (x *FilesystemType) StateLoad(m state.Source) {
}

func (x *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.filesystem"
}

func (x *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
	}
}

func (x *filesystem) beforeSave() {}

func (x *filesystem) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Filesystem)
	m.Save(1, &x.devMinor)
}

func (x *filesystem) afterLoad() {}

func (x *filesystem) StateLoad(m state.Source) {
	m.Load(0, &x.Filesystem)
	m.Load(1, &x.devMinor)
}

func (x *dir) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.dir"
}

func (x *dir) StateFields() []string {
	return []string{
		"dirRefs",
		"InodeAttrs",
		"InodeNoDynamicLookup",
		"InodeNotSymlink",
		"InodeDirectoryNoNewChildren",
		"OrderedChildren",
		"locks",
		"dentry",
	}
}

func (x *dir) beforeSave() {}

func (x *dir) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.dirRefs)
	m.Save(1, &x.InodeAttrs)
	m.Save(2, &x.InodeNoDynamicLookup)
	m.Save(3, &x.InodeNotSymlink)
	m.Save(4, &x.InodeDirectoryNoNewChildren)
	m.Save(5, &x.OrderedChildren)
	m.Save(6, &x.locks)
	m.Save(7, &x.dentry)
}

func (x *dir) afterLoad() {}

func (x *dir) StateLoad(m state.Source) {
	m.Load(0, &x.dirRefs)
	m.Load(1, &x.InodeAttrs)
	m.Load(2, &x.InodeNoDynamicLookup)
	m.Load(3, &x.InodeNotSymlink)
	m.Load(4, &x.InodeDirectoryNoNewChildren)
	m.Load(5, &x.OrderedChildren)
	m.Load(6, &x.locks)
	m.Load(7, &x.dentry)
}

func (x *cpuFile) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.cpuFile"
}

func (x *cpuFile) StateFields() []string {
	return []string{
		"implStatFS",
		"DynamicBytesFile",
		"maxCores",
	}
}

func (x *cpuFile) beforeSave() {}

func (x *cpuFile) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.implStatFS)
	m.Save(1, &x.DynamicBytesFile)
	m.Save(2, &x.maxCores)
}

func (x *cpuFile) afterLoad() {}

func (x *cpuFile) StateLoad(m state.Source) {
	m.Load(0, &x.implStatFS)
	m.Load(1, &x.DynamicBytesFile)
	m.Load(2, &x.maxCores)
}

func (x *implStatFS) StateTypeName() string {
	return "pkg/sentry/fsimpl/sys.implStatFS"
}

func (x *implStatFS) StateFields() []string {
	return []string{}
}

func (x *implStatFS) beforeSave() {}

func (x *implStatFS) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *implStatFS) afterLoad() {}

func (x *implStatFS) StateLoad(m state.Source) {
}

func init() {
	state.Register((*dirRefs)(nil))
	state.Register((*kcovInode)(nil))
	state.Register((*kcovFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*dir)(nil))
	state.Register((*cpuFile)(nil))
	state.Register((*implStatFS)(nil))
}