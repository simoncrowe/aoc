use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::string::String;

fn main() -> io::Result<()> {
    let input = File::open("03_input.txt")?;
    let lines = BufReader::new(input).lines();
    let total_priorities: u32 = lines 
        .map(|line| line.unwrap())
        .map(common_item_priority)
        .sum();
    println!("Sum of priorities of common items is {total_priorities}");
    Ok(())
}

fn common_item_priority(rucksack_contents: String) -> u32 {
    let (comp_one, comp_two) = rucksack_contents
        .as_bytes()
        .split_at(rucksack_contents.len() / 2);
    let common_item = comp_one
        .into_iter()
        .filter(|item| comp_two.contains(item))
        .next()
        .unwrap();
    priority(common_item)
}

fn priority(item: &u8) -> u32 {
    if *item <= 90 {
        // Z is has the value 91 in unicode
        (item - 38).into()
    } else {
        (item - 96).into()
    }
}
