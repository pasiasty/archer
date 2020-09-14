# Gravity archer

In this game you will try to kill the other archers in the space while taking into the consideration the force of gravity of surrounding planets!

## Developer guide

1. Install buffalo
1. Install nodejs
1. Run `docker-compose -f docker-compose.deps.yml up -d` to run all dependencies (like SQL server).
1. Run `buffalo setup` to prepare your environment
1. Run `buffalo test` to ensure that all tests are passing
1. Run `buffalo dev` to run a daemon that will watch your source code, rebuild and launch it automatically. The dev instance will be served under [http://127.0.0.1:3000](http://127.0.0.1:3000) address.

After making sure that all dependencies are running you may also run the latest image of Archer (without having to rebuild anything) by executing:

`docker-compose -f docker-compose.latest.yml up -d`
