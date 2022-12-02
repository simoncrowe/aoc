use std::fs;

fn main() {
    let cals_flat = fs::read_to_string("01_input.txt").unwrap();
    let mut totals: Vec<u32> = cals_flat.split("\n\n").map(sum_cals).collect();
    totals.sort_by(|a, b| b.cmp(a));
    let top_three = &totals[..3];
    println!("Sum of top three is {}", top_three.iter().sum::<u32>());
}

fn sum_cals(cals_flat: &str) -> u32 {
    cals_flat
        .trim()
        .split("\n")
        .map(|cals| cals.parse::<u32>().unwrap())
        .sum()
}
