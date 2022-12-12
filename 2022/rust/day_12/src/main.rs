use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader, Lines};
use std::ops;

const LOWECASE_A_ORD: u8 = 97;
const UPPERCASE_S_ORD: u8 = 83;
const UPPERCASE_E_ORD: u8 = 69;

pub fn main() {
    let input_path = "/home/sc/git/aoc/2022/input/09_input.txt";
    println!(
        "Answer to part one: {}",
        compute_shortest_path_length(&input_path).unwrap()
    );
}

fn compute_shortest_path_length(input_path: &str) -> io::Result<usize> {
    let grid = Terrain::new(input_path).unwrap();
    Ok(42)
}

fn height_from_ord(ord: u8) -> u8 {
    if ord == UPPERCASE_S_ORD {
        0
    } else if ord == UPPERCASE_E_ORD {
        25
    } else {
        ord - LOWECASE_A_ORD
    }
}

struct Terrain {
    elevations: Vec<u8>,
    width: usize,
    start: (usize, usize),
    end: (usize, usize),
}

impl Terrain {
    pub fn new(input_path: &str) -> io::Result<Terrain> {
        let input = File::open(input_path)?;
        let lines = BufReader::new(input).lines();
        let mut elevations: Vec<u8> = Vec::new();
        let mut width: usize = 0;
        let mut start: (usize, usize) = (0, 0);
        let mut end: (usize, usize) = (0, 0);
        for (row_idx, line) in lines
            .map(|ln| ln.unwrap())
            .collect::<Vec<String>>()
            .into_iter()
            .rev()
            .enumerate()
        {
            for (col_idx, ord) in line.as_bytes().iter().map(|o| *o).enumerate() {
                println!("{}, {}: {}", col_idx, row_idx, ord);
                if ord == UPPERCASE_S_ORD {
                    start.0 = col_idx;
                    start.1 = row_idx;
                } else if ord == UPPERCASE_E_ORD {
                    end.0 = col_idx;
                    end.1 = row_idx;
                }
                elevations.push(height_from_ord(ord));
            }
            width = line.len()
        }
        let terrain = Terrain {
            elevations,
            width,
            start,
            end,
        };
        Ok(terrain)
    }
    pub fn get_elevation(&self, x: usize, y: usize) -> u8 {
        *self.elevations.get((y * self.width) + x).unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_terain_new() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        let grid = Terrain::new(test_input_path).unwrap();
        assert_eq!((5, 2), grid.end);
        assert_eq!(25, grid.get_elevation(5, 2));
        assert_eq!((0, 4), grid.start);
        assert_eq!(0, grid.get_elevation(0, 4));
    }

    #[test]
    fn test_part_one_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        assert_eq!(31, compute_shortest_path_length(&test_input_path).unwrap());
    }
}
