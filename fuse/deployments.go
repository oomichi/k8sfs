package fuse

import (
	"log"
	"strings"

	"github.com/yuuichi-fujioka/k8sfs/k8s"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"

	v1beta1 "k8s.io/api/extensions/v1beta1"
)

type deploymentsDir struct {
	nodefs.File
	defaultDir
	Namespace string
}

func NewDeploymentsDir(ns string) (string, *deploymentsDir) {

	return "deploy", &deploymentsDir{
		File:       nodefs.NewDefaultFile(),
		defaultDir: NewDefaultDir(),
		Namespace:  ns,
	}
}

func (f *deploymentsDir) GetAttr(out *fuse.Attr) fuse.Status {
	SetAttrTimeCluster(out)
	out.Size = 4096 // block size?
	if readOnlyMode {
		out.Mode = fuse.S_IFDIR | 0555
	} else {
		out.Mode = fuse.S_IFDIR | 0755
	}
	return fuse.OK
}

func (f *deploymentsDir) GetFile() nodefs.File {
	return f
}

func (f *deploymentsDir) GetDir(name string) DirEntry {
	if name == "" {
		return f
	}

	names := strings.Split(name, "/")

	for k, child := range f.dirs {
		if k == names[0] {
			return child.GetDir(strings.Join(names[1:], "/"))
		}
	}

	return nil
}

func (f *deploymentsDir) Unlink(name string) (code fuse.Status) {
	log.Printf("Unlink: %s at %s", name, "deploy")
	// TODO
	deployName := strings.TrimSuffix(name, ".yaml")
	err := k8s.DeleteDeployment(f.Namespace, deployName)
	if err != nil {
		return fuse.EIO
	}
	return fuse.OK
}

func (f *deploymentsDir) Mkdir(name string, mode uint32) fuse.Status {
	log.Printf("Mkdir: %s at %s", name, "deploy")
	// TODO
	return fuse.ENOSYS
}

func (f *deploymentsDir) Rmdir() (code fuse.Status) {
	log.Printf("Rmdir: %s", "deploy")
	// TODO
	return fuse.ENOSYS
}

func (f *deploymentsDir) Create(name string, flags uint32, mode uint32) (file nodefs.File, code fuse.Status) {
	log.Printf("Create: %s on %s with 0x%x 0x%x", name, "deploy", flags, mode)
	// TODO
	return nil, fuse.ENOSYS
}

func (f *deploymentsDir) AddDeployment(deploy *v1beta1.Deployment) {
	if !f.UpdateDeployment(deploy) {
		newFile := NewDeploymentFile(deploy)
		f.files[newFile.Name] = newFile
	}
}

func (f *deploymentsDir) UpdateDeployment(deploy *v1beta1.Deployment) (updated bool) {

	updated = false

	name := deploy.Name
	for _, file := range f.files {
		if file.Name == name+".yaml" {
			UpdateDeploymentFile(file, deploy)
			updated = true
			break
		}
	}
	return
}

func (f *deploymentsDir) DeleteDeployment(deploy *v1beta1.Deployment) {

	name := deploy.Name

	delete(f.dirs, name)
	delete(f.files, name+".yaml")

}
