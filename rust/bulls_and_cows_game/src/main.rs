extern crate rand;

use std::io;
use rand::Rng;


#[derive(Debug, Eq, PartialEq)]
struct Number(char, char, char, char);

enum CheckResults {
    Guessed,
    NotGuessed { bulls: i32, cows: i32 },
}

impl Number {
    fn from_correct_string(input_string: String) -> Number {
        let chars: Vec<char> = input_string.chars().collect();
        Number(chars[0], chars[1], chars[2], chars[3])
    }

    fn has_digit(&self, ch: char) -> bool {
        return self.0 == ch
            || self.1 == ch
            || self.2 == ch
            || self.3 == ch;
    }
}

fn main() {
    println!("Bulls and cows game");
    println!("-------------------");

    let secret = generate_secret_number();
    let mut guesses_count = 0;

    loop {
        println!("Please input your 4 number guess:");

        let guess = read_input();
        if guess.len() != 4 {
            continue;
        }

        let input_number = Number::from_correct_string(guess);
        guesses_count = guesses_count + 1;

        match check_input(&secret, &input_number) {
            CheckResults::Guessed => {
                println!("You won!");
                println!("Guesses count: {}", guesses_count);
                break;
            }
            CheckResults::NotGuessed { bulls, cows } => {
                println!("bulls: {}, cows: {} \n", bulls, cows);
                continue;
            }
        }
    }
}

fn check_input(secret: &Number, input_number: &Number) -> CheckResults {
    if secret == input_number {
        return CheckResults::Guessed;
    }

    let mut bulls = 0;
    let mut cows = 0;

    //TODO Number struct to be re-written in Vec, to be able to generify that check
    if secret.0 == input_number.0 {
        bulls = bulls + 1;
    } else if secret.has_digit(input_number.0) {
        cows = cows + 1;
    }

    if secret.1 == input_number.1 {
        bulls = bulls + 1;
    } else if secret.has_digit(input_number.1) {
        cows = cows + 1;
    }

    if secret.2 == input_number.2 {
        bulls = bulls + 1;
    } else if secret.has_digit(input_number.2) {
        cows = cows + 1;
    }

    if secret.3 == input_number.3 {
        bulls = bulls + 1;
    } else if secret.has_digit(input_number.3) {
        cows = cows + 1;
    }

    return CheckResults::NotGuessed { bulls, cows };
}

fn read_input() -> String {
    let mut guess = String::new();
    io::stdin().read_line(&mut guess)
        .expect("Failed to read used input line");

    String::from(guess.trim())
}

fn generate_secret_number() -> Number {
    let digits = ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'];
    let mut used_digits: Vec<char> = Vec::new();

    let first_digit = generate_digit(&digits, &mut used_digits);
    let second_digit = generate_digit(&digits, &mut used_digits);
    let third_digit = generate_digit(&digits, &mut used_digits);
    let fourth_digit = generate_digit(&digits, &mut used_digits);

    Number(first_digit, second_digit, third_digit, fourth_digit)
}

fn generate_digit(numbers: &[char; 10], used_numbers: &mut Vec<char>) -> char {
    loop {
        let num = rand::thread_rng().choose(numbers).unwrap();
        if !used_numbers.contains(num) {
            used_numbers.push(*num);
            return *num;
        }
    };
}
