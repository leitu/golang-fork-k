# Overview

## Why it comes from

This project is inspired by the keynote in kubecon Seattle.
You can find the link from https://kccna18.sched.com/event/GsxY/keynote-developing-kubernetes-services-at-airbnb-scale-melanie-cebula-software-engineer-airbnb

and Video here https://youtu.be/ytu3aUCwlSg

## How it works

I think k tool is using javascript. .. right now, it only works with same folder structure based on Melanie's Keynote, you can checkout the keynote. Regarding build and deploy, that is purely through the docker and kubectl command level,somehow it hardcoded right now, accept values will be in TODO list.

```
$ make
$ ./k
Usage:
  k [command]

Available Commands:
  build       build the docker image
  deploy      deploy environment
  generate    Generate all folders
  help        Help about any command

Flags:
  -h, --help   help for k

Use "k [command] --help" for more information about a command.
```

## TODO

- [ ] kubectl environment

- [ ] cmd.exec