use std::fs;

fn main() {
    let rucksacks_contents = fs::read_to_string("03_input.txt").unwrap();
    let total_priorities: u32 = rucksacks_contents
        .trim()
        .split("\n")
        .collect::<Vec<&str>>()
        .chunks(3)
        .map(badge_priority)
        .sum();
    println!("Sum of priorities of badges is {total_priorities}");
}

fn badge_priority(rucksacks: &[&str]) -> u32 {
    let one = rucksacks[0].as_bytes();
    let two = rucksacks[1].as_bytes();
    let three = rucksacks[2].as_bytes();
    let badge_item = one
        .iter()
        .filter(|item| two.contains(item) && three.contains(item))
        .next()
        .unwrap();
    priority(badge_item)
}

fn priority(item: &u8) -> u32 {
    if *item <= 90 {
        // Z is the byte 91 in unicode
        (item - 38).into()
    } else {
        (item - 96).into()
    }
}
