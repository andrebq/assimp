# What is this?

The Open Asset Import (http://assimp.sourceforge.net/) is a library capable of reading a variety of 3D formats and present them in a single API (aiScene).

This project contains wrappers for that library and also a similar api (assimp.Scene) to manipulate the data inside a Go program.

The Go structures are implemented in pure go, so you can use it even without the Open Asset Import lib, only the package assimp/conv need the library.

# How to use?

You can use your 3D models in two ways:

* use the assimp binary (assimp/assimp) to convert your models to a pure go representation (using gob) and then inside your 3d application use the packages (assimp/enc, assimp) to load and manipulate the model.

* in your 3d application import the (assimp/conv assimp) packages to convert and manipulate the model.

# Do you have any samples?

You can read the unit tests for each package to see how the API is used or you can look at https://github.com/andrebq/exp and see a 3d application that can display your models (the application is limited to small models at this moment.)

# Can I help?

Yes, just do the usual fork/edit/request thing.

# Who is responsible for this?

If you want to know more about me, just go to: http://resume.amoraes.info/en-US/ there you can see my profiles and link to my personal blog.

# Where I can have more information about this project?

Just check the commit messages and there is a possibility that I will post something at http://amoraes.info/
