use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

const LOWECASE_A_ORD: u8 = 97;
const UPPERCASE_S_ORD: u8 = 83;
const UPPERCASE_E_ORD: u8 = 69;

pub fn main() {
    let input_path = "/home/sc/git/aoc/2022/input/12_input.txt";
    println!(
        "Answer to part one: {}",
        compute_shortest_path_length(&input_path).unwrap()
    );
}

fn compute_shortest_path_length(input_path: &str) -> io::Result<usize> {
    let grid = Terrain::new(input_path).unwrap();
    let mut paths: Vec<Vec<Pos>> = Vec::new();
    let mut paths_to_dest: Vec<Vec<Pos>> = Vec::new();
    paths.push(vec![grid.start]);
    let mut traversed: HashSet<Pos> = HashSet::new();
    let mut depth = 0;
    loop {
        println!("\nChecking depth {}", depth);
        depth += 1;

        let mut fully_traversed = true;
        let mut new_paths: Vec<Vec<Pos>> = Vec::new();
        for path in paths.iter_mut().filter(|p| *p.last().unwrap() != grid.end) {
            let path_copy = path.clone();
            let cur_pos = path_copy.last().unwrap();
            let possible_steps = get_possible_steps(&grid, *cur_pos);
            let selected_steps = select_steps(cur_pos, possible_steps, &grid.end);
            //let traversed_copy: HashSet<Pos> = HashSet::from_iter(path.clone().into_iter());
            let traversed_copy = traversed.clone();
            let mut steps_iter = selected_steps
                .into_iter()
                .filter(|pos| !traversed_copy.contains(pos));
            match steps_iter.next() {
                Some(pos) => {
                    path.push(pos);
                    traversed.insert(pos);
                    if pos == grid.end {
                        paths_to_dest.push(path.clone());
                    } else {
                        fully_traversed = false;
                    }
                }
                None => continue,
            }
            while let Some(pos) = steps_iter.next() {
                //println!("Splitting to new line at {pos:?}");
                let mut new_path = path.clone();
                new_path.push(pos);
                traversed.insert(pos);
                if pos == grid.end {
                    paths_to_dest.push(new_path);
                } else {
                    new_paths.push(new_path);
                }
            }
        }
        paths.append(&mut new_paths);
        println!("total paths: {}", paths.len());
        println!("paths to dest: {}", paths_to_dest.len());
        let mut fewest_steps = 0;
        match paths_to_dest.iter().map(|path| path.len()).min() {
            Some(len) => {
                fewest_steps = len - 1;
                println!("Fewest steps {}", fewest_steps);
            }
            None => {}
        }
        if fully_traversed {
            break;
        }
    }
    Ok(paths_to_dest.iter().map(|path| path.len()).min().unwrap() - 1)
}

fn select_steps(current: &Pos, potential: Vec<Pos>, target: &Pos) -> Vec<Pos> {
    let perfect_direction: Vec<Pos> = potential.clone().into_iter().filter(|p| towards(current, p, target, 1)).collect();
    if perfect_direction.len() > 0 {
        return perfect_direction;
    } 
    potential.clone().into_iter().filter(|p| towards(current, p, target, 2)).collect()
}

fn towards(current: &Pos, potential: &Pos, target: &Pos, tolerance: usize) -> bool {
    let mut away_count = 0;
    if current.x < target.x && potential.x < current.x {
        away_count += 1; 
    }
    if current.y < target.y && potential.y < current.y {
        away_count += 1; 
    }
    if current.x > target.x && potential.x > current.x {
        away_count += 1; 
    }
    if current.y > target.y && potential.y > current.y {
        away_count += 1; 
    }
    away_count < tolerance 
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
    //println!("Found {} directions for {}, {}", dirs.len(), pos.x, pos.y);
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
                let bytes = vec![ord];
                let letter = std::str::from_utf8(&bytes).unwrap();
                //println!("{}, {}: {} ({})", row_idx, col_idx, ord, letter);
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
    fn test_part_one_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/12_test_input.txt";
        assert_eq!(31, compute_shortest_path_length(&test_input_path).unwrap());
    }
}
