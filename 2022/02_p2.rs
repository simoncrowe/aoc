use std::fs;

fn main() {
    let games_text = fs::read_to_string("02_input.txt").unwrap();
    let total_score: u32 = games_text.split("\n").map(play).sum();
    println!("The total score would be {total_score}");
}

fn play(game_text: &str) -> u32 {
    match game_text.split_once(" ") {
        Some(("A", "X")) => 3 + 0, // Scissors lose to rock
        Some(("A", "Y")) => 1 + 3, // Rock draws with rock
        Some(("A", "Z")) => 2 + 6, // Paper beats rock
        Some(("B", "X")) => 1 + 0, // Rock loses to paper
        Some(("B", "Y")) => 2 + 3, // Paper draws with paper
        Some(("B", "Z")) => 3 + 6, // Scissors beat paper
        Some(("C", "X")) => 2 + 0, // Paper loses to scissors
        Some(("C", "Y")) => 3 + 3, // Scissors draw with scissors
        Some(("C", "Z")) => 1 + 6, // Rock beats scissors
        Some((_, _)) => 0,
        None => 0,
    }
}
