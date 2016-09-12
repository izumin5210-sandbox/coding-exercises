puts gets.split(" ").map(&:to_i).tap { |(a, b)| break a >= b } ? "OK" : "NG"
