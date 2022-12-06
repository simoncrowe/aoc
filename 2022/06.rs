use std::collections::HashSet;
use std::fs;
use std::iter::FromIterator;

fn main() {
    let input = fs::read("./06_input.txt").unwrap();
    println!("Part one answer {}", get_end_of_unique_sequence(&input, 4));
    println!("Part two answer {}", get_end_of_unique_sequence(&input, 14));
}

fn get_end_of_unique_sequence(input: &Vec<u8>, seq_len: usize) -> usize {
    input[..]
        .windows(seq_len)
        .enumerate()
        .filter(|(_idx, vals)| HashSet::<u8>::from_iter(vals.iter().cloned()).len() == seq_len)
        .map(|(idx, _vals)| idx + seq_len)
        .next()
        .unwrap()
}
