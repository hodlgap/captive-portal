--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Debian 16.0-1.pgdg120+1)
-- Dumped by pg_dump version 16.0

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
-- Name: update_auth_acknowledgment_log_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_auth_acknowledgment_log_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.auth_acknowledgment_log_updated_at = now();
    RETURN NEW;
END;
$$;


--
-- Name: update_auth_attempt_log_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_auth_attempt_log_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.auth_attempt_log_updated_at = now();
    RETURN NEW;
END;
$$;


--
-- Name: update_gateway_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.update_gateway_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.gateway_updated_at = now();
    RETURN NEW;
END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: auth_acknowledgment_log; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auth_acknowledgment_log (
    auth_acknowledgment_log_id bigint NOT NULL,
    auth_acknowledgment_log_gateway_name_hash character varying(64) NOT NULL,
    auth_acknowledgment_log_raw_payload text NOT NULL,
    auth_acknowledgment_log_created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    auth_acknowledgment_log_updated_at timestamp without time zone
);


--
-- Name: TABLE auth_acknowledgment_log; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.auth_acknowledgment_log IS 'client approval log';


--
-- Name: COLUMN auth_acknowledgment_log.auth_acknowledgment_log_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_acknowledgment_log.auth_acknowledgment_log_id IS 'primary key';


--
-- Name: COLUMN auth_acknowledgment_log.auth_acknowledgment_log_gateway_name_hash; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_acknowledgment_log.auth_acknowledgment_log_gateway_name_hash IS 'gateway name hash';


--
-- Name: COLUMN auth_acknowledgment_log.auth_acknowledgment_log_raw_payload; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_acknowledgment_log.auth_acknowledgment_log_raw_payload IS 'raw payload string in gateway ack request';


--
-- Name: COLUMN auth_acknowledgment_log.auth_acknowledgment_log_created_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_acknowledgment_log.auth_acknowledgment_log_created_at IS 'created timestamp';


--
-- Name: COLUMN auth_acknowledgment_log.auth_acknowledgment_log_updated_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_acknowledgment_log.auth_acknowledgment_log_updated_at IS 'updated timestamp';


--
-- Name: auth_attempt_log; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auth_attempt_log (
    auth_attempt_log_id bigint NOT NULL,
    auth_attempt_log_client_type character varying(30) NOT NULL,
    auth_attempt_log_client_interface character varying(20) NOT NULL,
    auth_attempt_log_client_ip character varying(45) NOT NULL,
    auth_attempt_log_client_mac_address character varying(12) NOT NULL,
    auth_attempt_log_client_gateway_name text NOT NULL,
    auth_attempt_log_client_url text NOT NULL,
    auth_attempt_log_client_hash_id character varying(64) NOT NULL,
    auth_attempt_log_origin_url text NOT NULL,
    auth_attempt_log_theme_spec_path text NOT NULL,
    auth_attempt_log_opennds_version character varying(20) NOT NULL,
    auth_attempt_log_created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    auth_attempt_log_updated_at timestamp without time zone,
    auth_attempt_log_rhid text NOT NULL,
    auth_attempt_log_session_length_minutes bigint NOT NULL,
    auth_attempt_log_upload_rate_threshold bigint NOT NULL,
    auth_attempt_log_download_rate_threshold bigint NOT NULL,
    auth_attempt_log_upload_quota bigint NOT NULL,
    auth_attempt_log_download_quota bigint NOT NULL,
    auth_attempt_log_custom_value json
);


--
-- Name: TABLE auth_attempt_log; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.auth_attempt_log IS 'attempt for gateway auth';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_id IS 'primary key';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_type; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_type IS 'client type. eg) cpi_url';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_interface; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_interface IS 'client network interface eg) br-lan';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_ip; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_ip IS 'lan ip address of client. includes ipv6';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_mac_address; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_mac_address IS 'client mac address';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_gateway_name; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_gateway_name IS 'gateway name client wants to auth';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_url; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_url IS 'auth url for gateway';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_client_hash_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_client_hash_id IS 'client unique value that identified in gateway LAN';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_origin_url; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_origin_url IS 'origin url from gateway auth request';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_theme_spec_path; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_theme_spec_path IS 'theme_spec file in gateway auth request';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_opennds_version; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_opennds_version IS 'opennds version in gateway auth request';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_created_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_created_at IS 'created timestamp';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_updated_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_updated_at IS 'updated timestamp';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_rhid; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_rhid IS 'hashed client hid with gateway password';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_session_length_minutes; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_session_length_minutes IS 'session length minutes';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_upload_rate_threshold; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_upload_rate_threshold IS 'upload rate threshold (KB/s)';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_download_rate_threshold; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_download_rate_threshold IS 'download rate threshold (KB/s)';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_upload_quota; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_upload_quota IS 'upload quota (KB)';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_download_quota; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_download_quota IS 'download quota (KB)';


--
-- Name: COLUMN auth_attempt_log.auth_attempt_log_custom_value; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.auth_attempt_log.auth_attempt_log_custom_value IS 'json serialized custom tags';


--
-- Name: gateway; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.gateway (
    gateway_id bigint NOT NULL,
    gateway_name text NOT NULL,
    gateway_mac_address character varying(12) NOT NULL,
    gateway_name_hash character varying(64) NOT NULL,
    gateway_password_key character varying(256) NOT NULL,
    gateway_created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    gateway_updated_at timestamp without time zone
);


--
-- Name: TABLE gateway; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON TABLE public.gateway IS 'router gateway';


--
-- Name: COLUMN gateway.gateway_id; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_id IS 'primary key';


--
-- Name: COLUMN gateway.gateway_name; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_name IS 'gateway name';


--
-- Name: COLUMN gateway.gateway_mac_address; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_mac_address IS 'mac address of gateway wlan interface';


--
-- Name: COLUMN gateway.gateway_name_hash; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_name_hash IS 'hashed gateway name';


--
-- Name: COLUMN gateway.gateway_password_key; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_password_key IS 'fas key for gateway';


--
-- Name: COLUMN gateway.gateway_created_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_created_at IS 'created timestamp';


--
-- Name: COLUMN gateway.gateway_updated_at; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.gateway.gateway_updated_at IS 'updated timestamp';


--
-- Name: auth_acknowledgment_log auth_acknowledgment_log_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auth_acknowledgment_log
    ADD CONSTRAINT auth_acknowledgment_log_pk PRIMARY KEY (auth_acknowledgment_log_id);


--
-- Name: auth_attempt_log auth_attempt_log_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auth_attempt_log
    ADD CONSTRAINT auth_attempt_log_pk PRIMARY KEY (auth_attempt_log_id);


--
-- Name: gateway gateway_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_pk PRIMARY KEY (gateway_id);


--
-- Name: gateway gateway_pk2; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.gateway
    ADD CONSTRAINT gateway_pk2 UNIQUE (gateway_mac_address);


--
-- Name: auth_attempt_log_auth_attempt_log_client_hash_id_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX auth_attempt_log_auth_attempt_log_client_hash_id_index ON public.auth_attempt_log USING btree (auth_attempt_log_client_hash_id);


--
-- Name: auth_attempt_log_auth_attempt_log_client_mac_address_index; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX auth_attempt_log_auth_attempt_log_client_mac_address_index ON public.auth_attempt_log USING btree (auth_attempt_log_client_mac_address);


--
-- Name: gateway_gateway_name_hash_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX gateway_gateway_name_hash_uindex ON public.gateway USING btree (gateway_name_hash);


--
-- Name: auth_acknowledgment_log update_auth_acknowledgment_log_updated_at_trigger; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER update_auth_acknowledgment_log_updated_at_trigger BEFORE UPDATE ON public.auth_acknowledgment_log FOR EACH ROW EXECUTE FUNCTION public.update_auth_acknowledgment_log_updated_at();


--
-- Name: auth_attempt_log update_auth_attempt_log_updated_at_trigger; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER update_auth_attempt_log_updated_at_trigger BEFORE UPDATE ON public.auth_attempt_log FOR EACH ROW EXECUTE FUNCTION public.update_auth_attempt_log_updated_at();


--
-- Name: gateway update_gateway_updated_at_trigger; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER update_gateway_updated_at_trigger BEFORE UPDATE ON public.gateway FOR EACH ROW EXECUTE FUNCTION public.update_gateway_updated_at();


--
-- PostgreSQL database dump complete
--

