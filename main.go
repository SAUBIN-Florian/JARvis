package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func WriteOutput(path string, name string) {
	fmt.Printf("A repository will be created at path '%s', with the name: %s\n", path, name)
	var base_path = path + "/" + name

	err := os.Mkdir(base_path, 0750)
	if err != nil {
		fmt.Println("ERROR: Can't create the repository project, something went wrong")
		return
	}

	os.MkdirAll(base_path+"/build", 0750)
	os.MkdirAll(base_path+"/src", 0750)

	var content = `package src;

class Main {
	public static void main(String[] args) {
		System.out.println("Generated with JARvis CLI");
	}
}
`

	err = os.WriteFile(base_path+"/src"+"/Main.java", []byte(content), 0660)
	if err != nil {
		os.Exit(1)
	}
}

func GenerateMakefile(template_file string, path string, name string) error {
	var content []byte
	var err error
	var default_makefile = `JAVAC=javac
JAVA=java
JAR=jar

SRC=$(wildcard src/*.java)
MAIN_PATH=src.Main
BUILD=build

jarify: clean
	@mkdir -p build
	@$(JAVAC) -d $(BUILD) $(SRC)
	@$(JAR) cfe $(BUILD)/App.jar $(MAIN_PATH) -C $(BUILD) .

run: jarify
	@$(JAVA) -jar $(BUILD)/App.jar

clean:
	@rm -rf $(BUILD)
`

	if template_file != "" {
		content, err = os.ReadFile(template_file)
		if err != nil {
			return fmt.Errorf("Can't read the template file provided: %v", err)
		}
	} else {
		content = []byte(default_makefile)
	}

	var base_path = path + "/" + name

	return os.WriteFile(base_path+"/Makefile", content, 0644)
}

func main() {
	var build_path = flag.String("path", "./", "Define the path were the java build will be wrote.")
	var repo_name = flag.String("name", "output", "Define the name of your java project")
	var makefile_template = flag.String("template", "", "Define the Makefile template to generate your jar project")
	var force = flag.Bool("force", false, "Overwrite your output java repository if already exist")
	flag.Parse()

	if *force {
		exec.Command("rm", "-rf", *build_path+"/"+*repo_name).Run()
	}

	WriteOutput(*build_path, *repo_name)
	GenerateMakefile(*makefile_template, *build_path, *repo_name)
}
