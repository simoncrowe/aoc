use std::fs;

fn main() {
    let games_text = fs::read_to_string("02_input.txt").unwrap();
    let total_score: u32 = games_text.split("\n").map(score).sum();
    println!("The total score would be {total_score}");
}

fn score(game_text: &str) -> u32 {
    match game_text.split_once(" ") {
        Some(("A", "X")) => 1 + 3,
        Some(("A", "Y")) => 2 + 6,
        Some(("A", "Z")) => 3 + 0,
        Some(("B", "X")) => 1 + 0,
        Some(("B", "Y")) => 2 + 3,
        Some(("B", "Z")) => 3 + 6,
        Some(("C", "X")) => 1 + 6,
        Some(("C", "Y")) => 2 + 0,
        Some(("C", "Z")) => 3 + 3,
        Some((_, _)) => 0,
        None => 0,
    }
}
