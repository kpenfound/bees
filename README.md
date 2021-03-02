# A Bee AI Simulation

## Concepts

### World

The world is a generated space containing bee hives, bees, and flowers

### Bees

Bees are intended to travel from hive to flowers and back.  Their goal
is to collect nectar from the flowers and bring it to the hive.  In visiting
a flower, the flower will become pollinated.  A bee can only be away from it's
hive for so long.
  
### Hive

A hive is the homebase for a bee and collects nectar.  The hive can convert
nectar into new bees.

### Flowers

Flowers contain nectar for the bees to collect and need to be pollinated.  Once
pollinated, a flower can create a new flower.


## TODO
- * Nomad api
- * Job lifecycles
- * Data models
- * Bee vision
- bee ai
- run!

## Limitations
- Can't do the docker on Fedora
- Can't do envoyproxy on arm (works on arm64)
- Cant do network bridge on WSL2
- ~~For some reason can't get CNI working on Fedora~~ Kernel needs to be compiled with BRIDGE enabled
