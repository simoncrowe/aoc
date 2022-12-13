use std::fs::File;
use std::io::{self, prelude::*, BufReader};

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
    let mut paths: Vec<Vec<(Direction, (usize, usize))>> = Vec::new();
    paths.push(vec![(Direction::Unknown, grid.start)]);
    loop {
        let mut fully_traversed = true;
        let mut new_paths: Vec<Vec<(Direction, (usize, usize))>> = Vec::new();
        for path in paths.iter_mut() {
            let (from_dir, cur_pos) = path.last().unwrap();
            let directions = get_directions(&grid, *cur_pos, *from_dir);
            let mut dir_iter = directions.into_iter();
            match dir_iter.next() {
                Some((dir, pos)) => {
                    path.push((dir, pos ));
                    fully_traversed = false;
                },
                None => {continue}
            }
            while let Some((dir, pos)) = dir_iter.next() {
                new_paths.push(vec![(dir, pos )]); 
            }
        }
        paths.append(&mut new_paths);
        if fully_traversed {
            break;
        }
    }
    Ok(paths.iter().map(|path| path.len()).max().unwrap())
}

fn get_directions(grid: &Terrain, pos: (usize, usize), from: Direction) -> Vec<(Direction, (usize, usize) )> {
    let cur_val = grid.get_elevation(pos.0, pos.1).unwrap();
    let mut dirs: Vec<(Direction, (usize, usize))> = Vec::new();
    println!("Checking directions for {}, {}", pos.0, pos.1);
    if !matches!(from,  Direction::Up) && pos.1 != grid.height {
        match grid.get_elevation(pos.0, pos.1 + 1)  {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push((Direction::Up, (pos.0, pos.1 + 1)));
                }
            },
            None => {}
        }
    }
    if !matches!(from, Direction::Right) && pos.0 != grid.width {
        match grid.get_elevation(pos.0 + 1, pos.1) {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push((Direction::Right, (pos.0 + 1, pos.1)));
                }
            },
            None => {}
        }
    }
    if !matches!(from, Direction::Down) && pos.1 != 0 {
        match grid.get_elevation(pos.0, pos.1 - 1) {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push((Direction::Down, (pos.0, pos.1 - 1)));
                }
            },
            None => {}
        }
    }
    if !matches!(from, Direction::Left) && pos.0 != 0 {
        match grid.get_elevation(pos.0 - 1, pos.1) {
            Some(elevation) => { 
                if cur_val + 1 >= elevation {
                    dirs.push((Direction::Left, (pos.0 - 1, pos.1)));
                }
            },
            None => {}
        }
    }
    dirs
}

#[derive(Copy, Clone)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
    Unknown,
}

struct Terrain {
    elevations: Vec<u8>,
    width: usize,
    height: usize,
    start: (usize, usize),
    end: (usize, usize),
}

impl Terrain {
    pub fn new(input_path: &str) -> io::Result<Terrain> {
        let input = File::open(input_path)?;
        let lines = BufReader::new(input).lines();
        let mut elevations: Vec<u8> = Vec::new();
        let mut width: usize = 0;
        let mut height: usize = 0;
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
                if ord == UPPERCASE_S_ORD {
                    start.0 = col_idx;
                    start.1 = row_idx;
                } else if ord == UPPERCASE_E_ORD {
                    end.0 = col_idx;
                    end.1 = row_idx;
                }
                elevations.push(height_from_ord(ord));
            }
            width = line.len();
            height += 1;
        }
        let terrain = Terrain {
            elevations,
            width,
            height,
            start,
            end,
        };
        Ok(terrain)
    }
    pub fn get_elevation(&self, x: usize, y: usize) -> Option<u8> {
        
        match self.elevations.get((y * self.width) + x) {
            Some(elevation) => Some(*elevation),
            None => None
        }
    }
}

fn height_from_ord(ord: u8) -> u8 {
    println!("{}", ord);
    if ord == UPPERCASE_S_ORD {
        0
    } else if ord == UPPERCASE_E_ORD {
        25
    } else {
        ord - LOWECASE_A_ORD
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
        assert_eq!(25, grid.get_elevation(5, 2).unwrap());
        assert_eq!((0, 4), grid.start);
        assert_eq!(0, grid.get_elevation(0, 4).unwrap());
    }

    #[test]
    fn test_part_one_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        assert_eq!(31, compute_shortest_path_length(&test_input_path).unwrap());
    }
}
