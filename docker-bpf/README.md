# BPF enabled box

Shy box to try out syscall profiling...

## Hello World
After the provision has finished (and the reboot has taken place), start the kernelscope container.

```
$ docker run -d --name kernelscope --privileged -v /lib/modules:/lib/modules:ro -v /usr/src:/usr/src:ro -v /etc/localtime:/etc/localtime:ro -p 8080:8080 qnib/kernelscope
```
