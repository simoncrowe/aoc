use std::collections::{HashSet, VecDeque};
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

const LOWECASE_A_ORD: u8 = 97;
const UPPERCASE_S_ORD: u8 = 83;
const UPPERCASE_E_ORD: u8 = 69;

pub fn main() {
    let input_path = "/home/sc/git/aoc/2022/input/12_input.txt";
    println!(
        "Answer to part one: {}",
        compute_shortest_path_length_from_start_to_end(&input_path).unwrap()
    );
    println!(
        "Answer to part two: {}",
        compute_shortest_path_length_from_lowest_to_end(&input_path).unwrap()
    );
}

fn compute_shortest_path_length_from_start_to_end(input_path: &str) -> io::Result<u32> {
    let grid = Terrain::new(input_path).expect("Should be able to parse terrain");
    Ok(compute_shortest_path_length(&grid, grid.start, grid.end))
}

fn compute_shortest_path_length_from_lowest_to_end(input_path: &str) -> io::Result<u32> {
    let grid = Terrain::new(input_path).expect("Should be able to parse terrain");
    let shortest_length = grid
        .get_positions(0)
        .map(|start| compute_shortest_path_length(&grid, start, grid.end))
        .min()
        .expect("There should be a shortest path");
    Ok(shortest_length)
}

fn compute_shortest_path_length(grid: &Terrain, start: Pos, end: Pos) -> u32 {
    let mut traversed: HashSet<Pos> = HashSet::new();
    let mut traversal_queue: VecDeque<(Pos, u32)> = VecDeque::new();
    traversal_queue.push_back((start, 0));
    while let Some((cur_pos, cur_depth)) = traversal_queue.pop_front() {
        let possible_steps = get_possible_steps(&grid, cur_pos);
        let traversed_copy = traversed.clone();
        let mut steps_iter = possible_steps
            .into_iter()
            .filter(|pos| !traversed_copy.contains(pos));
        while let Some(pos) = steps_iter.next() {
            traversed.insert(pos);
            if pos == end {
                return cur_depth + 1;
            } else {
                traversal_queue.push_back((pos, cur_depth + 1));
            }
        }
    }
    u32::MAX 
}

fn get_possible_steps(grid: &Terrain, pos: Pos) -> Vec<Pos> {
    let cur_val = grid.get_elevation(pos.x, pos.y).unwrap();
    let mut dirs: Vec<Pos> = Vec::new();
    if pos.y <= grid.height {
        match grid.get_elevation(pos.x, pos.y + 1) {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push(Pos {
                        x: pos.x,
                        y: pos.y + 1,
                    });
                }
            }
            None => {}
        }
    }
    if pos.x <= grid.width {
        match grid.get_elevation(pos.x + 1, pos.y) {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push(Pos {
                        x: pos.x + 1,
                        y: pos.y,
                    });
                }
            }
            None => {}
        }
    }
    if pos.y != 0 {
        match grid.get_elevation(pos.x, pos.y - 1) {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push(Pos {
                        x: pos.x,
                        y: pos.y - 1,
                    });
                }
            }
            None => {}
        }
    }
    if pos.x != 0 {
        match grid.get_elevation(pos.x - 1, pos.y) {
            Some(elevation) => {
                if cur_val + 1 >= elevation {
                    dirs.push(Pos {
                        x: pos.x - 1,
                        y: pos.y,
                    });
                }
            }
            None => {}
        }
    }
    dirs
}
#[derive(Debug, PartialEq, Eq, Hash, Copy, Clone)]
struct Pos {
    x: usize,
    y: usize,
}

struct Terrain {
    elevations: Vec<u8>,
    width: usize,
    height: usize,
    start: Pos,
    end: Pos,
}

impl Terrain {
    pub fn new(input_path: &str) -> io::Result<Terrain> {
        let input = File::open(input_path)?;
        let lines = BufReader::new(input).lines();
        let mut elevations: Vec<u8> = Vec::new();
        let mut width: usize = 0;
        let mut height: usize = 0;
        let mut start = Pos { x: 0, y: 0 };
        let mut end = Pos { x: 0, y: 0 };
        for (row_idx, line) in lines
            .map(|ln| ln.unwrap())
            .collect::<Vec<String>>()
            .into_iter()
            .rev()
            .enumerate()
        {
            for (col_idx, ord) in line.as_bytes().iter().map(|o| *o).enumerate() {
                if ord == UPPERCASE_S_ORD {
                    start.x = col_idx;
                    start.y = row_idx;
                } else if ord == UPPERCASE_E_ORD {
                    end.x = col_idx;
                    end.y = row_idx;
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
            None => None,
        }
    }

    pub fn get_positions(&self, elevation: u8) -> impl Iterator<Item = Pos> + '_ {
        self.elevations
            .clone()
            .into_iter()
            .enumerate()
            .filter(move |(_idx, e)| e == &elevation)
            .map(|(idx, _e)| Pos {
                x: idx % self.width,
                y: idx / self.width,
            })
    }
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_terain_new() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        let grid = Terrain::new(test_input_path).unwrap();
        assert_eq!(Pos { x: 5, y: 2 }, grid.end);
        assert_eq!(25, grid.get_elevation(5, 2).unwrap());
        assert_eq!(Pos { x: 0, y: 4 }, grid.start);
        assert_eq!(0, grid.get_elevation(0, 4).unwrap());
    }

    #[test]
    fn test_terain_get_positions() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        let grid = Terrain::new(test_input_path).unwrap();
        let lowest: HashSet<Pos> = grid.get_positions(0).collect();
        let mut expected_lowest: HashSet<Pos> = HashSet::new();
        expected_lowest.insert(Pos { x: 0, y: 0 });
        expected_lowest.insert(Pos { x: 0, y: 1 });
        expected_lowest.insert(Pos { x: 0, y: 2 });
        expected_lowest.insert(Pos { x: 0, y: 3 });
        expected_lowest.insert(Pos { x: 0, y: 4 });
        expected_lowest.insert(Pos { x: 1, y: 4 });
        assert_eq!(lowest, expected_lowest);
    }

    #[test]
    fn test_part_one_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        assert_eq!(
            31,
            compute_shortest_path_length_from_start_to_end(&test_input_path).unwrap()
        );
    }

    #[test]
    fn test_part_two_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        assert_eq!(
            29,
            compute_shortest_path_length_from_lowest_to_end(&test_input_path).unwrap()
        );
    }
}
