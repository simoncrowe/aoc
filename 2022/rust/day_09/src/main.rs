use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::ops;

pub fn main() {
    let input_path = "/home/sc/git/aoc/2022/input/09_input.txt";
    println!(
        "Answer to part one: {}",
        answer_part_one(&input_path).unwrap()
    );
}

fn answer_part_one(input_path: &str) -> io::Result<usize> {
    let mut head: Point = Point::new(0, 0);
    let mut tail: Point = Point::new(0, 0);
    let mut trail: HashSet<Point> = HashSet::new();
    trail.insert(tail);
    let input = File::open(input_path)?;
    for line in BufReader::new(input).lines().map(|ln| ln.unwrap()) {
        println!("\n{line}\n");
        let (direction, mag) = line.split_once(" ").unwrap();
        let magnitude = mag.parse::<i32>().unwrap();
        let head_offset = Point::unit_vector_from_direction(&direction);

        for _ in 0..magnitude {
            head = head + head_offset;
            println!("Head pos: {head:?}; head offset: {head_offset:?}");
            let tail_offset = get_tail_offset(&head, &tail);
            tail = tail + tail_offset;
            trail.insert(tail);
            println!("Tail pos: {tail:?}; tail offset: {tail_offset:?}");
            println!("Tail len {}\n", trail.len());
        }
    }
    Ok(trail.len())
}

fn get_tail_offset(head: &Point, tail: &Point) -> Point {
    let diff = *head - *tail;
    println!("Diff of head and tail: {diff:?}");
    let mut x_offset = 0;
    if diff.x > 1 || (diff.x > 0 && diff.y > 1) {
        x_offset = 1;
    } else if diff.x < -1 || (diff.x < 0 && diff.y < -1) {
        x_offset = -1;
    }
    let mut y_offset = 0;
    if diff.y > 1 || (diff.x > 1 && diff.y > 0) {
        y_offset = 1;
    } else if diff.y < -1 || (diff.x < -1 && diff.y < 0) {
        y_offset = -1;
    }
    Point::new(x_offset, y_offset)
}

#[derive(Debug, PartialEq, Eq, Hash, Copy, Clone)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    pub fn new(x: i32, y: i32) -> Point {
        Point { x, y }
    }

    pub fn unit_vector_from_direction(direction: &str) -> Point {
        match direction {
            "U" => Point::new(0, 1),
            "R" => Point::new(1, 0),
            "D" => Point::new(0, -1),
            "L" => Point::new(-1, 0),
            _ => unreachable!("Direction should not deviate from U, R, D or L"),
        }
    }
}

impl ops::Add<Point> for Point {
    type Output = Point;

    fn add(self, _rhs: Point) -> Point {
        Point::new(self.x + _rhs.x, self.y + _rhs.y)
    }
}

impl ops::Sub<Point> for Point {
    type Output = Point;

    fn sub(self, _rhs: Point) -> Point {
        Point::new(self.x - _rhs.x, self.y - _rhs.y)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one_example() {
        let test_input_path = "/home/sc/git/aoc/2022/input/09_test_input.txt";
        assert_eq!(13, answer_part_one(&test_input_path).unwrap());
    }
}
