/**
 * Cria o database
 */
CREATE DATABASE transparency
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Portuguese_Brazil.1252'
    LC_CTYPE = 'Portuguese_Brazil.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

/**
 * Cria a tabela
 */
CREATE TABLE public.scorecard
(
    unitid character varying(10) COLLATE pg_catalog."default" NOT NULL,
    opeid character varying(10) COLLATE pg_catalog."default",
    opeid6 character varying(10) COLLATE pg_catalog."default",
    instnm character varying(100) COLLATE pg_catalog."default",
    city character varying(100) COLLATE pg_catalog."default",
    stabbr character varying(2) COLLATE pg_catalog."default",
    insturl character varying(250) COLLATE pg_catalog."default",
    CONSTRAINT scorecard_pkey PRIMARY KEY (unitid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.scorecard
    OWNER to postgres;

/**
 * Copia os dados do arquivo preparado para 7 colunas sem o cabeçalho.
 * IMPORTANTE: substituir o diretorio informado pelo diretorio atual da sua máquina.
 */
COPY public.scorecard 
FROM 'D:\projetos\Apenas7colunas2SemCabecaco2.csv' 
WITH delimiter ';'