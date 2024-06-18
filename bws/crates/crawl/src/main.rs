mod config;
mod http;
mod source;
mod util;

use config::{Config, SourceEnum, SOURCE_URL_SEARCH};
use source::{wikidich::Wikidich, Source};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    env_logger::init();
    log::info!("Start crawling...");

    let mut config = Config::new();
    config.initialize();

    let mut sources: Vec<Box<dyn Source>> = Vec::new();

    let source_names: Vec<String> = config.sources();
    if source_names.contains(&SourceEnum::WIKIDICH.to_string()) {
        let mut wikidich = Wikidich::default();
        let wikidich_config = config.get_source(SourceEnum::WIKIDICH).unwrap();

        let startup_url = wikidich_config.get(SOURCE_URL_SEARCH).unwrap();
        match wikidich.crawl_booklist(&wikidich_config, startup_url).await {
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
