use unidecode::unidecode;

pub fn get_slug(str: &String) -> String {
    let ascii_str = unidecode(str.as_str());
    ascii_str.replace(" ", "-")
}
