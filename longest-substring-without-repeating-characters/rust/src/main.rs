use std::collections::HashMap;

struct Solution {

}

impl Solution {
    pub fn length_of_longest_substring(&self, s: String) -> i32 {
        let mut map: HashMap<char, u8> = HashMap::new();

        let mut i = 0;
        let mut j = 0;
        let mut max = 0;
        
        for ch in s.chars() {
            map.entry(ch).or_default();
            // map.entry(ch).or_insert(0);

            if let Some(x) = map.get_mut(&ch) {
                *x += 1;
            }

            println!("{:?}", map);
            while map.get(&ch).unwrap() > &1  {
                let i_c = s.chars().nth(i).unwrap();
                if let Some(x) = map.get_mut(&i_c) {
                    *x -= 1;
                }
                i += 1;
            }
            println!("{:?}", map);

            if j - i + 1 > max {
                max = j - i + 1;
                let sub = &s[i..=j];
                println!("sub: {:?}, max: {}", sub, max);
            }

            j += 1;
        }


        println!("{:?}", map);
        println!("{}", max);

        // return s.len() as i32;

        return max as i32;
    }
}

fn main() {
    let s = Solution{};
    let len = s.length_of_longest_substring(String::from("hello"));
    println!("{}", len);
    let len = s.length_of_longest_substring(String::from("pwwkew"));
    println!("{}", len);
    // println!("Hello, world!");
}
