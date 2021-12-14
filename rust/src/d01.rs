pub fn parse(input: &str) -> Vec<u32> {
    input.lines().map(|line| line.parse().unwrap()).collect()
}

pub fn p1(depths: Vec<u32>) -> usize {
    depths
        .iter()
        .zip(depths.iter().skip(1))
        .filter(|(a, b)| b > a)
        .count()
}

pub fn p2(depths: Vec<u32>) -> usize {
    depths
        .iter()
        .zip(depths.iter().skip(3))
        .filter(|(a, d)| d > a)
        .count()
}

#[cfg(test)]
mod test {
    use super::*;
    #[test]
    fn test_parse() {
        let input = "123
456
9";
        assert_eq!(parse(input), vec![123, 456, 9]);
    }

    #[test]
    fn test_p1() {
        let depths = vec![1, 2, 1, 3];
        assert_eq!(p1(depths), 2);

        let s = std::fs::read_to_string("../inputs/1/input.txt").unwrap();
        assert_eq!(p1(parse(&s)), 1832);
    }

    #[test]
    fn test_p2() {
        let depths = vec![1, 2, 1, 3, 5, 2, 1];
        assert_eq!(p2(depths), 3);

        let s = std::fs::read_to_string("../inputs/1/input.txt").unwrap();
        assert_eq!(p2(parse(&s)), 1858);
    }
}
