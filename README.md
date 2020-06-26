# turbocharger

A tool to apply [Turbo](https://ai.googleblog.com/2019/08/turbo-improved-rainbow-colormap-for.html) [colormap](https://github.com/engelsjk/cturbo) styling to an image.

## Use

```
turbocharger -i test/shuttle.jpg -o test
```

![](test/shuttle-turbo.jpg)

## Features

Supports only PNG, JPG and JPEG images for now. Flag ```-i``` (required) for the input image filepath and flag ```-o``` (default '.') for the output image folder. Creates a new image in the output folder with the filename ```'{filename}-turbo.{ext}'```. 


## Install

```
go get github.com/engelsjk/turbocharger
```

## Notes

Shuttle image courtesy of [nasa.gov](https://www.nasa.gov/mission_pages/shuttle/shuttlemissions/sts132/multimedia/fd1/Image_Gallery_Collection_archive_6.html).
