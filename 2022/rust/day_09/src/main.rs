use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::ops;

pub fn main() {
    let input_path = "/home/sc/git/aoc/2022/input/09_input.txt";
    println!(
        "Answer to part one: {}",
        compute_trail_length(2, &input_path).unwrap()
    );
    println!(
        "Answer to part two: {}",
        compute_trail_length(10, &input_path).unwrap()
    );
}

fn compute_trail_length(knots_count: usize, input_path: &str) -> io::Result<usize> {
    let mut rope: Vec<Vector2> = (0..knots_count).map(|_| Vector2::new(0, 0)).collect();
    let mut trail: HashSet<Vector2> = HashSet::new();
    trail.insert(rope[knots_count - 1]);

    let input = File::open(input_path)?;
    for line in BufReader::new(input).lines().map(|ln| ln.unwrap()) {
        let (direction, mag) = line.split_once(" ").unwrap();
        let magnitude = mag.parse::<i32>().unwrap();
        let head_offset = Vector2::unit_vector_from_direction(&direction);

        for _ in 0..magnitude {
            rope[0] = rope[0] + head_offset;
            for i in 1..knots_count {
                let tail_offset = get_tail_offset(&rope[i - 1], &rope[i]);
                rope[i] = rope[i] + tail_offset;
            }
            trail.insert(rope[knots_count - 1]);
        }
    }
    Ok(trail.len())
}

fn get_tail_offset(head: &Vector2, tail: &Vector2) -> Vector2 {
    let diff = *head - *tail;
    let mut x_offset = 0;
    if diff.x > 1 || (diff.x > 0 && (diff.y > 1 || diff.y < -1)) {
        x_offset = 1;
    } else if diff.x < -1 || (diff.x < 0 && (diff.y > 1 || diff.y < -1)) {
        x_offset = -1;
    }
    let mut y_offset = 0;
    if diff.y > 1 || (diff.y > 0 && (diff.x > 1 || diff.x < -1)) {
        y_offset = 1;
    } else if diff.y < -1 || (diff.y < 0 && (diff.x > 1 || diff.x < -1)) {
        y_offset = -1;
    }
    Vector2::new(x_offset, y_offset)
}

#[derive(Debug, PartialEq, Eq, Hash, Copy, Clone)]
struct Vector2 {
    x: i32,
    y: i32,
}

impl Vector2 {
    pub fn new(x: i32, y: i32) -> Vector2 {
        Vector2 { x, y }
    }

    pub fn unit_vector_from_direction(direction: &str) -> Vector2 {
        match direction {
            "U" => Vector2::new(0, 1),
            "R" => Vector2::new(1, 0),
            "D" => Vector2::new(0, -1),
            "L" => Vector2::new(-1, 0),
            _ => unreachable!("Direction should not deviate from U, R, D or L"),
        }
    }
}

impl ops::Add<Vector2> for Vector2 {
    type Output = Vector2;

    fn add(self, _rhs: Vector2) -> Vector2 {
        Vector2::new(self.x + _rhs.x, self.y + _rhs.y)
    }
}

impl ops::Sub<Vector2> for Vector2 {
    type Output = Vector2;

    fn sub(self, _rhs: Vector2) -> Vector2 {
        Vector2::new(self.x - _rhs.x, self.y - _rhs.y)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/09_test_input.txt";
        assert_eq!(13, compute_trail_length(2, &test_input_path).unwrap());
    }

    #[test]
    fn test_part_two_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/09_test_input.txt";
        assert_eq!(1, compute_trail_length(10, &test_input_path).unwrap());
    }
}
