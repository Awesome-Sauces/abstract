//use std::env;
use std::fs;


pub fn load() {
    let bavel_avm = "avm_libs/bavel.avm".to_owned();
    let console_avm: String = "avm_libs/console.avm".to_owned();

    execute("x86-BAVEL".to_owned(), bavel_avm);
    execute("x86-CONSLE".to_owned(), console_avm);
}

fn execute(instance_name: String, filename: String) {
    let opcodes = digest(filename);

    for code in opcodes {
        println!("{instance_name}: <{code}>");
    }
}

fn digest(filename: String) -> Vec<String> {

    let contents = fs::read_to_string(filename)
        .expect("Should have been able to read the file");

    let mut word: String = "".to_owned();

    let mut str: bool = false;


    let mut opcodes: Vec<String> = Vec::new();

    for character in contents.chars() {
        if (character.is_whitespace() && !str && !character.eq(&'\n')) || character.eq(&';'){

            opcodes.push(word);
            word = "".to_owned();

            if character.eq(&';') {
                opcodes.push("END".to_owned());
            }

        }else {
            if character.is_ascii_punctuation() {
                if character.eq(&'"') {
                    str = !str;
                    continue;
                }
            }

            if !character.eq(&'\n') {
                word = format!("{word}{character}");
            }
            
        }
    }

    return opcodes;

    
}