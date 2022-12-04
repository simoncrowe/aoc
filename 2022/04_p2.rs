use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::string::String;

fn main() -> io::Result<()> {
    let input = File::open("04_input.txt")?;
    let lines = BufReader::new(input).lines();
    let contained_count = lines
        .map(|line| line.unwrap())
        .filter(overlaps)
        .collect::<Vec<String>>()
        .len();
    println!("Count of overlapping assignments {contained_count}");
    Ok(())
}

fn overlaps(assignment: &String) -> bool {
    let (a_txt, b_txt) = assignment.split_once(",").unwrap();
    let (a_start, a_end) = a_txt.split_once("-").unwrap();
    let (b_start, b_end) = b_txt.split_once("-").unwrap();

    let mut a_rng = a_start.parse::<u32>().unwrap()..=a_end.parse::<u32>().unwrap();
    let b_rng = b_start.parse::<u32>().unwrap()..=b_end.parse::<u32>().unwrap();

    a_rng.any(|a| b_rng.contains(&a))
}
