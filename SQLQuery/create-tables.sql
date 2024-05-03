CREATE TABLE IF NOT EXISTS public.broadband
(
    id SERIAL PRIMARY KEY,
    created timestamp without time zone NOT NULL DEFAULT now(),
    address character varying(32),  
    latency DOUBLE PRECISION,
    downloadspeed DOUBLE PRECISION,
    uploadspeed DOUBLE PRECISION,
    isp character varying(256)
)