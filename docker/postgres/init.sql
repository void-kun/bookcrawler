--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Databases
--

-- Dumped from database version 12.17
-- Dumped by pg_dump version 12.17

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

\connect bookcrawler

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: book_state; Type: TYPE; Schema: public; Owner: zrik
--

CREATE TYPE public.book_state AS ENUM (
    'đang tiến hành',
    'hoàn thành',
    'thái giám'
);


ALTER TYPE public.book_state OWNER TO zrik;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: tbl_author; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_author (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    name_slug character varying(255) NOT NULL,
    book_total integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.tbl_author OWNER TO zrik;

--
-- Name: tbl_author_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_author_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_author_id_seq OWNER TO zrik;

--
-- Name: tbl_author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_author_id_seq OWNED BY public.tbl_author.id;


--
-- Name: tbl_book; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_book (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    name_slug character varying(255) NOT NULL,
    summary text NOT NULL,
    state public.book_state NOT NULL,
    source_id integer NOT NULL,
    author_id integer NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    update_at timestamp with time zone DEFAULT now() NOT NULL,
    create_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tbl_book OWNER TO zrik;

--
-- Name: tbl_book_author; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_book_author (
    id integer NOT NULL,
    book_id integer NOT NULL,
    author_id integer NOT NULL,
    update_at timestamp with time zone DEFAULT now() NOT NULL,
    create_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tbl_book_author OWNER TO zrik;

--
-- Name: tbl_book_author_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_book_author_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_book_author_id_seq OWNER TO zrik;

--
-- Name: tbl_book_author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_book_author_id_seq OWNED BY public.tbl_book_author.id;


--
-- Name: tbl_book_category; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_book_category (
    id integer NOT NULL,
    book_id integer NOT NULL,
    category_id integer NOT NULL,
    update_at timestamp with time zone DEFAULT now() NOT NULL,
    create_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tbl_book_category OWNER TO zrik;

--
-- Name: tbl_book_category_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_book_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_book_category_id_seq OWNER TO zrik;

--
-- Name: tbl_book_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_book_category_id_seq OWNED BY public.tbl_book_category.id;


--
-- Name: tbl_book_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_book_id_seq OWNER TO zrik;

--
-- Name: tbl_book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_book_id_seq OWNED BY public.tbl_book.id;


--
-- Name: tbl_category; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_category (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    name_slug character varying(255) NOT NULL,
    book_total integer DEFAULT 0 NOT NULL,
    is_active boolean DEFAULT true NOT NULL
);


ALTER TABLE public.tbl_category OWNER TO zrik;

--
-- Name: tbl_category_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_category_id_seq OWNER TO zrik;

--
-- Name: tbl_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_category_id_seq OWNED BY public.tbl_category.id;


--
-- Name: tbl_chapter; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_chapter (
    id integer NOT NULL,
    book_id integer NOT NULL,
    title character varying(255) NOT NULL,
    title_slug character varying(255) NOT NULL,
    content text NOT NULL,
    update_at timestamp with time zone DEFAULT now() NOT NULL,
    create_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tbl_chapter OWNER TO zrik;

--
-- Name: tbl_chapter_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_chapter_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_chapter_id_seq OWNER TO zrik;

--
-- Name: tbl_chapter_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_chapter_id_seq OWNED BY public.tbl_chapter.id;


--
-- Name: tbl_source; Type: TABLE; Schema: public; Owner: zrik
--

CREATE TABLE public.tbl_source (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    name_slug character varying(255) NOT NULL,
    url character varying(255) NOT NULL,
    book_total integer DEFAULT 0 NOT NULL,
    is_active boolean DEFAULT true NOT NULL
);


ALTER TABLE public.tbl_source OWNER TO zrik;

--
-- Name: tbl_source_id_seq; Type: SEQUENCE; Schema: public; Owner: zrik
--

CREATE SEQUENCE public.tbl_source_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tbl_source_id_seq OWNER TO zrik;

--
-- Name: tbl_source_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: zrik
--

ALTER SEQUENCE public.tbl_source_id_seq OWNED BY public.tbl_source.id;


--
-- Name: tbl_author id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_author ALTER COLUMN id SET DEFAULT nextval('public.tbl_author_id_seq'::regclass);


--
-- Name: tbl_book id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book ALTER COLUMN id SET DEFAULT nextval('public.tbl_book_id_seq'::regclass);


--
-- Name: tbl_book_author id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_author ALTER COLUMN id SET DEFAULT nextval('public.tbl_book_author_id_seq'::regclass);


--
-- Name: tbl_book_category id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_category ALTER COLUMN id SET DEFAULT nextval('public.tbl_book_category_id_seq'::regclass);


--
-- Name: tbl_category id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_category ALTER COLUMN id SET DEFAULT nextval('public.tbl_category_id_seq'::regclass);


--
-- Name: tbl_chapter id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_chapter ALTER COLUMN id SET DEFAULT nextval('public.tbl_chapter_id_seq'::regclass);


--
-- Name: tbl_source id; Type: DEFAULT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_source ALTER COLUMN id SET DEFAULT nextval('public.tbl_source_id_seq'::regclass);


--
-- Data for Name: tbl_author; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_author (id, name, name_slug, book_total) FROM stdin;
\.


--
-- Data for Name: tbl_book; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_book (id, name, name_slug, summary, state, source_id, author_id, is_active, update_at, create_at) FROM stdin;
\.


--
-- Data for Name: tbl_book_author; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_book_author (id, book_id, author_id, update_at, create_at) FROM stdin;
\.


--
-- Data for Name: tbl_book_category; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_book_category (id, book_id, category_id, update_at, create_at) FROM stdin;
\.


--
-- Data for Name: tbl_category; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_category (id, name, name_slug, book_total, is_active) FROM stdin;
\.


--
-- Data for Name: tbl_chapter; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_chapter (id, book_id, title, title_slug, content, update_at, create_at) FROM stdin;
\.


--
-- Data for Name: tbl_source; Type: TABLE DATA; Schema: public; Owner: zrik
--

COPY public.tbl_source (id, name, name_slug, url, book_total, is_active) FROM stdin;
\.


--
-- Name: tbl_author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_author_id_seq', 1, false);


--
-- Name: tbl_book_author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_book_author_id_seq', 1, false);


--
-- Name: tbl_book_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_book_category_id_seq', 1, false);


--
-- Name: tbl_book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_book_id_seq', 1, false);


--
-- Name: tbl_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_category_id_seq', 1, false);


--
-- Name: tbl_chapter_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_chapter_id_seq', 1, false);


--
-- Name: tbl_source_id_seq; Type: SEQUENCE SET; Schema: public; Owner: zrik
--

SELECT pg_catalog.setval('public.tbl_source_id_seq', 1, false);


--
-- Name: tbl_author tbl_author_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_author
    ADD CONSTRAINT tbl_author_pkey PRIMARY KEY (id);


--
-- Name: tbl_book_author tbl_book_author_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_author
    ADD CONSTRAINT tbl_book_author_pkey PRIMARY KEY (id);


--
-- Name: tbl_book_category tbl_book_category_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_category
    ADD CONSTRAINT tbl_book_category_pkey PRIMARY KEY (id);


--
-- Name: tbl_book tbl_book_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book
    ADD CONSTRAINT tbl_book_pkey PRIMARY KEY (id);


--
-- Name: tbl_category tbl_category_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_category
    ADD CONSTRAINT tbl_category_pkey PRIMARY KEY (id);


--
-- Name: tbl_chapter tbl_chapter_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_chapter
    ADD CONSTRAINT tbl_chapter_pkey PRIMARY KEY (id);


--
-- Name: tbl_source tbl_source_pkey; Type: CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_source
    ADD CONSTRAINT tbl_source_pkey PRIMARY KEY (id);


--
-- Name: tbl_book_author tbl_book_author_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_author
    ADD CONSTRAINT tbl_book_author_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.tbl_author(id);


--
-- Name: tbl_book_author tbl_book_author_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_author
    ADD CONSTRAINT tbl_book_author_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.tbl_book(id);


--
-- Name: tbl_book tbl_book_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book
    ADD CONSTRAINT tbl_book_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.tbl_author(id);


--
-- Name: tbl_book_category tbl_book_category_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_category
    ADD CONSTRAINT tbl_book_category_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.tbl_book(id);


--
-- Name: tbl_book_category tbl_book_category_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book_category
    ADD CONSTRAINT tbl_book_category_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.tbl_category(id);


--
-- Name: tbl_book tbl_book_source_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_book
    ADD CONSTRAINT tbl_book_source_id_fkey FOREIGN KEY (source_id) REFERENCES public.tbl_source(id);


--
-- Name: tbl_chapter tbl_chapter_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: zrik
--

ALTER TABLE ONLY public.tbl_chapter
    ADD CONSTRAINT tbl_chapter_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.tbl_book(id);


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 12.17
-- Dumped by pg_dump version 12.17

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: zrik
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO zrik;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: zrik
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

