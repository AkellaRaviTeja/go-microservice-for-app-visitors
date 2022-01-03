# docker commands

## Pull the docker images

> **docker pull mongo** \
> **docker pull mongo-express** \
> **docker pull golang**

## Create a network for mongo

> **docker network create mongo-network**

_6b22e220102f8c1b90aafed9e1a1b322f00220af456f9025aef7f5850f8fac7a_

## Check the network created

> **docker network ls**

_NETWORK ID NAME DRIVERSCOPE_ \
 _043d023d4b77 bridge bridge local_ \
 _5a0fdf7addeb host host local_ \
 _6b22e220102f mongo-network bridge local_ \
 _c3e419abfacc none null local_

## Provide the network option while running the container

> _-d : detached mode_ \
> _-p : port number(default for mongo is 27017)_ \
> _-e : environmental variable_ \
> _MONGO_INITDB_ROOT_USERNAME : mongo username_ \
> _MONGO_INITDB_ROOT_PASSWORD : mongo password_ \
> _ME_CONFIG_MONGODB_URL : mongo running url_ \
> _-name : Container name(Choose any name)_ \
> _--net : Network that was created in the above step_

> **docker run -p 27017:27017 -d** \
> **-e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password** \
> **--name mongodb --net mongo-network mongo**

_a944534ed1b1ba4c8ae58bd22d71c0b9f768c19d7a617eb4059c9ffb9a303167_

## To check whats happening inside the container log it

> **docker logs a944534ed1b1ba4c8ae58bd22d71c0b9f768c19d7a617eb4059c9ffb9a303167**

## Run the mongo express

> _ME_CONFIG_MONGODB_ADMINUSERNAME : mongodb admin username_ \
> _ME_CONFIG_MONGODB_ADMINPASSWORD : mongodd admin password_ \
> _ME_CONFIG_MONGODB_SERVER : Container name of the mongodb_

> **docker run -p 8081:8081 -d** \
> **-e ME_CONFIG_MONGODB_ADMINUSERNAME=admin** \
> **-e ME_CONFIG_MONGODB_ADMINPASSWORD=password** \
> **-e ME_CONFIG_MONGODB_SERVER=mongodb** \
> **--name mongo-express --net mongo-network mongo-express**

_b15a02422a5aa3bc9770e56cda1b82dab00afce7309cc76eefeed6dc4025c07d_

## Check the containers running

> **docker ps**

_CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES_ \
_b15a02422a5a mongo-express "tini -- /docker-ent…" 23 seconds ago Up 21 seconds 0.0.0.0:8081->8081/tcp mongo-express_ \
_a944534ed1b1 mongo "docker-entrypoint.s…" 6 minutes ago Up 6 minutes 0.0.0.0:27017->27017/tcp mongodb_

## Now hit [http://localhost:8081/] and you should see mongo-express running
