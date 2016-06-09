## FS Interface

I needed to mock file system access for a project and I found [this](http://stackoverflow.com/questions/16742331/how-to-mock-abstract-filesystem-in-go) discussion, but no good libraries for mocking - so I started building one.

### FS
The fs package contains two interfaces; Filesystem and File as well as an implementation of Filesystem that wraps the file system function the standard os library.

### FS Mock
This package creates mocks of the Filesystem for testing.

This package is already in a decent state to use for most testing but it needs a lot of work to polish up and add some edge case testing - out of space errors, bad permissions.