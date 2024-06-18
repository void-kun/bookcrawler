pub mod wikidich;

use async_trait::async_trait;

use crate::util;
use crate::Config;
use std::collections::HashMap;

#[async_trait]
pub trait Source {
    /// Extract metadata(categories, ...) for source.
    async fn crawl_metadata(
        &mut self,
        config: &HashMap<String, String>,
    ) -> Result<(), Box<dyn std::error::Error>>;
    /// Extract book info.
    async fn crawl_booklist(
        &mut self,
        config: &HashMap<String, String>,
        url: &String,
    ) -> anyhow::Result<(), Box<dyn std::error::Error>>;
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

#[derive(Debug)]
pub enum BookStatus {
    CONTINUES,
    PENDING,
    FINISHED,
    NOTYET,
}

impl BookStatus {
    pub fn as_str(&self) -> &'static str {
        match self {
            BookStatus::CONTINUES => "Còn tiếp",
            BookStatus::PENDING => "Tạm ngưng",
            BookStatus::FINISHED => "Hoàn thành",
            BookStatus::NOTYET => "Chưa xác minh",
        }
    }
}

pub fn book_status(status: &str) -> BookStatus {
    match status {
        "Còn tiếp" => BookStatus::CONTINUES,
        "Tạm ngưng" => BookStatus::PENDING,
        "Hoàn thành" => BookStatus::FINISHED,
        "Chưa xác minh" => BookStatus::NOTYET,
        _ => panic!("Error: book status not exists."),
    }
}

#[derive(Debug)]
pub struct Author {
    name: String,
    origin_name: String,
}

#[derive(Debug)]
pub struct Rate {
    visibility: u32,
    star: u32,
    comments: u32,
}

#[derive(Debug)]
pub struct Comment {
    owner: String,
    interval: u64,
    message: String,
    children: Vec<Comment>,
}

#[derive(Debug)]
pub struct Book {
    title: String,
    author: Author,
    url: String,
    categories: Vec<Category>,
    summary: String,
    avatar: String,
    rates: Rate,
    status: BookStatus,
    comments: Comment,
}
