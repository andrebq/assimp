# Open Asset Importer

This application uses the OpenAssetLibrary to import a scene and output's a gob encoded
structure (see github.com/andrebq/assimp) that can be further manipulated using just go
whithout the need for cgo packages.

The usage is very simple just call:

assimp -if path-to-your-input-file [-of path-to-dest-file.gob]

The default -of is std-out, the input file is required