interface
=========

messing around with interfaces in the go language
I found most of the tutorials for them very confusing them- they didn't seem useful at all, just a lot of extra work for nothing.  After messing around with them in this program I've learned how to make them useful for me, basically by adding traditional object-oriented behavior back into go.  

In this program, Man and Woman are both sub-types of Person.  Person has a myGender method and a changeGender method.  Man and Woman pick up those automatically, which means they fit the Human interface.  so I can call the alterGender function on either of them.  If I made alterGender take a Person as its argument, this code wouldn't work, because unlike a "real" OOP language, Mand and Woman aren't "really" a Person, they just have its methods.
