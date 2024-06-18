use crate::config::*;

use super::{Book, Category, Source};
use async_trait::async_trait;
use regex::Regex;
use scraper::{selectable::Selectable, Html, Selector};
use std::collections::HashMap;

#[derive(Debug, Default)]
pub struct Wikidich {
    pub starturl: String,
    pub books: Vec<Book>,
    pub categories: HashMap<String, Vec<Category>>,
}

impl Wikidich {
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

    async fn extract_booklist(
        &mut self,
        config: &HashMap<String, String>,
        url: &String,
    ) -> anyhow::Result<Vec<String>, Box<dyn std::error::Error>> {
        let resp = reqwest::get(url).await?;

        let document = Html::parse_document(resp.text().await?.as_str());
        let selector = Selector::parse("li.author-item").unwrap();
        let base_url = config.get(SOURCE_URL).unwrap();

        let mut result: Vec<String> = Vec::new();
        for element in document.select(&selector) {
            // extract title & url
            let book_name_sel = Selector::parse(".name-col > a").unwrap();
            let book_name_item = element.select(&book_name_sel).next().unwrap();
            let url = base_url.to_string() + book_name_item.attr("href").unwrap();

            result.push(url);
        }

        Ok(result)
    }

    async fn extract_bookinfo(
        &mut self,
        config: &HashMap<String, String>,
        url: &String,
    ) -> anyhow::Result<Wikidich, Box<dyn std::error::Error>> {
        let wikidich: Wikidich = Wikidich::default();

        let resp = reqwest::get(url).await?;
        let document = Html::parse_fragment(resp.text().await?.as_str());

        let title_sel = Selector::parse(config.get(BOOK_INFO_TITLE).unwrap()).unwrap();
        let title = document.select(&title_sel).next().unwrap().inner_html();

        log::info!("{}", title);
        let stats_sel = Selector::parse(config.get(BOOK_INFO_STATS).unwrap()).unwrap();
        let mut stats = document.select(&stats_sel);

        let visibility = stats.nth(0).unwrap().inner_html();
        let star = stats.nth(1).unwrap().inner_html();
        let comments = stats.nth(2).unwrap().inner_html();

        log::info!(
            "visibility {}, star {}, comments {}",
            visibility,
            star,
            comments
        );
        Ok(wikidich)
    }
}

#[async_trait]
impl Source for Wikidich {
    async fn crawl_metadata(
        &mut self,
        config: &HashMap<String, String>,
    ) -> Result<(), Box<dyn std::error::Error>> {
        // metadata - url
        self.starturl = config.get(SOURCE_URL_SEARCH).unwrap().to_string();

        // metadata - categories
        match self.extract_categories(config).await {
            Ok(_) => log::info!("Extract category successfull!"),
            Err(e) => return Err(e),
        }
        Ok(())
    }

    async fn crawl_booklist(
        &mut self,
        config: &HashMap<String, String>,
        url: &String,
    ) -> anyhow::Result<(), Box<dyn std::error::Error>> {
        let mut book_urls: Vec<String> = Vec::new();
        for page in (0..20).step_by(20) {
            let page_regex = Regex::new(r"start=(?P<start>\d+)&").unwrap();
            let current_page: u32 = page_regex.captures(&url).unwrap()["start"].parse().unwrap();

            let next_url = url.replace(
                format!("&start={}&", current_page).as_str(),
                format!("&start={}&", page).as_str(),
            );

            let mut result = match self.extract_booklist(config, &next_url).await {
                Ok(value) => value,
                Err(e) => return Err(e),
            };
            book_urls.append(&mut result);
        }

        let mut books: Vec<Wikidich> = Vec::new();
        for book_url in &book_urls {
            let result: Wikidich = self.extract_bookinfo(config, book_url).await.unwrap();
            books.push(result);
        }

        Ok(())
    }
}
