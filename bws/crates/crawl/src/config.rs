use std::{
    collections::HashMap,
    fs,
    path::{Path, PathBuf},
};

pub enum SourceEnum {
    WIKIDICH,
    METRUYENCV,
}

impl SourceEnum {
    pub fn as_str(&self) -> &'static str {
        match self {
            SourceEnum::WIKIDICH => "wikidich",
            SourceEnum::METRUYENCV => "metruyencv",
        }
    }

    pub fn to_string(&self) -> String {
        match self {
            SourceEnum::WIKIDICH => "wikidich".to_string(),
            SourceEnum::METRUYENCV => "metruyencv".to_string(),
        }
    }
}

pub const SOURCE_URL: &str = "URL";
pub const SOURCE_URL_SEARCH: &str = "URL_SEARCH";
pub const CATEGORY_WRAPPER: &str = "CATEGORY_WRAPPER_PATH";
pub const CATEGORY_TYPE: &str = "CATEGORY_TYPE_PATH";
pub const CATEGORY_GROUP: &str = "CATEGORY_GROUP_PATH";
pub const CATEGORY_TITLE: &str = "CATEGORY_TITLE_PATH";
pub const CATEGORY_ID: &str = "CATEGORY_ID_PATH";

const CONFIG_PATH: &str = "config.toml";

#[derive(Debug)]
pub struct Config {
    item: HashMap<String, HashMap<String, String>>,
}

impl Config {
    pub fn new() -> Self {
        Config {
            item: HashMap::new(),
        }
    }

    pub fn initialize(&mut self) {
        // get filepath of config
        let current_dir = std::env::current_dir().unwrap();
        let filepath = current_dir.join(Path::new(CONFIG_PATH));
        log::info!("Load config from: {}", filepath.display());

        self.load_toml(filepath);
    }

    pub fn get_source(&self, source: SourceEnum) -> Option<&HashMap<String, String>> {
        let source = source.as_str();
        self.item.get(source)
    }

    pub fn sources(&self) -> Vec<String> {
        let keys: Vec<&String> = self.item.keys().collect();
        let mut source: Vec<String> = Vec::new();

        for key in keys {
            source.push(key.to_string());
        }
        source
    }

    fn load_toml(&mut self, filepath: PathBuf) {
        let toml_content = fs::read_to_string(filepath)
            .expect(format!("File {} is not found.", CONFIG_PATH).as_str());
        let parsed_toml = toml::from_str(&toml_content).expect("Failed to parse toml values");

        // pass toml values to HashMap
        if let toml::Value::Table(table) = parsed_toml {
            for (key, value) in table {
                match value {
                    toml::Value::Table(nested_table) => {
                        if !self.item.contains_key(&key.to_string()) {
                            self.item.insert(key.to_string(), HashMap::new());
                        }
                        for (nkey, nvalue) in nested_table {
                            if let Some(n_map) = self.item.get_mut(&key) {
                                n_map.insert(nkey, nvalue.to_string().replace("\"", ""));
                            }
                        }
                    }
                    other => {
                        panic!("The non-string, non-table value of toml file, {:?}", other);
                    }
                }
            }
        }
    }
}
