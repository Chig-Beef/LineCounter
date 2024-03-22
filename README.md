# LineCounter
This is a simple program that takes a few arguments.
First, give the program the file type you want to search for.
And example would be `go`.
This would count all the files that end with `.go`.
Then, you give it a directory to search, such as `path/to/directory`.
The program will then search the directory, and count every line in every file that has the correct file type.

You can also give line counter multiple fileTypes to search, such as `js`, `html`, and `css`.
Here are some example queries:
```
> lc go C:/Users/Documents/Code/Go/MyProject
Total: 500 with 0 errors.
```
```
> lc js html css ../../JavaScript/MyWebsite
Total: 78 with 1 errors.
```
