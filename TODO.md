
#### Use the following process for image crunching:
```
        runString = {
            "jpeg": u"jpegoptim" + exe + " -f --strip-all '%(file)s'",
            "png": u"optipng" + exe + " -force -o7 '%(file)s'&&advpng" + exe + " -z4 '%(file)s' && pngcrush -rem gAMA -rem alla -rem cHRM -rem iCCP -rem sRGB -rem time '%(file)s' '%(file)s.bak' && mv '%(file)s.bak' '%(file)s'"
        }
```
