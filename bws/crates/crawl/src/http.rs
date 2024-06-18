use reqwest::Client;

#[derive(Debug)]
pub struct Http {
    client: Client,
}

impl Http {
    fn new() -> Self {
        let client = Client::new();
    }
}
