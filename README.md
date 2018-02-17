# k8s FUSE

Manage k8s via filesystem

NOTE: At this time, it is just possible to see namespace yet.

## Usage

NOTE: The feature is not tested yet.

* Mount

```
$ go get -u github.com/yuuichi-fujioka/k8sfs
$ k8sfs -kubeconfig ~/.kube/config /mnt/k8s
```

* Check

```
$ ls -la /mnt/k8s  # Can see namespace
total 4
drwxr-xr-x 1 ubuntu  ubuntu     0  1月  1  1970 .
drwxr-xr-x 8 root    root    4096  1月 30 15:43 ..
drwxr-xr-x 0 ubuntu  ubuntu   234 12月 27 19:05 default
-rw-r--r-- 1 ubuntu  ubuntu   234 12月 27 19:05 default.yaml
drwxr-xr-x 0 ubuntu  ubuntu   242 12月 27 19:05 kube-public
-rw-r--r-- 1 ubuntu  ubuntu   242 12月 27 19:05 kube-public.yaml
drwxr-xr-x 0 ubuntu  ubuntu   241 12月 27 19:05 kube-system
-rw-r--r-- 1 ubuntu  ubuntu   241 12月 27 19:05 kube-system.yaml
$ cat /mnt/k8s/default.yaml
metadata:
  creationTimestamp: 2017-12-27T10:05:30Z
  name: default
  resourceVersion: "29"
  selfLink: /api/v1/namespaces/default
  uid: 786341fa-eaed-11e7-9088-52540051b57e
spec:
  finalizers:
  - kubernetes
status:
  phase: Active
$ ls -la /mnt/k8s/default/ # Can see directories for each resource
total 0
drwxr-xr-x 1 ubuntu  ubuntu  234  2月  8 08:53 .
drwxr-xr-x 1 ubuntu  ubuntu    0  1月  1  1970 ..
drwxr-xr-x 0 ubuntu  ubuntu    0  2月  8 08:53 po
$ ls -la /mnt/k8s/default/po # ポッド一覧
total 0
drwxr-xr-x 1 ubuntu  ubuntu     0  2月  8 08:53 .
drwxr-xr-x 0 ubuntu  ubuntu   234  2月  8 08:53 ..
-rw-r--r-- 1 ubuntu  ubuntu  2260  2月 17 10:16 nginx-8586cf59-zltvn.yaml
$ cat /mnt/k8s/default/po/nginx-8586cf59-zltvn.yaml  # podのyaml
metadata:
  creationTimestamp: 2018-02-17T01:16:08Z
  generateName: nginx-8586cf59-
  labels:
    pod-template-hash: "41427915"
    run: nginx
  name: nginx-8586cf59-zltvn
  namespace: default
  ownerReferences:
  - apiVersion: extensions/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: nginx-8586cf59
    uid: 21f6e115-1380-11e8-bf44-52540051b57e
  resourceVersion: "1213848"
  selfLink: /api/v1/namespaces/default/pods/nginx-8586cf59-zltvn
  uid: 21fab404-1380-11e8-bf44-52540051b57e
spec:
  containers:
  - image: nginx
    imagePullPolicy: Always
    name: nginx
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-8h9cs
      readOnly: true
  dnsPolicy: ClusterFirst
  nodeName: worker-2
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-8h9cs
    secret:
      defaultMode: 420
      secretName: default-token-8h9cs
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: 2018-02-17T01:16:08Z
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: 2018-02-17T01:16:12Z
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: 2018-02-17T01:16:08Z
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://2886f6794d2c404ec22916ff18e1136b73d27ed060450c1645845c04f0495952
    image: nginx:latest
    imageID: docker-pullable://nginx@sha256:98ade51c31ad73126a8fb9990efb3a3e7aba9a258d61fdac42599bf54ae955ca
    lastState: {}
    name: nginx
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: 2018-02-17T01:16:12Z
  hostIP: 192.168.0.1
  phase: Running
  podIP: 10.11.133.200
  qosClass: BestEffort
  startTime: 2018-02-17T01:16:08Z
```


* Cleanup

```
$ killall k8sfs
$ fusermount -u /mnt/k8s
```