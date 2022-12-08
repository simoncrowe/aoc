use std::fs::File;
use std::io::{self, prelude::*, BufReader};

const GRID_SIZE: usize = 99;

#[derive(Hash)]
struct Vector2 {
    pub x: i32,
    pub y: i32,
}

fn main() -> io::Result<()> {
    let input = File::open("08_input.txt")?;
    let mut grid = [[0u8; GRID_SIZE]; GRID_SIZE];
    for (y_idx, row) in BufReader::new(input).lines().enumerate() {
        for (x_idx, height) in row?.chars().enumerate() {
            grid[x_idx][y_idx] = height.to_string().parse::<u8>().unwrap();
        }
    }

    println!("Answer to part one {}", 42);
    Ok(())
}

fn get_visible_coords(dir: Vector2, grid: &[[u8; GRID_SIZE]]) -> Vec<Vector2> {
    let mut x: usize = if dir.x == -1 { GRID_SIZE - 1 } else { 0 };
    let mut y: usize = if dir.x == -1 { GRID_SIZE - 1 } else { 0 };
    let mut coords: Vec<Vector2> = Vec::new();
    loop {
        
        x += dir.x;
        y += dir.y;
        if x == GRID_SIZE || x < 0 || y == GRID_SIZE || y < 0 {
            break;
        }
    }
    coords
}
