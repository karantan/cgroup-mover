# Control Group Mover (cgroup-mover) ![gha build](https://github.com/karantan/cgroup-mover/workflows/Go/badge.svg)

`cgroup-mover` is a simple tool that helps moving processes from one cgroup to a another
cgroup.

The idea is to have one cgroup for processes that you want to limit because they consume
too much server resources.

Idealy this should be done when a child process is spawn, but this is sometims hard to do
(e.g. php-fpm child workers).

That's why we check for child workers every 5 seconds because:
1. We don't want to check too often and cause problems (e.g. file locking)
2. Some processes need a bit more resources at startup (first 1-2 sec) so we try being
nice to them.

## Usage

```bash
$ cgroup-mover --help
Usage of cgroup-mover:
  -new string
    	Cgroup TO which all child processes will be moved
  -user string
      User of which all processes will be moved (e.g. foo_bar)
  -old string
    	Cgroup FROM all child processes will be moved

$ cgroup-mover --old grp1 --new grp2
```
