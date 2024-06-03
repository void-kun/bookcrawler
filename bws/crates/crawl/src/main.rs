mod config;
mod source;
mod util;

use config::{Config, SourceEnum};
use source::{wikidich::Wikidich, Source};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    env_logger::init();
    log::info!("Start crawling...");

    // Load config
    let mut config = Config::new();
    config.initialize();

    let mut sources: Vec<Box<dyn Source>> = Vec::new();

    let source_names: Vec<String> = config.sources();
    if source_names.contains(&SourceEnum::WIKIDICH.to_string()) {
        let mut wikidich = Wikidich::new();
        match wikidich.crawl_booklist(&config, "https://truyenwikidich.net/tim-kiem?qs=1&status=5794f03dd7ced228f4419192&gender=5794f03dd7ced228f4419196&tc=&tf=0&m=6&y=2024&q=&start=0&vo=2".to_string()).await {

        // match wikidich.crawl_metadata(&config).await {
            Ok(_) => {
                sources.push(Box::new(wikidich));
            }
            Err(e) => {
                log::error!("Extract category error: {}", e);
            }
        }
    }
    if source_names.contains(&"metruyencv".to_string()) {}

    // crawling
    // for source in sources {
    //     let _ = source.crawl().await;
    // }
    Ok(())
}
