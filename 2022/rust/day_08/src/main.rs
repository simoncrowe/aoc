use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

const GRID_SIZE: usize = 99;

pub fn main() -> io::Result<()> {
    let input = File::open("/home/sc/git/aoc/2022/input/08_input.txt")?;
    let mut grid = [[0u8; GRID_SIZE]; GRID_SIZE];
    for (y_idx, row) in BufReader::new(input).lines().enumerate() {
        for (x_idx, height) in row?.chars().enumerate() {
            grid[x_idx][y_idx] = height.to_string().parse::<u8>().unwrap();
        }
    }
    let coords = get_visible_coords(&grid);
    println!("Answer to part one {}", coords.len());
    let scores = get_scenic_scores(&grid);
    println!("Answer to part two {}", scores.into_iter().max().unwrap());
    Ok(())
}

fn get_visible_coords(grid: &[[u8; GRID_SIZE]]) -> HashSet<(usize, usize)> {
    let mut coords: HashSet<(usize, usize)> = HashSet::new();
    for x in 0..GRID_SIZE {
        coords.insert((x, 0));
        let mut max: u8 = grid[x][0];
        for y in 1..GRID_SIZE {
            if grid[x][y] > max {
                coords.insert((x, y));
                max = grid[x][y];
            }
        }
    }
    for x in 0..GRID_SIZE {
        coords.insert((x, GRID_SIZE - 1));
        let mut max: u8 = grid[x][GRID_SIZE - 1];
        for y in (0..GRID_SIZE - 1).rev() {
            if grid[x][y] > max {
                coords.insert((x, y));
                max = grid[x][y];
            }
        }
    }
    for y in 0..GRID_SIZE {
        coords.insert((0, y));
        let mut max: u8 = grid[0][y];
        for x in 1..GRID_SIZE {
            if grid[x][y] > max {
                coords.insert((x, y));
                max = grid[x][y];
            }
        }
    }
    for y in 0..GRID_SIZE {
        coords.insert((GRID_SIZE - 1, y));
        let mut max: u8 = grid[GRID_SIZE - 1][y];
        for x in (0..GRID_SIZE - 1).rev() {
            if grid[x][y] > max {
                coords.insert((x, y));
                max = grid[x][y];
            }
        }
    }
    coords
}

fn get_scenic_scores(grid: &[[u8; GRID_SIZE]]) -> Vec<u32> {
    let mut scores = Vec::new();
    for x in 0..GRID_SIZE {
        for y in 0..GRID_SIZE {
            let current = grid[x][y];
            let mut up = 0;
            for look_y in (y + 1)..GRID_SIZE {
                up += 1;
                if grid[x][look_y] >= current {
                    break;
                }
            }
            let mut right = 0;
            for look_x in (x + 1)..GRID_SIZE {
                right += 1;
                if grid[look_x][y] >= current {
                    break;
                }
            }
            let mut down = 0;
            for look_y in (0..y).rev() {
                down += 1;
                if grid[x][look_y] >= current {
                    break;
                }
            }
            let mut left = 0;
            for look_x in (0..x).rev() {
                left += 1;
                if grid[look_x][y] >= current {
                    break;
                }
            }
            scores.push(up * right * down * left);
        }
    }
    scores
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one_example() {
        let grid = [
            [3, 2, 6, 3, 3],
            [0, 5, 5, 3, 5],
            [3, 5, 3, 5, 3],
            [7, 1, 3, 4, 9],
            [3, 2, 2, 9, 0],
        ];
        let coords = get_visible_coords(&grid);
        assert_eq!(coords.len(), 21);
    }
    #[test]
    fn test_part_two_example() {
        let grid = [
            [3, 2, 6, 3, 3],
            [0, 5, 5, 3, 5],
            [3, 5, 3, 5, 3],
            [7, 1, 3, 4, 9],
            [3, 2, 2, 9, 0],
        ];
        let scores = get_scenic_scores(&grid);
        assert_eq!(scores.into_iter().max().unwrap(), 8);
    }
}
