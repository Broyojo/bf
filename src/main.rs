use std::{fs, io::Result};

fn main() -> Result<()> {
    let source = fs::read_to_string("hanoi.b")?;

    println!("{:?}", source);

    let mut program = "use std::io::{self, Read, Write};

fn main() {
    let mut mem = [0u8; 30000];
    let mut dp = 0;"
        .to_string();

    for symbol in source.chars() {
        program += match symbol {
            '>' => "dp += 1;",
            '<' => "dp -= 1;",
            '+' => "mem[dp] += 1;",
            '-' => "mem[dp] -= 1;",
            '.' => "print!(\"{}\", mem[dp] as char);\nio::stdout().flush().unwrap();",
            ',' => "/* TODO */",
            '[' => "while mem[dp] != 0 {",
            ']' => "}",
            _ => continue,
        };
        program.push('\n');
    }

    program.push('}');

    fs::write("program.rs", program)?;

    Ok(())
}
