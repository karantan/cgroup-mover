# Control Group Mover (cgroup-mover) ![gha build](https://github.com/karantan/cgroup-mover/workflows/Go/badge.svg)

`cgroup-mover` is a simple tool that helps moving child processes to a different cgroup.

The idea is to have one cgroup for the master process and a different one for its child
processes. This way we can define e.g. max CPU usage for each child process.

Idealy this should be done when a child process is spawn, but this is sometims hard to do
(e.g. php-fpm child workers).

We check for child workers every 2 seconds because:
1. We don't want to check too often and cause problems (e.g. file locking)
2. Some processes need a bit more resources at startup (first 1-2 sec) so we try being
nice to them.

## Usage

```bash
$ cgroup-mover --help
Usage of cgroup-mover:
  -new string
    	Cgroup TO which all child processes will be moved
  -old string
    	Cgroup FROM all child processes will be moved

$ cgroup-mover --old grp1 --new grp2
```
