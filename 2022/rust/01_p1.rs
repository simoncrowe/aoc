use std::fs;

fn main() {
    let cals_flat = fs::read_to_string("01_input.txt").unwrap();
    let max_total_cals = cals_flat.split("\n\n").map(sum_cals).max().unwrap();
    println!("Max is {}", max_total_cals);
}

fn sum_cals(cals_flat: &str) -> u32 {
    cals_flat
        .trim()
        .split("\n")
        .map(|cals| cals.parse::<u32>().unwrap())
        .sum()
}
