# turbocharger

A tool to apply [engelsjk/colormap](https://github.com/engelsjk/colormap) styling to an image.

## Example

```bash
turbocharger -i test/shuttle.jpg -o test
```

![](test/shuttle-turbo.jpg)

Creates a new image in the output folder with the filename ```'{filename}-{palette}.{ext}'```.

## Usage

```bash
turbocharger --help
```

```bash
Usage of turbocharger:
  -i string
        input filepath
  -l    list palettes
  -o string
        output dir (default ".")
  -p string
        palette (default "turbo")
```

## Features

Supports only PNG, JPG and JPEG images. 

## Install

```
go get github.com/engelsjk/turbocharger
```

## Notes

Shuttle image courtesy of [nasa.gov](https://www.nasa.gov/mission_pages/shuttle/shuttlemissions/sts132/multimedia/fd1/Image_Gallery_Collection_archive_6.html).
