# Intro

The code here is just an experiment to a wrapper around the Open Asset Importer (assimp) library.

The model used "Bob.blend" is the same released in the assimp SDK, see the bob.source.txt for more information.

# How it works

The code load the model using assimp then convert it to pure Go representation, using the Go representation, the model is rendered using OpenGL/GLFW.

# Go Dependencies

* banthar OpenGL bindings (github.com/banthar/gl and github.com/banthar/glu).
* jteeuwen bindings for GLFW (github.com/jteeuwen/glfw).

# C Dependencies

* Open Asset Import (http://assimp.sourceforge.net/ - assimp--3.0.1270-sdk)
* GLFW (http://www.glfw.org/ - 2.7.6) -> I used the 2.7.6 version because the 2.7.5 wasn't working on my 64 bit machine.
* GLEW (http://glew.sourceforge.net/ - 1.7.0) -> I had to configure my LD_LIBRARY_PATH to search for /usr/lib64 in order to start the program.
* OpenGL (libgl1-mesa-dev/dri/glx) and GLU (libglu1-mesa-dev) -> install from your own linux distribution

I compiled assimp/glfw/glew using the standard configuration, note that GLEW installs under /usr/lib (or /usr/lib64 in my case), the others go under /usr/local.


