use std::fs;

fn main() {
    let rucksacks_contents = fs::read_to_string("03_input.txt").unwrap();
    let total_priorities: u32 = rucksacks_contents
        .trim()
        .split("\n")
        .map(common_item_priority)
        .sum();
    println!("Sum of priorities of common items is {total_priorities}");
}

fn common_item_priority(rucksack_contents: &str) -> u32 {
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
