use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let input = File::open("03_input.txt")?;
    let mut lines = BufReader::new(input).lines();
    let mut total: u32 = 0;
    while let Some(one) = lines.next() {
        let two = lines.next().unwrap();
        let three = lines.next().unwrap();
        total += badge_priority(&one?.as_bytes(), &two?.as_bytes(), &three?.as_bytes());
    }
    println!("Sum of priorities of badges is {total}");
    Ok(())
}

fn badge_priority(bag_one: &[u8], bag_two: &[u8], bag_three: &[u8]) -> u32 {
    let badge_item = bag_one
        .iter()
        .filter(|item| bag_two.contains(item) && bag_three.contains(item))
        .next()
        .unwrap();
    priority(badge_item)
}

fn priority(item: &u8) -> u32 {
    // Z is 90 in utf-8; a is 97
    if *item <= 90 {
        (item - 38).into()
    } else {
        (item - 96).into()
    }
}
