use crate::config::*;

use super::{Category, Source};
use async_trait::async_trait;
use scraper::{selectable::Selectable, Html, Selector};
use std::collections::HashMap;

#[derive(Debug)]
pub struct Wikidich {
    starturl: String,
    categories: HashMap<String, Vec<Category>>,
}

impl Wikidich {
    /// Init wikidich instance
    pub fn new() -> Self {
        Wikidich {
            starturl: String::from(""),
            categories: HashMap::new(),
        }
    }

    /// Extract wikidich metadata - categories
    async fn extract_categories(
        &mut self,
        config: &HashMap<String, String>,
    ) -> anyhow::Result<(), Box<dyn std::error::Error>> {
        let url: String = config.get(SOURCE_URL_SEARCH).unwrap().to_string();
        let resp = reqwest::get(url).await?;

        let document = Html::parse_document(resp.text().await?.as_str());
        let selector = Selector::parse(config.get(CATEGORY_WRAPPER).unwrap()).unwrap();

        for element in document.select(&selector) {
            // extract category type
            let category_type_sel = Selector::parse(config.get(CATEGORY_TYPE).unwrap()).unwrap();
            let category_type_item = element.select(&category_type_sel);

            for ele in category_type_item {
                let category_type: &str = if ele.text().collect::<Vec<_>>().len() > 0 {
                    ele.text().collect::<Vec<_>>()[0]
                } else {
                    ""
                };

                if category_type != "" {
                    if !self.categories.contains_key(category_type) {
                        self.categories
                            .insert(category_type.to_string(), Vec::new());
                    }
                    // extract category group
                    let category_group_sel =
                        Selector::parse(config.get(CATEGORY_GROUP).unwrap()).unwrap();

                    for group_ele in element.select(&category_group_sel) {
                        // extract category - element id
                        let id_sel = Selector::parse(config.get(CATEGORY_ID).unwrap()).unwrap();
                        let id_item = group_ele.select(&id_sel).next().unwrap();
                        let id = id_item.value().attr("id").unwrap().to_string();

                        // extract category - element name
                        let name_sel =
                            Selector::parse(config.get(CATEGORY_TITLE).unwrap()).unwrap();
                        let name_item = group_ele.select(&name_sel).next().unwrap();
                        let name = name_item.inner_html();

                        let category = Category::new(name, id);
                        if let Some(category_t) =
                            self.categories.get_mut(&category_type.to_string())
                        {
                            category_t.push(category);
                        }
                    }
                }
            }
        }
        Ok(())
    }
}

#[async_trait]
impl Source for Wikidich {
    async fn crawl_metadata(&mut self, meta: &Config) -> Result<(), Box<dyn std::error::Error>> {
        let config = meta.get_source(SourceEnum::WIKIDICH).unwrap();

        // metadata - url
        self.starturl = config.get(SOURCE_URL_SEARCH).unwrap().to_string();

        // metadata - categories
        match self.extract_categories(config).await {
            Ok(_) => log::info!("Extract category successfull!"),
            Err(e) => return Err(e),
        }

        Ok(())
    }

    async fn crawl(&self) -> anyhow::Result<(), Box<dyn std::error::Error>> {
        Ok(())
    }
}
