pub mod wikidich;

use async_trait::async_trait;

use crate::util;
use crate::Config;

#[async_trait]
pub trait Source {
    /// Extract wikidich metadata
    async fn crawl_metadata(&mut self, meta: &Config) -> Result<(), Box<dyn std::error::Error>>;
    async fn crawl(&self) -> anyhow::Result<(), Box<dyn std::error::Error>>;
}

#[derive(Debug)]
pub struct Category {
    name: String,
    id: String,
    slug: String,
}

impl Category {
    pub fn new(name: String, id: String) -> Self {
        let slug = util::get_slug(&name).to_lowercase();

        Category { name, id, slug }
    }
}
