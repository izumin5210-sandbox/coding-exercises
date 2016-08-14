puts gets.chomp.chars.inject("") { |s, c| s += c; (c == ?B) ? s[0..-3] : s }
