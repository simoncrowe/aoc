use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::string::String;

fn main() -> io::Result<()> {
    let input = File::open("04_input.txt")?;
    let lines = BufReader::new(input).lines();
    let contained_count = lines
        .map(|line| line.unwrap())
        .filter(contains)
        .collect::<Vec<String>>()
        .len();
    println!("Count of assignemnts where one contains the other {contained_count}");
    Ok(())
}

fn contains(assignment: &String) -> bool {
    let (first, second) = assignment.split_once(",").unwrap();
    let (first_start, first_end) = first.split_once("-").unwrap();
    let (second_start, second_end) = second.split_once("-").unwrap();

    let a_start = first_start.parse::<u32>().unwrap();
    let a_end = first_end.parse::<u32>().unwrap();
    let b_start = second_start.parse::<u32>().unwrap();
    let b_end = second_end.parse::<u32>().unwrap();

    let a_contains_b = a_start <= b_start && a_end >= b_end;
    let b_contains_a = b_start <= a_start && b_end >= a_end;
    a_contains_b || b_contains_a
}
