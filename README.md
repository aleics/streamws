# streamws
Learning project using Websockets for sending media files throw HTTP by [Aleix Casanovas](https://github.com/aleics).
This repository is divided in two parts: client and server. Where in every part, different functionalities will be used.

## streamwsimage
This is the basic part of the project. A websocket channel will be used for sending images to the client.
Multiple frames from a video were generated using [ffmpeg](http://ffmpeg.org/). This multiple images will be the ones used.

On the client side, the user can set the period of the transmition of images (from 10ms to 200ms).

## streamwsvideo
Implementing...

## annotations
The test video used is a 720P24 downloaded [here](https://mango.blender.org/).