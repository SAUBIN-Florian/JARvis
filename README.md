# 🤖☕ JARvis
Super simple Golang CLI tool for generating Java projects in 2 parts

1️⃣ It will generate a repository anywhere you want on your computer  
2️⃣ Now you will only use pure Java with the help of Makefile commands  


# Informations
✅ Tested with go version 1.23.0  
✅ Tested on linux Debian 12  
📦 Generate `.jar` files automatically  
⚡ Fast, portable, no dependencies  


This project is a non-standard super simple one, ultimately the only thing (for now) this CLI can do
is generating Java projects, but they are a bit different...


## 📌 Theses projects WILL NOT uses any traditionals tools like MAVEN or GRADLE !
It's simply a way to quickly create basic Java project for prototyping or experimenting stuff.
It use Makefile for building a JAR an execute it but more on that later...
It did not follow the classical Java files structure like src/com/example/too_much_stuff...


## ⚡ Quick start

🤖 JARvis part:  
```bash
# 1️⃣ Clone the repo and navigate to it
$ cd JARvis/

# 2️⃣ Build it (no go mod)
$ go build -o JARvis main.go

# 3️⃣ Run your shiny new tool!
$ ./JARvis
```

By default JARvis will create a repository named output at path ./  

```bash
# Change the name of your java project:
$ ./JARvis -name goffee

# Change the path:
$ ./JARvis -path /my_favorite_place_to_think_about_the_meaning_of_life

# Force it to overwrite an existing repo with the same name
$ ./JARvis -name goffee -force

# You can always use the help
$ ./JARvis -help
```

☕ Java part:  
```bash
# go to your freshly created project
$ cd output/

# it will create a .jar and execute it
$ make run
```

You can add repositories and .java files directly in the ./src repo 


# 🚀 Future
I want to add more build options to the Golang CLI

The main focus for the improvement is more on the Java side of this project
I want to add the possibility to download and setup externals packages and libraries from Maven Central


# 💣 Beware
This project is a goofy one, built for my personal use for prototyping. I strongly discourage you to use this for creating scalable and production ready apps...
For more informations check the LICENSE file