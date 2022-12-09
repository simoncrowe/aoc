use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

const GRID_SIZE: usize = 99;

pub fn answer() -> io::Result<()> {
    let input = File::open("/home/sc/git/aoc/2022/input/08_input.txt")?;
    let mut grid = [[0u8; GRID_SIZE]; GRID_SIZE];
    for (y_idx, row) in BufReader::new(input).lines().enumerate() {
        for (x_idx, height) in row?.chars().enumerate() {
            grid[x_idx][y_idx] = height.to_string().parse::<u8>().unwrap();
        }
    }
    let coords = get_visible_coords(&grid);
    println!("Answer to part one {}", coords.len());
    Ok(())
}

fn get_visible_coords(grid: &[[u8; GRID_SIZE]]) -> HashSet<(usize, usize)> {
    let mut coords: HashSet<(usize, usize)> = HashSet::new();
    println!("From top");
    for x in 0..GRID_SIZE {
        coords.insert((x, 0));
        let mut max: u8 = grid[x][0];
        println!("Add {}, {}: {}", x, 0, grid[x][0]);
        for y in 1..GRID_SIZE {
            if grid[x][y] > max {
                coords.insert((x, y));
                println!("Add {}, {}: {}", x, y, grid[x][y]);
                max = grid[x][y];
            }
        }
        println!("Next row");
    }
    println!("From bottom");
    for x in 0..GRID_SIZE {
        coords.insert((x, GRID_SIZE - 1));
        let mut max: u8 = grid[x][GRID_SIZE - 1];
        println!("Add {}, {}: {}", x, GRID_SIZE - 1, grid[x][GRID_SIZE - 1]);
        for y in (0..GRID_SIZE - 1).rev() {
            if grid[x][y] > max {
                coords.insert((x, y));
                println!("Add {}, {}: {}", x, y, grid[x][y]);
                max = grid[x][y];
            }
        }
        println!("Next row");
    }
    println!("From left");
    for y in 0..GRID_SIZE {
        coords.insert((0, y));
        println!("Add {}, {}: {}", 0, y, grid[0][y]);
        let mut max: u8 = grid[0][y];
        for x in 1..GRID_SIZE {
            if grid[x][y] > max {
                coords.insert((x, y));
                println!("Add {}, {}: {}", x, y, grid[x][y]);
                max = grid[x][y];
            }
        }
        println!("Next column");
    }
    println!("From right");
    for y in 0..GRID_SIZE {
        coords.insert((GRID_SIZE - 1, y));
        println!("Add {}, {}: {}", GRID_SIZE - 1, y, grid[GRID_SIZE - 1][y]);
        let mut max: u8 = grid[GRID_SIZE - 1][y];
        for x in (0..GRID_SIZE - 1).rev() {
            if grid[x][y] > max {
                coords.insert((x, y));
                println!("Add {}, {}: {}", x, y, grid[x][y]);
                max = grid[x][y];
            }
        }
        println!("Next column");
    }
    coords
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
            [3, 2, 2, 9, 0]
        ];
        let coords = get_visible_coords(&grid);
        assert_eq!(coords.len(), 21);
    }
}
