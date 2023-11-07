# Construct
The language used for instructing the AVM (Abstract Virtual Machine).
Construct can be written in it's lowest level (Direct AVM calls) or in a higher level
that allows for OOP and more complex loops. Construct code may be different for each 
kind of AVM. For example say a node is running AVM x86 then the amount of AVM calls
availiable will be significantly larger. Although the architecture of the AVM being run by a node
does not only affect the amount of AVM calls availiable. Here is a list of all the X architectures and
all the things they affect:

AVM x86 (32GB RAM RECOMENDED):
    - 86 Different AVM Calls
    - 860 Pointers availiable for a given program
    - 86 Megabytes of Storage Availiable for a given program (Maximum)
    - 8.6 Megabytes of RAM (Allowing for 115 programs per gigabyte of the node's RAM)
    - Native Support for Higher-level Construct (e.g. construct-v1.3.0>=)
    - 86 tps (minimum)
AVM x64 (16GB RAM RECOMENDED):
    - 64 Different AVM Calls
    - 640 Pointers availiable for a given program
    - 64 Megabytes of Storage Availiable for a given program (Maximum)
    - 6.4 Megabytes of RAM (Allowing for 150 programs per gigabyte of node's RAM)
    - 64 tps (minimum)
AVM x32 - STANDARD (8GB RAM RECOMENDED):
    - 32 Different AVM Calls
    - 320 Pointers availiable for a given program
    - 3.2 Megabytes of Storage Availiable for a given program (Maximum)
    - 0.32 Megabytes of RAM (Allowing for 3125 programs per gigabyte of node's RAM)
    - 32 tps (minimum)
AVM x16 - MINIMUM (4GB RAM RECOMENDED):
    - 32 Different AVM Calls
    - 16 Pointers availiable for a given program
    - 1 Megabytes of Storage Availiable for a given program (Maximum)
    - 0.16 Megabytes of RAM (Allowing for 3125 programs per gigabyte of node's RAM)
    - 16 tps (minimum)

When compiling a higher-level of Construct or any language that compiles to the lowest level of Construct it
is recomended not to compile to AVM x16 as it may result in broken pointers and low storage. The use of 
AVM x16 is for delegating tiny-programs to nodes that aren't handling the majority of the server. AVM x32
is more of a normal version of the AVM, the AVM x32 is robust enough to allow high-level compilation to
low-level with a low-risk of pointer leaks.

AVM x64 is a less capable version and tuned down version of AVM x86. AVM x64 does not allow native support for higher-level construct
but it does allow for larger programs and with a larger library of AVM calls. As the Abstract Network gets larger, the need for scaling solutions will arise,
a promissing idea for scaling is implementing new architectures that are still backwards compatible with X architecture verions of the AVM. For example,
the HYPR architecture which would allow for the faster processing of construct programs and blocks. This would be acheieved through a new system for storing and 
accessing state data.

# X-Architecture
The X-Architecture is currently the only architecture supported by the Abstract Network. Any new Architectures must
have the ability to have backwards-compatability with the X-Architecture. From now own we will refrence the X-Architecture
as the AVMX.

The AVMX works on opcodes