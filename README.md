# Realms like Minecraft hosting service

## Features

- Switch between all versions of Minecraft server
- Switch between multiple worlds (at least they have to be compatible with selected server version)
- Customize selected server: whitelisted users, message of the day, thumbnail
- Upload worlds from other servers
- Download the worlds as zip files
- Customize Minecraft server images with plugins (BuildCraft, mini games)

## Components

- Imaginarium - prepares the gallery of server images
- Vault - hosts the worlds
- Server - runs the images against the worlds with defined configuration

## Resources

All resources could be modelled with Kubernetes CRDs and interaction with them could be performed in terms of regular API verbs.

### ServerImageTemplate

The version specific template of the Dockerfile the final instance of the ServerImage will be built from. Some server versions may have platform specific components or just to have some plugins preinstalled. In same time some generic flow might be extracted from for the bulk of images: JRE version, jar file URL etc.

Properties:

- Name - the user friendly name of the template that will be used during the build
- FileContent - raw file content with template statements
- Values - input values for template building

### ServerImage

- Name - user friendly name of the image that will be used during the current server startup
- ServerImageTemplate - the template to be used for building the image
- Values - template arguments

....
_TBD_

## Challenges

- Resources migration during new releases
