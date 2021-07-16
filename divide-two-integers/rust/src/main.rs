struct Solution {}
impl Solution {
    pub fn divide(&self, dividend: i32, divisor: i32) -> i32 {
        if dividend == i32::MIN && divisor == -1 {
            return i32::MAX;
        }

        let flag_diff = (dividend < 0) ^ (divisor < 0);

        let mut up = if dividend > 0 { -dividend } else { dividend };
        let down = if divisor > 0 { -divisor } else { divisor };

        let mut quotient: i32 = 0;

        while up <= down {
            let mut i = 0;
            let mut tmp = down;
            while up - tmp <= tmp {
                tmp <<= 1;
                i += 1;
            }
            up -= tmp;

            quotient += 1 << i;
        }

        quotient = if flag_diff { -quotient } else { quotient };

        return quotient;

        // println!("{}", flag_diff);
        // println!("***> {} / {}", dividend, divisor);
        // println!("===> {} / {} = {}", dividend, divisor, quotient);
    }
}
fn main() {
    let s = Solution {};
    println!("{}", s.divide(5, 2));
    println!("{}", s.divide(-5, 2));
    println!("{}", s.divide(-5, -2));
    println!("{}", s.divide(i32::MIN, -1));
    println!("{}", s.divide(i32::MIN, -2));

    // << double always
    // >> not double when negative

    // println!("{}>>{}={}", -3, 1, -3>>1);
    // println!("{}<<{}={}", -3, 1, -3<<1);

    // println!("{}>>{}={}", -5, 1, -5>>1);
    // println!("{}<<{}={}", -5, 1, -5<<1);
}
